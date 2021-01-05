// Copyright 2019 Form3 Financial Cloud
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package configs provides library configuration functionality
package configs

import (
	"log"
	"os"
	"sync"
	"time"
)

type properties struct {
	BaseAPIURL          string
	HttpClientTimeout   time.Duration
	HttpRecordVersion   string
	HttpDefaultPageSize string
}

var (
	once sync.Once
	p    *properties
)

func Props() *properties {
	once.Do(func() {
		p = &properties{
			BaseAPIURL:          getEnvOrDefaultString("BASE_API_URL", "http://localhost:8080/v1"),
			HttpClientTimeout:   getEnvOrDefaultDuration("HTTP_CLIENT_REQ_TIME_OUT", time.Minute),
			HttpRecordVersion:   getEnvOrDefaultString("HTTP_RECORD_VERSION", "0"),
			HttpDefaultPageSize: getEnvOrDefaultString("HTTP_DEFAULT_PAGE_SIZE", "2"),
		}
	})
	return p
}

func getEnvOrDefaultDuration(key string, d time.Duration) time.Duration {
	env := os.Getenv(key)
	if isEmpty(env) {
		return d
	}

	i, err := time.ParseDuration(env)
	if err != nil {
		log.Fatalf("fail to convert: %s with value: %s in to time duration: %s", key, env, err)
	}
	return i
}

func getEnvOrDefaultString(key, d string) string {
	env := os.Getenv(key)
	if isEmpty(env) {
		return d
	}
	return env
}

func isEmpty(v string) bool {
	return len(v) == 0
}
