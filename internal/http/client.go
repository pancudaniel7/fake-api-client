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

import (
	"encoding/json"
	"fmt"
	"github.com/pancudaniel7/fake-api-client/configs"
	"github.com/pancudaniel7/fake-api-client/pkg/errors"
	"io"
	"net/http"
	"sync"
)

// An clientAPI represents the struct type used
// to create singleton http client object.
type clientAPI struct {
	HTTPClient *http.Client
}

var (
	once sync.Once
	c    *clientAPI
)

// APIClient function return the clientAPI singleton object.
// The object is used to communicate using http with the main server.
// APIClient is initialised once (lazy load) when it used for the first time,
// also it contains HttpClientTimeout property.
func APIClient() *clientAPI {
	once.Do(func() {
		c = &clientAPI{
			HTTPClient: &http.Client{
				Timeout: configs.Properties().HttpClientTimeout,
			},
		}
	})
	return c
}

// SendRequest method can send http requests and return http responses.
// The method uses req as request object.
// expCode is used to verify if the response is the correct one,
// if the response dose not contains the expected status code the method
// will return ErrorResponse containing description about the error.
// the resData object is used to receive response body data or to send request body data.
// if response cannot be decoded the method will return SyntaxError.
func (c *clientAPI) SendRequest(req *http.Request, expCode int, resData interface{}) error {
	req.Header.Set("Accept", "application/json")

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if err = handleExpectedStatusCode(*res, expCode); err != nil {
		return err
	}

	fullResponse := Body{
		Data: resData,
	}

	if resData != nil {
		if err = json.NewDecoder(res.Body).Decode(&fullResponse); err != nil {
			return err
		}
	}
	return nil
}

// CreateRequest function is used to create request using http verb method,
// reqUrl and data request body using NewRequest golang object.
// If the object fails to return request will return custom RequestError.
func CreateRequest(method, reqUrl string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, reqUrl, body)
	if err != nil {
		return nil, errors.RequestError{
			Message:  fmt.Sprintf("fail to create %s request object: %s", method, err),
			CausedBy: err}
	}

	return req, nil
}

// BuildPagination function return url query parameters regarding pageNum and pageSize values.
func BuildPagination(pageNum, pageSize string) string {
	if len(pageNum) == 0 {
		return ""
	} else if len(pageSize) == 0 {
		return pageNumberLabel + pageNum + "&" +
			pageSizeLabel + configs.Properties().HttpDefaultPageSize
	}

	return pageNumberLabel + pageNum + "&" +
		pageSizeLabel + pageSize
}

// handleExpectedStatusCode function is used to handle ErrorResponse content type coming from requests.
// The returned type is ResponseError containing all details is needed regarding the error.
func handleExpectedStatusCode(res http.Response, expCode int) error {
	if res.StatusCode != expCode {
		var errRes ErrorResponse

		errMsg := "request error with different status code, expected: %d but returned: %d with error message: %s"

		var err error
		if err = json.NewDecoder(res.Body).Decode(&errRes); err != nil {
			decErrMsg := "fail to decode error response body with error message: %s"

			return errors.ResponseError{
				Message:    fmt.Sprintf(errMsg, expCode, res.StatusCode, fmt.Sprintf(decErrMsg, err)),
				StatusCode: res.StatusCode,
				CausedBy:   err}
		}

		return errors.ResponseError{
			Message:    fmt.Sprintf(errMsg, expCode, res.StatusCode, errRes.Message),
			StatusCode: res.StatusCode}
	}
	return nil
}
