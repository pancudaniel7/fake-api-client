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

package model

import (
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