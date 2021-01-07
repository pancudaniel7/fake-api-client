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
// limitations under the License.cd

package errors

import "fmt"

// ResponseError struct defines a fail response.
type ResponseError struct {
	StatusCode int
	Message    string
	CausedBy   error
}

// Error returns error string response for ResponseError.
func (e ResponseError) Error() string {
	return fmt.Sprintf("%s, status code: %d caused by: %s", e.Message, e.StatusCode, e.CausedBy)
}

// ResponseError struct defines a fail request.
type RequestError struct {
	Message  string
	CausedBy error
}

// Error returns error string response for RequestError.
func (e RequestError) Error() string {
	return fmt.Sprintf("%s, caused by: %s", e.Message, e.CausedBy)
}
