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
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestAccountCreation(t *testing.T) {
	acc := readFileAsAccount("data/account.json")
	resAcc := create(acc).(*api.Account)

	assert.EqualValues(t, acc.Attributes, resAcc.Attributes)
}

func create(acc api.Account) api.Resource {
	resAcc, err := acc.Create()
	if err != nil {
		log.Fatalf("Fail to create account: %s", err)
	}
	return resAcc
}

func readFileAsAccount(path string) api.Account {
	accJsonBytes := readFileAsBytes(path)

	acc := api.Account{}
	if err := json.Unmarshal(accJsonBytes, &acc); err != nil {
		log.Fatalf("Could not unmarshal acc bytes object: %s", err)
	}
	return acc
}
