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

import "sync"

// resourcePromise struct is used to keep all needed objects in order to create
// asynchronous promise functionality
type resourcePromise struct {
	wg  sync.WaitGroup
	res Resource
	err error
}

// NewResourcePromise returns a new promise and calls immediately the function f
// in a new goroutine.
// Also before the new goroutine method increments the wait group in order for being
// able to reproduce asynchronous functionality.
func NewResourcePromise(f func() (Resource, error)) *resourcePromise {
	p := &resourcePromise{}

	p.wg.Add(1)
	go func() {
		p.res, p.err = f()
		p.wg.Done()
	}()
	return p
}

// Then method is used to call a f function if the promise call was successful.
// The method waits for goroutine using wait group, and after that calls the successful
// f function if err object is nil.
func (p *resourcePromise) Then(r func(res Resource)) *resourcePromise {
	go func() {
		p.wg.Wait()
		if p.err == nil {
			r(p.res)
		}
	}()
	return p
}

// Cache method is used to call an e function if the promise call failed.
// The method waits for goroutine using wait group, and after that calls the failed
// e function if err object is not nil.
func (p *resourcePromise) Cache(e func(err error)) {
	go func() {
		p.wg.Wait()
		if p.err != nil {
			e(p.err)
		}
	}()
}
