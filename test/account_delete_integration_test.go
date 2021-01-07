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

func TestAccountDelete(t *testing.T) {
	acc := readFileAsAccount("data/account.json")

	_, err := acc.Create()
	if err != nil {
		log.Fatalf("fail to create first account resource: %s", err)
	}

	err = acc.Delete()
	assert.Nil(t, err)
}

func TestFailAccountDeleteWithInvalidId(t *testing.T) {
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
