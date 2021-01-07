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

package api

import (
	"bytes"
	"encoding/json"
	"github.com/pancudaniel7/fake-api-client/configs"
	_http "github.com/pancudaniel7/fake-api-client/internal/http"
	"net/http"
	"time"
)

// Account struct is used for defining Account resource.
type Account struct {
	ID             string     `json:"id"`
	CreatedOn      time.Time  `json:"created_on"`
	ModifiedOn     time.Time  `json:"modified_on"`
	OrganisationID string     `json:"organisation_id"`
	Type           string     `json:"type"`
	Version        int        `json:"version"`
	Attributes     Attributes `json:"attributes"`
}

// Attributes struct is used for defining Account resource attributes.
type Attributes struct {
	AccountNumber               string   `json:"account_number"`
	AccountClassification       string   `json:"account_classification"`
	AccountMatchingOptOut       bool     `json:"account_matching_opt_out"`
	AlternativeBankAccountNames []string `json:"alternative_bank_account_names"`
	BankID                      string   `json:"bank_id"`
	BankIDCode                  string   `json:"bank_id_code"`
	BaseCurrency                string   `json:"base_currency"`
	Bic                         string   `json:"bic"`
	Country                     string   `json:"country"`
	CustomerID                  string   `json:"customer_id"`
	JointAccount                bool     `json:"joint_account"`
	Iban                        string   `json:"iban"`
}

// Create account method used for creating account resource type.
// The method creates the request for creating account resource
// or returns jsonError if it cannot parse the reqBody object,
// or RequestError if cannot create the request,
// or it returns the account if the account was successful created.
func (a Account) Create() (Resource, error) {
	reqBody := _http.Body{Data: a}

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

	resAcc := &Account{}
	return resAcc, _http.APIClient().SendRequest(req, http.StatusCreated, resAcc)
}

// List method returns all account list if pageNum and pageSize are empty,
// or specific account list by the pageNum and pageSize.
// Also this method returns RequestError if the request could not be created
// or if the response returns an error content.
func (a Account) List(pageNum, pageSize string) ([]Resource, error) {

	pagParam := _http.BuildPagination(pageNum, pageSize)
	reqUrl := configs.Properties().BaseAPIURL + _http.AccountPath
	if len(pagParam) != 0 {
		reqUrl += "?" + pagParam
	}

	req, err := _http.CreateRequest(http.MethodGet, reqUrl, nil)
	if err != nil {
		return nil, err
	}

	resAccList := &[]Account{}
	if err = _http.APIClient().SendRequest(req, http.StatusOK, resAccList); err != nil {
		return nil, err
	}

	return convertSlicesAccountToResource(*resAccList), nil
}

// ListById method returns one account entity requested by the account id.
// Also this method returns RequestError if the request could not be created
// or if the response returns an error content.
func (a Account) ListById() (Resource, error) {

	reqUrl := configs.Properties().BaseAPIURL + _http.AccountPath +
		"/" + a.ID

	req, err := _http.CreateRequest(http.MethodGet, reqUrl, nil)
	if err != nil {
		return nil, err
	}

	resAcc := &Account{}
	return resAcc, _http.APIClient().SendRequest(req, http.StatusOK, resAcc)
}

// Delete method delete account entity by account id.
// Also this method returns RequestError if the request could not be created
// or if the response returns an error content.
func (a Account) Delete() error {
	reqUrl := configs.Properties().BaseAPIURL + _http.AccountPath +
		"/" + a.ID +
		"?" + _http.VersionLabel + configs.Properties().HttpRecordVersion

	req, err := _http.CreateRequest(http.MethodDelete, reqUrl, nil)
	if err != nil {
		return err
	}

	return _http.APIClient().SendRequest(req, http.StatusNoContent, nil)
}

// convertSlicesAccountToResource helps with converting list of Account types list
// in to a list of Resource types list.
func convertSlicesAccountToResource(accList []Account) []Resource {
	resAccRes := []Resource{}
	for _, resAcc := range accList {
		resAccRes = append(resAccRes, resAcc)
	}
	return resAccRes
}
