// Copyright 2020 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package cmd

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/clivern/rhino/core/controller"
	"github.com/clivern/rhino/core/middleware"
	"github.com/clivern/rhino/core/model"
	"github.com/clivern/rhino/core/module"

	"github.com/drone/envsubst"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var config string

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start Rhino Server",
	Run: func(cmd *cobra.Command, args []string) {
		configUnparsed, err := ioutil.ReadFile(config)

		if err != nil {
			panic(fmt.Sprintf(
				"Error while reading config file [%s]: %s",
				config,
				err.Error(),
			))
		}

		configParsed, err := envsubst.EvalEnv(string(configUnparsed))

		if err != nil {
			panic(fmt.Sprintf(
				"Error while parsing config file [%s]: %s",
				config,
				err.Error(),
			))
		}

		viper.SetConfigType("json")
		err = viper.ReadConfig(bytes.NewBuffer([]byte(configParsed)))

		if err != nil {
			panic(fmt.Sprintf(
				"Error while loading configs [%s]: %s",
				config,
				err.Error(),
			))
		}

		if viper.GetString("log.output") != "stdout" {
			fs := module.FileSystem{}
			dir, _ := filepath.Split(viper.GetString("log.output"))

			if !fs.DirExists(dir) {
				if _, err := fs.EnsureDir(dir, 777); err != nil {
					panic(fmt.Sprintf(
						"Directory [%s] creation failed with error: %s",
						dir,
						err.Error(),
					))
				}
			}

			if !fs.FileExists(viper.GetString("log.output")) {
				f, err := os.Create(viper.GetString("log.output"))
				if err != nil {
					panic(fmt.Sprintf(
						"Error while creating log file [%s]: %s",
						viper.GetString("log.output"),
						err.Error(),
					))
				}
				defer f.Close()
			}
		}

		if viper.GetString("log.output") == "stdout" {
			gin.DefaultWriter = os.Stdout
			log.SetOutput(os.Stdout)
		} else {
			f, _ := os.Create(viper.GetString("log.output"))
			gin.DefaultWriter = io.MultiWriter(f)
			log.SetOutput(f)
		}

		lvl := strings.ToLower(viper.GetString("log.level"))
		level, err := log.ParseLevel(lvl)

		if err != nil {
			level = log.InfoLevel
		}

		log.SetLevel(level)

		if viper.GetString("app.mode") == "prod" {
			gin.SetMode(gin.ReleaseMode)
			gin.DefaultWriter = ioutil.Discard
			gin.DisableConsoleColor()
		}

		if viper.GetString("log.format") == "json" {
			log.SetFormatter(&log.JSONFormatter{})
		} else {
			log.SetFormatter(&log.TextFormatter{})
		}

		r := gin.Default()

		r.Use(middleware.Cors())

		r.GET("/favicon.ico", func(c *gin.Context) {
			c.String(http.StatusNoContent, "")
		})

		r.GET("/", controller.Index)
		r.GET("/_health", controller.Health)
		r.GET("/api/requests", controller.Requests)

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
			if strings.ToLower(route.Request.Method) == "get" {
				r.GET(route.Path, controller.Mock)
			} else if strings.ToLower(route.Request.Method) == "post" {
				r.POST(route.Path, controller.Mock)
			} else if strings.ToLower(route.Request.Method) == "put" {
				r.PUT(route.Path, controller.Mock)
			} else if strings.ToLower(route.Request.Method) == "delete" {
				r.DELETE(route.Path, controller.Mock)
			} else if strings.ToLower(route.Request.Method) == "patch" {
				r.PATCH(route.Path, controller.Mock)
			} else if strings.ToLower(route.Request.Method) == "head" {
				r.HEAD(route.Path, controller.Mock)
			} else if strings.ToLower(route.Request.Method) == "options" {
				r.OPTIONS(route.Path, controller.Mock)
			}
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
	},
}

func init() {
	serveCmd.Flags().StringVarP(&config, "config", "c", "config.prod.yml", "Absolute path to config file (required)")
	serveCmd.MarkFlagRequired("config")
	rootCmd.AddCommand(serveCmd)
}
