// Copyright (c) 2016 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package rpc

import (
	"github.com/uber-go/uberfx/core"
	"github.com/uber-go/uberfx/modules"

	"go.uber.org/yarpc/encoding/json"
)

type CreateJsonRegistrantsFunc func(service core.ServiceHost) []json.Registrant

func JsonModule(hookup CreateJsonRegistrantsFunc, options ...modules.ModuleOption) core.ModuleCreateFunc {
	return func(mi core.ModuleCreateInfo) ([]core.Module, error) {
		if mod, err := newYarpcJsonModule(mi, hookup, options...); err == nil {
			return []core.Module{mod}, nil
		} else {
			return nil, err
		}

	}
}

func newYarpcJsonModule(mi core.ModuleCreateInfo, createService CreateJsonRegistrantsFunc, options ...modules.ModuleOption) (*YarpcModule, error) {

	reg := func(mod *YarpcModule) {
		procs := createService(mi.Host)

		if procs != nil {
			for _, proc := range procs {
				json.Register(mod.rpc, proc)
			}
		}
	}

	return newYarpcModule(mi, reg, options...)
}