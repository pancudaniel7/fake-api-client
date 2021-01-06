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

// +build integration

package test

import (
	"encoding/json"
	"github.com/pancudaniel7/fake-api-client/pkg/api"
	"github.com/pancudaniel7/fake-api-client/pkg/errors"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)


func TestAllAccountListing(t *testing.T) {

	expAccList := []api.Account{
		readFileAsAccount("data/account.json"),
		readFileAsAccount("data/second-account.json")}

	for _, acc := range expAccList {
		_, err := acc.Create()
		if err != nil {
			log.Fatalf("fail to create account resource: %s", err)
		}
	}

	accResList, err := expAccList[0].List("", "")
	if err != nil {
		log.Fatalf("fail to list all accounts: %s", err)
	}

	assert.EqualValues(t, 2, len(accResList))

	assert.EqualValues(t, expAccList[0].ID, accResList[0].(api.Account).ID)
	assert.EqualValues(t, expAccList[0].Type, accResList[0].(api.Account).Type)
	assert.EqualValues(t, expAccList[0].OrganisationID, accResList[0].(api.Account).OrganisationID)
	assert.EqualValues(t, expAccList[0].Attributes, accResList[0].(api.Account).Attributes)

	assert.EqualValues(t, expAccList[1].ID, accResList[1].(api.Account).ID)
	assert.EqualValues(t, expAccList[1].Type, accResList[1].(api.Account).Type)
	assert.EqualValues(t, expAccList[1].OrganisationID, accResList[1].(api.Account).OrganisationID)
	assert.EqualValues(t, expAccList[1].Attributes, accResList[1].(api.Account).Attributes)

	for _, acc := range expAccList {
		deleteAccount(acc)
	}
}

func TestPageAccountListing(t *testing.T) {

	expAccList := []api.Account{
		readFileAsAccount("data/account.json"),
		readFileAsAccount("data/second-account.json"),
		readFileAsAccount("data/third-account.json"),
		readFileAsAccount("data/fourth-account.json")}

	for _, acc := range expAccList {
		_, err := acc.Create()
		if err != nil {
			log.Fatalf("fail to create account resource: %s", err)
		}
	}

	accResList, err := expAccList[0].List("0", "")
	if err != nil {
		log.Fatalf("fail to list all accounts: %s", err)
	}

	assert.EqualValues(t, 2, len(accResList))

	assert.EqualValues(t, expAccList[0].ID, accResList[0].(api.Account).ID)
	assert.EqualValues(t, expAccList[0].Type, accResList[0].(api.Account).Type)
	assert.EqualValues(t, expAccList[0].OrganisationID, accResList[0].(api.Account).OrganisationID)
	assert.EqualValues(t, expAccList[0].Attributes, accResList[0].(api.Account).Attributes)

	assert.EqualValues(t, expAccList[1].ID, accResList[1].(api.Account).ID)
	assert.EqualValues(t, expAccList[1].Type, accResList[1].(api.Account).Type)
	assert.EqualValues(t, expAccList[1].OrganisationID, accResList[1].(api.Account).OrganisationID)
	assert.EqualValues(t, expAccList[1].Attributes, accResList[1].(api.Account).Attributes)

	for _, acc := range expAccList {
		deleteAccount(acc)
	}
}

func TestPageAndSizeAccountListing(t *testing.T) {

	expAccList := []api.Account{
		readFileAsAccount("data/account.json"),
		readFileAsAccount("data/second-account.json"),
		readFileAsAccount("data/third-account.json"),
		readFileAsAccount("data/fourth-account.json")}

	for _, acc := range expAccList {
		_, err := acc.Create()
		if err != nil {
			log.Fatalf("fail to create account resource: %s", err)
		}
	}

	accResList, err := expAccList[0].List("1", "1")
	if err != nil {
		log.Fatalf("fail to list all accounts: %s", err)
	}

	assert.EqualValues(t, 1, len(accResList))

	assert.EqualValues(t, expAccList[1].ID, accResList[0].(api.Account).ID)
	assert.EqualValues(t, expAccList[1].Type, accResList[0].(api.Account).Type)
	assert.EqualValues(t, expAccList[1].OrganisationID, accResList[0].(api.Account).OrganisationID)
	assert.EqualValues(t, expAccList[1].Attributes, accResList[0].(api.Account).Attributes)

	for _, acc := range expAccList {
		deleteAccount(acc)
	}
}

func TestFailAccountListingWithInvalidPageNumber(t *testing.T) {
	acc := readFileAsAccount("data/account.json")

	_, err := acc.Create()
	if err != nil {
		log.Fatalf("fail to create first account resource: %s", err)
	}

	resAcc, err := acc.List("999999999", "")

	assert.EqualValues(t, len(resAcc), 0)

	deleteAccount(acc)
}

func TestAccountListingById(t *testing.T) {

	expAccList := []api.Account{
		readFileAsAccount("data/account.json"),
		readFileAsAccount("data/second-account.json"),
		readFileAsAccount("data/third-account.json"),
		readFileAsAccount("data/fourth-account.json")}

	for _, acc := range expAccList {
		_, err := acc.Create()
		if err != nil {
			log.Fatalf("fail to create account resource: %s", err)
		}
	}

	resAcc, err := expAccList[2].ListById()
	if err != nil {
		log.Fatalf("fail to list all accounts: %s", err)
	}

	actAcc := resAcc.(*api.Account)

	assert.EqualValues(t, expAccList[2].ID, actAcc.ID)
	assert.EqualValues(t, expAccList[2].Type, actAcc.Type)
	assert.EqualValues(t, expAccList[2].OrganisationID, actAcc.OrganisationID)
	assert.EqualValues(t, expAccList[2].Attributes, actAcc.Attributes)

	for _, acc := range expAccList {
		deleteAccount(acc)
	}
}

func TestFailAccountListingByInvalidId(t *testing.T) {
	acc := readFileAsAccount("data/account.json")

	_, err := acc.Create()
	if err != nil {
		log.Fatalf("fail to create first account resource: %s", err)
	}

	tempID := acc.ID
	acc.ID = "wrong id value"
	_, actErr := acc.ListById()

	expErr := errors.ResponseError{
		StatusCode: 400,
		Message:    "request error with different status code, expected: 200 but returned: 400 with error message: id is not a valid uuid",
		CausedBy:   nil,
	}

	assert.EqualValues(t, expErr, actErr)

	acc.ID = tempID
	deleteAccount(acc)
}

func readFileAsAccount(path string) api.Account {
	accJsonBytes := readFileAsBytes(path)

	acc := api.Account{}
	if err := json.Unmarshal(accJsonBytes, &acc); err != nil {
		log.Fatalf("Fail to unmarshal account json file bytes: %s", err)
	}
	return acc
}
