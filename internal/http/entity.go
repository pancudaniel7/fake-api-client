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

package http

// ErrorResponse type needed to read error response content.
type ErrorResponse struct {
	Message string `json:"error_message" validate:"required"`
}

// Body type is used for response nad request http body content.
type Body struct {
	Data interface{} `json:"data" validate:"required"`
}
