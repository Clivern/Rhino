// Copyright 2020 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/clivern/rhino/internal/app/model"
	"github.com/clivern/rhino/internal/app/module"

	"github.com/gin-gonic/gin"
)

// Debug controller
func Debug(c *gin.Context) {
	var bodyBytes []byte

	// Workaround for issue https://github.com/gin-gonic/gin/issues/1651
	if c.Request.Body != nil {
		bodyBytes, _ = ioutil.ReadAll(c.Request.Body)
	}

	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

	logger, _ := module.NewLogger()

	defer logger.Sync()

	route := model.GetRoute(c.FullPath(), "")

	rand.Seed(time.Now().UnixNano())

	failCount, _ := strconv.Atoi(strings.ReplaceAll(route.Chaos.FailRate, "%", ""))

	if rand.Intn(100) < failCount {
		c.Status(http.StatusInternalServerError)
		return
	}

	latencySeconds, _ := strconv.Atoi(strings.ReplaceAll(route.Chaos.Latency, "s", ""))

	time.Sleep(time.Duration(latencySeconds) * time.Second)

	header, _ := json.Marshal(c.Request.Header)

	logger.Info(fmt.Sprintf(
		"%s:%s %s %s",
		c.Request.Method,
		c.Request.URL,
		header,
		string(bodyBytes),
	))

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}
