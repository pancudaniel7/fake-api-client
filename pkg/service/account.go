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

package service

import (
	"bytes"
	"encoding/json"
	"github.com/pancudaniel7/fake-api-client/configs"
	_http "github.com/pancudaniel7/fake-api-client/internal/http"
	"github.com/pancudaniel7/fake-api-client/pkg/model"
	"net/http"
)

type Account struct{}

// Create account method used for creating account resource type.
// The method creates the request for creating account resource
// or returns jsonError if it cannot parse the reqBody object,
// or RequestError if cannot create the request,
// or it returns the account if the account was successful created.
func (a Account) Create(acc model.Resource) (model.Resource, error) {
	reqBody := _http.Body{Data: acc}

	b, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	reqUrl := configs.Properties().BaseAPIURL + _http.AccountPath
	body := bytes.NewReader(b)

	req, err := _http.CreateRequest(http.MethodPost, reqUrl, body)
	if err != nil {
		return nil, err
	}

	resAcc := &model.Account{}
	return resAcc, _http.APIClient().SendRequest(req, http.StatusCreated, resAcc)
}

// List method returns all account list if pageNum and pageSize are empty,
// or specific account list by the pageNum and pageSize.
// Also this method returns RequestError if the request could not be created
// or if the response returns an error content.
func (a Account) List(pageNum, pageSize string) ([]model.Resource, error) {

	pagParam := _http.BuildPagination(pageNum, pageSize)
	reqUrl := configs.Properties().BaseAPIURL + _http.AccountPath
	if len(pagParam) != 0 {
		reqUrl += "?" + pagParam
	}

	req, err := _http.CreateRequest(http.MethodGet, reqUrl, nil)
	if err != nil {
		return nil, err
	}

	resAccList := &[]model.Account{}
	if err = _http.APIClient().SendRequest(req, http.StatusOK, resAccList); err != nil {
		return nil, err
	}

	return convertSlicesAccountToResource(*resAccList), nil
}

// ListBy method returns one account entity requested by the account id.
// Also this method returns RequestError if the request could not be created
// or if the response returns an error content.
func (a Account) ListBy(id string) (model.Resource, error) {

	reqUrl := configs.Properties().BaseAPIURL + _http.AccountPath +
		"/" + id

	req, err := _http.CreateRequest(http.MethodGet, reqUrl, nil)
	if err != nil {
		return nil, err
	}

	resAcc := &model.Account{}
	return resAcc, _http.APIClient().SendRequest(req, http.StatusOK, resAcc)
}

// DeleteBy method delete account entity by account id.
// Also this method returns RequestError if the request could not be created
// or if the response returns an error content.
func (a Account) DeleteBy(id string) error {
	reqUrl := configs.Properties().BaseAPIURL + _http.AccountPath +
		"/" + id +
		"?" + _http.VersionLabel + configs.Properties().HttpRecordVersion

	req, err := _http.CreateRequest(http.MethodDelete, reqUrl, nil)
	if err != nil {
		return err
	}

	return _http.APIClient().SendRequest(req, http.StatusNoContent, nil)
}

// convertSlicesAccountToResource helps with converting list of Account types list
// in to a list of Resource types list.
func convertSlicesAccountToResource(accList []model.Account) []model.Resource {
	resAccRes := []model.Resource{}
	for _, resAcc := range accList {
		resAccRes = append(resAccRes, resAcc)
	}
	return resAccRes
}
