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
