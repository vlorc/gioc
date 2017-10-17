// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package register

import (
	"github.com/vlorc/gioc/types"
)

func (rf *CoreRegisterFactory) Instance(selector types.Selector) (types.Register, error) {
	return NewRegister(selector), nil
}

func NewRegisterFactory() types.RegisterFactory {
	return &CoreRegisterFactory{}
}

func NewRegister(selector types.Selector) types.Register {
	return &CoreRegister{
		selector: selector,
	}
}
