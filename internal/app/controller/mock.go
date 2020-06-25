// Copyright 2020 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package controller

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/clivern/rhino/internal/app/model"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// Mock controller
func Mock(c *gin.Context) {
	var bodyBytes []byte

	// Workaround for issue https://github.com/gin-gonic/gin/issues/1651
	if c.Request.Body != nil {
		bodyBytes, _ = ioutil.ReadAll(c.Request.Body)
	}

	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

	header, _ := json.Marshal(c.Request.Header)

	parameters := make(map[string]string)

	for k, v := range c.Request.URL.Query() {
		parameters[k] = v[0]
	}

	route := model.GetRoute(c.FullPath(), c.Request.Method, parameters)

	rand.Seed(time.Now().UnixNano())

	failCount, _ := strconv.Atoi(strings.Replace(route.Chaos.FailRate, "%", "", -1))

	if rand.Intn(100) < failCount {
		log.WithFields(log.Fields{
			"method":     c.Request.Method,
			"url":        c.Request.URL.Path,
			"header":     header,
			"parameters": parameters,
			"body":       string(bodyBytes),
		}).Info("Failed Request")

		c.Status(http.StatusInternalServerError)
		return
	}

	latencySeconds, _ := strconv.Atoi(strings.Replace(route.Chaos.Latency, "s", "", -1))

	time.Sleep(time.Duration(latencySeconds) * time.Second)

	log.WithFields(log.Fields{
		"method":     c.Request.Method,
		"url":        c.Request.URL.Path,
		"header":     header,
		"parameters": parameters,
		"body":       string(bodyBytes),
	}).Info("Request Success")

	for _, header := range route.Response.Headers {
		c.Header(header.Key, header.Value)
	}

	if strings.Contains(route.Response.Body, "@json:") {
		path := strings.Replace(route.Response.Body, "@json:", "", -1)
		content, err := ioutil.ReadFile(path)

		if err != nil {
			panic(err)
		}

		route.Response.Body = string(content)
	}

	for _, param := range c.Params {
		route.Response.Body = strings.Replace(route.Response.Body, ":"+param.Key, param.Value, -1)
	}

	for key, value := range route.Request.Parameters {

		if !strings.HasPrefix(value, ":") {
			continue
		}

		route.Response.Body = strings.Replace(route.Response.Body, value, parameters[key], -1)
	}

	c.String(route.Response.StatusCode, route.Response.Body)
}
