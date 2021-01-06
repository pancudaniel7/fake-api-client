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

func TestAccountCreation(t *testing.T) {

	expAcc := readFileAsAccount("data/account.json")
	resResource, err := expAcc.Create()
	if err != nil {
		log.Fatalf("fail to create account resource: %s", err)
	}

	resAcc := resResource.(*api.Account)

	assert.EqualValues(t, expAcc.ID, resAcc.ID)
	assert.EqualValues(t, expAcc.Type, resAcc.Type)
	assert.EqualValues(t, expAcc.OrganisationID, resAcc.OrganisationID)
	assert.EqualValues(t, expAcc.Attributes, resAcc.Attributes)

	deleteAccount(expAcc)
}

func TestFailAccountCreationForSameId(t *testing.T) {
	acc := readFileAsAccount("data/account.json")

	_, actErr := acc.Create()
	if actErr != nil {
		log.Fatalf("fail to create first account resource: %s", actErr)
	}

	expErr := errors.ResponseError{
		StatusCode: 409,
		Message:    "request error with different status code, expected: 201 but returned: 409 with error message: Account cannot be created as it violates a duplicate constraint",
		CausedBy:   nil,
	}

	_, actErr = acc.Create()

	assert.NotNil(t, actErr)
	assert.EqualValues(t, expErr, actErr)

	deleteAccount(acc)
}

func TestFailAccountCreationWithInvalidValues(t *testing.T) {
	acc := readFileAsAccount("data/invalid-account.json")

	expErr := errors.ResponseError{
		StatusCode: 400,
		Message: "request error with different status code, " +
			"expected: 201 but returned: 400 with error message: " +
			"validation failure list:\nvalidation failure list:\nvalidation " +
			"failure list:\naccount_classification in body should be one of " +
			"[Personal Business]\naccount_number in body should match " +
			"'^[A-Z0-9]{0,64}$'\nbank_id in body should match '^[A-Z0-9]{0,16}$'\nbank_id_code " +
			"in body should match '^[A-Z]{0,16}$'\nbase_currency in body should match " +
			"'^[A-Z]{3}$'\nbic in body should match '^([A-Z]{6}[A-Z0-9]{2}|[A-Z]{6}[A-Z0-9]{5})$'\ncountry " +
			"in body should match '^[A-Z]{2}$'\niban in body should match '^[A-Z]{2}[0-9]{2}[A-Z0-9]{0,64}$'\nid " +
			"in body must be of type uuid: \"wrong uuid\"\norganisation_id in body must be of type uuid: \"wrong organisation id\"\ntype " +
			"in body should be one of [accounts]",
		CausedBy: nil,
	}

	_, actArr := acc.Create()

	assert.NotNil(t, actArr)
	assert.EqualValues(t, expErr, actArr)
}

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

func TestAccountListById(t *testing.T) {

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

func TestFailAccountListByInvalidId(t *testing.T) {
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

func TestFailAccountDeleteWithWrongId(t *testing.T) {
	acc := readFileAsAccount("data/account.json")

	_, err := acc.Create()
	if err != nil {
		log.Fatalf("fail to create first account resource: %s", err)
	}

	tempID := acc.ID
	acc.ID = "wrong id value"
	actErr := acc.Delete()

	expErr := errors.ResponseError{
		StatusCode: 400,
		Message:    "request error with different status code, expected: 204 but returned: 400 with error message: id is not a valid uuid",
		CausedBy:   nil,
	}

	assert.EqualValues(t, expErr, actErr)

	acc.ID = tempID
	deleteAccount(acc)
}

func deleteAccount(acc api.Account) {
	if err := acc.Delete(); err != nil {
		log.Fatalf("fail to delete account resource:  %s", err)
	}
}

func readFileAsAccount(path string) api.Account {
	accJsonBytes := readFileAsBytes(path)

	acc := api.Account{}
	if err := json.Unmarshal(accJsonBytes, &acc); err != nil {
		log.Fatalf("Fail to unmarshal account json file bytes: %s", err)
	}
	return acc
}
