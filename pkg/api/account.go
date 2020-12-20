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
	"fmt"
	"github.com/pancudaniel7/fake-api-client/configs"
	_http "github.com/pancudaniel7/fake-api-client/internal/http"
	"github.com/pancudaniel7/fake-api-client/pkg/errors"
	"net/http"
	"time"
)

type Account struct {
	ID             string     `json:"id"`
	CreatedOn      time.Time  `json:"created_on"`
	ModifiedOn     time.Time  `json:"modified_on"`
	OrganisationID string     `json:"organisation_id"`
	Type           string     `json:"type"`
	Version        int        `json:"version"`
	Attributes     Attributes `json:"attributes"`
}

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

func (a Account) Create() (Resource, error) {
	reqBody := _http.Body{Data: a}

	b, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	reqUrl := configs.Props().BaseAPIURL + accountPath
	req, err := http.NewRequest(http.MethodPost, reqUrl, bytes.NewReader(b))
	if err != nil {
		return nil, errors.RequestError{Message: fmt.Sprintf("fail to create account post request object: %s", err)}
	}

	resAcc := &Account{}
	return resAcc, _http.APIClient().SendRequest(req, http.StatusCreated, resAcc)
}

func (a Account) List() (Resource, error) {
	return nil, nil
}

func (a Account) Delete() (Resource, error) {
	return nil, nil
}
