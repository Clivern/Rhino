// Copyright 2020 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package model

import (
	"encoding/json"
	"time"
)

// Header struct
type Header struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// Request struct
type Request struct {
	Route      string    `json:"route"`
	URI        string    `json:"uri"`
	Method     string    `json:"method"`
	StatusCode int       `json:"statusCode"`
	Headers    []Header  `json:"headers"`
	Status     string    `json:"status"`
	Body       string    `json:"body"`
	Time       time.Time `json:"time"`
}

// LoadFromJSON update object from json
func (p *Request) LoadFromJSON(data []byte) (bool, error) {
	err := json.Unmarshal(data, &p)
	if err != nil {
		return false, err
	}
	return true, nil
}

// ConvertToJSON convert object to json
func (p *Request) ConvertToJSON() (string, error) {
	data, err := json.Marshal(&p)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
