// Copyright 2020 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package model

import (
	"strings"

	"github.com/spf13/viper"
)

// Route struct
type Route struct {
	Path    string `mapstructure:"path"`
	Request struct {
		Method string `mapstructure:"method"`
	} `mapstructure:"request"`
	Response struct {
		StatusCode int `mapstructure:"statusCode"`
		Headers    []struct {
			Key   string `mapstructure:"key"`
			Value string `mapstructure:"value"`
		} `mapstructure:"headers"`
		Body string `mapstructure:"body"`
	} `mapstructure:"response"`
	Chaos struct {
		Latency  string `mapstructure:"latency"`
		FailRate string `mapstructure:"failRate"`
	} `mapstructure:"chaos"`
}

// GetDebugRoutes get a list of debug routes
func GetDebugRoutes() ([]Route, error) {
	var routes []Route

	err := viper.UnmarshalKey("debug", &routes)

	if err != nil {
		return routes, err
	}

	return routes, nil
}

// GetMockRoutes get a list of mock routes
func GetMockRoutes() ([]Route, error) {
	var routes []Route

	err := viper.UnmarshalKey("mock", &routes)

	if err != nil {
		return routes, err
	}

	return routes, nil
}

// GetRoute get route object with path
func GetRoute(path string, method string) Route {
	debugRoutes, _ := GetDebugRoutes()

	for _, route := range debugRoutes {
		if path == route.Path {
			return route
		}
	}

	mockRoutes, _ := GetMockRoutes()

	for _, route := range mockRoutes {
		if path == route.Path && strings.ToLower(method) == strings.ToLower(route.Request.Method) {
			return route
		}
	}

	return Route{}
}
