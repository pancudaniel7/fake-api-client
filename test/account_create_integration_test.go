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
	"github.com/pancudaniel7/fake-api-client/pkg/errors"
	"github.com/pancudaniel7/fake-api-client/pkg/model"
	"github.com/pancudaniel7/fake-api-client/pkg/service"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestAccountCreation(t *testing.T) {

	expAcc := readFileAsAccount("data/account.json")
	a := service.Account{}

	resResource, err := a.Create(expAcc)
	if err != nil {
		log.Fatalf("fail to create account resource: %s", err)
	}

	actAcc := resResource.(*model.Account)

	assert.EqualValues(t, expAcc.ID, actAcc.ID)
	assert.EqualValues(t, expAcc.Type, actAcc.Type)
	assert.EqualValues(t, expAcc.OrganisationID, actAcc.OrganisationID)
	assert.EqualValues(t, expAcc.Attributes, actAcc.Attributes)

	deleteAccount(a, expAcc)
}

func TestFailAccountCreationForSameId(t *testing.T) {
	acc := readFileAsAccount("data/account.json")
	a := service.Account{}

	_, actErr := a.Create(acc)
	if actErr != nil {
		log.Fatalf("fail to create first account resource: %s", actErr)
	}

	expErr := errors.ResponseError{
		StatusCode: 409,
		Message:    "request error with different status code, expected: 201 but returned: 409 with error message: Account cannot be created as it violates a duplicate constraint",
		CausedBy:   nil,
	}

	_, actErr = a.Create(acc)

	assert.NotNil(t, actErr)
	assert.EqualValues(t, expErr, actErr)

	deleteAccount(a, acc)
}

func TestFailAccountCreationWithAllInvalidValues(t *testing.T) {
	acc := readFileAsAccount("data/invalid-account.json")
	a := service.Account{}

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
			"in body must be of type uuid: \"invalid uuid\"\norganisation_id in body must be of type uuid: \"invalid organisation id\"\ntype " +
			"in body should be one of [accounts]",
		CausedBy: nil,
	}

	_, actArr := a.Create(acc)

	assert.EqualValues(t, expErr, actArr)
}

func readFileAsAccount(path string) model.Account {
	accJsonBytes := readFileAsBytes(path)

	acc := model.Account{}
	if err := json.Unmarshal(accJsonBytes, &acc); err != nil {
		log.Fatalf("Fail to unmarshal account json file bytes: %s", err)
	}
	return acc
}
