// Copyright 2020 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package controller

import (
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/clivern/rhino/internal/app/model"

	"github.com/gin-gonic/gin"
)

// Debug controller
func Debug(c *gin.Context) {
	route := model.GetRoute(c.FullPath(), "")

	rand.Seed(time.Now().UnixNano())

	failCount, _ := strconv.Atoi(strings.ReplaceAll(route.Chaos.FailRate, "%", ""))

	if rand.Intn(100) < failCount {
		c.Status(http.StatusInternalServerError)
		return
	}

	latencySeconds, _ := strconv.Atoi(strings.ReplaceAll(route.Chaos.Latency, "s", ""))

	time.Sleep(time.Duration(latencySeconds) * time.Second)

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}
