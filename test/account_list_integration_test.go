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
	"github.com/pancudaniel7/fake-api-client/pkg/errors"
	"github.com/pancudaniel7/fake-api-client/pkg/model"
	"github.com/pancudaniel7/fake-api-client/pkg/service"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestAllAccountListing(t *testing.T) {

	expAccList := []model.Account{
		readFileAsAccount("data/account.json"),
		readFileAsAccount("data/second-account.json")}
	a := service.Account{}

	for _, acc := range expAccList {
		_, err := a.Create(acc)
		if err != nil {
			log.Fatalf("fail to create account resource: %s", err)
		}
	}

	accResList, err := a.List("", "")
	if err != nil {
		log.Fatalf("fail to list all accounts: %s", err)
	}

	assert.EqualValues(t, 2, len(accResList))

	assert.EqualValues(t, expAccList[0].ID, accResList[0].(model.Account).ID)
	assert.EqualValues(t, expAccList[0].Type, accResList[0].(model.Account).Type)
	assert.EqualValues(t, expAccList[0].OrganisationID, accResList[0].(model.Account).OrganisationID)
	assert.EqualValues(t, expAccList[0].Attributes, accResList[0].(model.Account).Attributes)

	assert.EqualValues(t, expAccList[1].ID, accResList[1].(model.Account).ID)
	assert.EqualValues(t, expAccList[1].Type, accResList[1].(model.Account).Type)
	assert.EqualValues(t, expAccList[1].OrganisationID, accResList[1].(model.Account).OrganisationID)
	assert.EqualValues(t, expAccList[1].Attributes, accResList[1].(model.Account).Attributes)

	for _, acc := range expAccList {
		deleteAccount(a, acc)
	}
}

func TestPageAccountListing(t *testing.T) {

	expAccList := []model.Account{
		readFileAsAccount("data/account.json"),
		readFileAsAccount("data/second-account.json"),
		readFileAsAccount("data/third-account.json"),
		readFileAsAccount("data/fourth-account.json")}

	a := service.Account{}

	for _, acc := range expAccList {
		_, err := a.Create(acc)
		if err != nil {
			log.Fatalf("fail to create account resource: %s", err)
		}
	}

	accResList, err := a.List("0", "")
	if err != nil {
		log.Fatalf("fail to list all accounts: %s", err)
	}

	assert.EqualValues(t, 2, len(accResList))

	assert.EqualValues(t, expAccList[0].ID, accResList[0].(model.Account).ID)
	assert.EqualValues(t, expAccList[0].Type, accResList[0].(model.Account).Type)
	assert.EqualValues(t, expAccList[0].OrganisationID, accResList[0].(model.Account).OrganisationID)
	assert.EqualValues(t, expAccList[0].Attributes, accResList[0].(model.Account).Attributes)

	assert.EqualValues(t, expAccList[1].ID, accResList[1].(model.Account).ID)
	assert.EqualValues(t, expAccList[1].Type, accResList[1].(model.Account).Type)
	assert.EqualValues(t, expAccList[1].OrganisationID, accResList[1].(model.Account).OrganisationID)
	assert.EqualValues(t, expAccList[1].Attributes, accResList[1].(model.Account).Attributes)

	for _, acc := range expAccList {
		deleteAccount(a, acc)
	}
}

func TestPageAndSizeAccountListing(t *testing.T) {

	expAccList := []model.Account{
		readFileAsAccount("data/account.json"),
		readFileAsAccount("data/second-account.json"),
		readFileAsAccount("data/third-account.json"),
		readFileAsAccount("data/fourth-account.json")}

	a := service.Account{}

	for _, acc := range expAccList {
		_, err := a.Create(acc)
		if err != nil {
			log.Fatalf("fail to create account resource: %s", err)
		}
	}

	accResList, err := a.List("1", "1")
	if err != nil {
		log.Fatalf("fail to list all accounts: %s", err)
	}

	assert.EqualValues(t, 1, len(accResList))

	assert.EqualValues(t, expAccList[1].ID, accResList[0].(model.Account).ID)
	assert.EqualValues(t, expAccList[1].Type, accResList[0].(model.Account).Type)
	assert.EqualValues(t, expAccList[1].OrganisationID, accResList[0].(model.Account).OrganisationID)
	assert.EqualValues(t, expAccList[1].Attributes, accResList[0].(model.Account).Attributes)

	for _, acc := range expAccList {
		deleteAccount(a, acc)
	}
}

func TestFailAccountListingWithInvalidPageNumber(t *testing.T) {
	acc := readFileAsAccount("data/account.json")
	a := service.Account{}

	_, err := a.Create(acc)
	if err != nil {
		log.Fatalf("fail to create first account resource: %s", err)
	}

	resAcc, err := a.List("999999999", "")

	assert.EqualValues(t, len(resAcc), 0)

	deleteAccount(a, acc)
}

func TestAccountListingById(t *testing.T) {

	expAccList := []model.Account{
		readFileAsAccount("data/account.json"),
		readFileAsAccount("data/second-account.json"),
		readFileAsAccount("data/third-account.json"),
		readFileAsAccount("data/fourth-account.json")}

	a := service.Account{}

	for _, acc := range expAccList {
		_, err := a.Create(acc)
		if err != nil {
			log.Fatalf("fail to create account resource: %s", err)
		}
	}

	resAcc, err := a.ListBy(expAccList[2].ID)
	if err != nil {
		log.Fatalf("fail to list all accounts: %s", err)
	}

	actAcc := resAcc.(*model.Account)

	assert.EqualValues(t, expAccList[2].ID, actAcc.ID)
	assert.EqualValues(t, expAccList[2].Type, actAcc.Type)
	assert.EqualValues(t, expAccList[2].OrganisationID, actAcc.OrganisationID)
	assert.EqualValues(t, expAccList[2].Attributes, actAcc.Attributes)

	for _, acc := range expAccList {
		deleteAccount(a, acc)
	}
}

func TestFailAccountListingByInvalidId(t *testing.T) {
	acc := readFileAsAccount("data/account.json")

	a := service.Account{}

	_, err := a.Create(acc)
	if err != nil {
		log.Fatalf("fail to create first account resource: %s", err)
	}

	tempID := acc.ID
	acc.ID = "wrong id value"
	_, actErr := a.ListBy(acc.ID)

	expErr := errors.ResponseError{
		StatusCode: 400,
		Message:    "request error with different status code, expected: 200 but returned: 400 with error message: id is not a valid uuid",
		CausedBy:   nil,
	}

	assert.EqualValues(t, expErr, actErr)

	acc.ID = tempID
	deleteAccount(a, acc)
}
