// Copyright 2020 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/clivern/rhino/internal/app/controller"
	"github.com/clivern/rhino/internal/app/model"
	"github.com/clivern/rhino/internal/app/util"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	var configFile string
	var get string

	flag.StringVar(&configFile, "config", "config.prod.json", "config")
	flag.StringVar(&get, "get", "", "get")
	flag.Parse()

	if get == "release" {
		fmt.Println(
			fmt.Sprintf(
				`Rhino Version %v Commit %v, Built @%v`,
				version,
				commit,
				date,
			),
		)
		return
	}

	config, err := util.ReadFile(configFile)

	if err != nil {
		panic(fmt.Sprintf(
			"Error while reading config file [%s]: %s",
			configFile,
			err.Error(),
		))
	}

	viper.SetConfigType("json")
	err = viper.ReadConfig(bytes.NewBuffer([]byte(config)))

	if err != nil {
		panic(fmt.Sprintf(
			"Error while loading configs [%s]: %s",
			configFile,
			err.Error(),
		))
	}

	if viper.GetString("app.mode") == "prod" {
		gin.SetMode(gin.ReleaseMode)
		gin.DefaultWriter = ioutil.Discard
		gin.DisableConsoleColor()
	}

	r := gin.Default()

	r.GET("/favicon.ico", func(c *gin.Context) {
		c.String(http.StatusNoContent, "")
	})

	r.GET("/_health", controller.Health)

	debugRoutes, err := model.GetDebugRoutes()

	if err != nil {
		panic(fmt.Sprintf(
			"Error while building debug routes from config file: %s",
			err.Error(),
		))
	}

	for _, route := range debugRoutes {
		r.Any(route.Path, controller.Debug)
	}

	mockRoutes, err := model.GetMockRoutes()

	if err != nil {
		panic(fmt.Sprintf(
			"Error while building mock routes from config file: %s",
			err.Error(),
		))
	}

	for _, route := range mockRoutes {
		r.Any(route.Path, controller.Mock)
	}

	var runerr error

	if viper.GetBool("app.tls.status") {
		runerr = r.RunTLS(
			fmt.Sprintf(":%s", strconv.Itoa(viper.GetInt("app.port"))),
			viper.GetString("app.tls.pemPath"),
			viper.GetString("app.tls.keyPath"),
		)
	} else {
		runerr = r.Run(
			fmt.Sprintf(":%s", strconv.Itoa(viper.GetInt("app.port"))),
		)
	}

	if runerr != nil {
		panic(runerr.Error())
	}
}
