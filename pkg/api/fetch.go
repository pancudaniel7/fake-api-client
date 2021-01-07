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

type resourcePromise struct {
	wg  sync.WaitGroup
	res Resource
	err error
}

func NewResourcePromise(f func() (Resource, error)) *resourcePromise {
	p := &resourcePromise{}

	p.wg.Add(1)
	go func() {
		p.res, p.err = f()
		p.wg.Done()
	}()
	return p
}

func (p *resourcePromise) Then(r func(res Resource)) *resourcePromise {
	go func() {
		p.wg.Wait()
		if p.err == nil {
			r(p.res)
		}
	}()
	return p
}

func (p *resourcePromise) Cache(e func(err error)) {
	go func() {
		p.wg.Wait()
		if p.err != nil {
			e(p.err)
		}
	}()
}
