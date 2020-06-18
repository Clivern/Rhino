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

// Debug controller
func Debug(c *gin.Context) {
	var bodyBytes []byte

	// Workaround for issue https://github.com/gin-gonic/gin/issues/1651
	if c.Request.Body != nil {
		bodyBytes, _ = ioutil.ReadAll(c.Request.Body)
	}

	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	header, _ := json.Marshal(c.Request.Header)

	route := model.GetRoute(c.FullPath(), "")

	rand.Seed(time.Now().UnixNano())

	failCount, _ := strconv.Atoi(strings.Replace(route.Chaos.FailRate, "%", "", -1))

	if rand.Intn(100) < failCount {
		log.WithFields(log.Fields{
			"method": c.Request.Method,
			"url":    c.Request.URL.Path,
			"header": header,
			"body":   string(bodyBytes),
		}).Info("Failed Request")

		c.Status(http.StatusInternalServerError)
		return
	}

	latencySeconds, _ := strconv.Atoi(strings.Replace(route.Chaos.Latency, "s", "", -1))

	time.Sleep(time.Duration(latencySeconds) * time.Second)

	log.WithFields(log.Fields{
		"method": c.Request.Method,
		"url":    c.Request.URL.Path,
		"header": header,
		"body":   string(bodyBytes),
	}).Info("Request Success")

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}
