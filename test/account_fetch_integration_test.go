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

func TestAccountFetchSuccessCreation(t *testing.T) {
	expAcc := readFileAsAccount("data/account.json")

	doneChan := make(chan int)
	p := api.NewResourcePromise(expAcc.Create)

	p.Then(func(res api.Resource) {
		actAcc := res.(*api.Account)

		assert.EqualValues(t, expAcc.ID, actAcc.ID)
		assert.EqualValues(t, expAcc.Type, actAcc.Type)
		assert.EqualValues(t, expAcc.OrganisationID, actAcc.OrganisationID)
		assert.EqualValues(t, expAcc.Attributes, actAcc.Attributes)

		doneChan <- 1
	})

	p.Cache(func(err error) {
		log.Fatalf("Fail to fetch account creation: %s", err)
	})
	<-doneChan

	deleteAccount(expAcc)
}

func TestAccountFetchFailCreation(t *testing.T) {
	expAcc := readFileAsAccount("data/invalid-account.json")

	doneChan := make(chan int)
	p := api.NewResourcePromise(expAcc.Create)

	p.Then(func(res api.Resource) {
		log.Fatalf("Should fail to fetch account creation: %s", res)
	})

	p.Cache(func(err error) {
		assert.EqualValues(t, err.(errors.ResponseError).StatusCode, 400)
		doneChan <- 1
	})
	<-doneChan
}
