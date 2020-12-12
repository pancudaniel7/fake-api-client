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

import "os"

// Properties type is used to keep all env variables library
// properties
type Properties struct {
	APIBaseURLV1 string
}

// p singleton points to library properties object
var p *Properties

// Props fetches environment or default value for every Properties field
// and returns a  pointer p *Properties of this the singleton object.
// If p *Properties is nil the properties will be fetched once.
func Props() *Properties {
	if p == nil {
		p = &Properties{
			APIBaseURLV1: getEnvOrDefault("FAKE_API_BASE_URL_V1", "http://localhost:8080/v1"),
		}
	}
	return p
}

// getEnvOrDefault retrieve env variable for eKey variable name
// or default value d if env variable dose not exists.
func getEnvOrDefault(eKey, d string) string {
	env := os.Getenv(eKey)
	if isEmpty(env) {
		return d
	}
	return env
}

// isEmpty return true or false if the v is a empty string or not
func isEmpty(v string) bool {
	return len(v) == 0
}
