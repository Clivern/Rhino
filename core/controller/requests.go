// Copyright 2020 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Requests controller
func Requests(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)
	c.Writer.Write([]byte(`[{"time": "2020/10/19","uri": "/v2/api/order/1","route": "v2/api/order/:id","statusCode": 200,"method": "GET","headers": [{"key": "x-auth", "value": "123"}],"body": "{\"id\": \"1\"}","status": "success"}]`))
	return
}
