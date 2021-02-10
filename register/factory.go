// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package register

import (
	"github.com/vlorc/gioc/types"
)

func (*coreRegisterFactory) Instance(selector types.Selector) (types.Register, error) {
	return NewRegister(selector), nil
}

func NewRegisterFactory() types.RegisterFactory {
	return &coreRegisterFactory{}
}

func NewRegister(selector types.Selector) types.Register {
	return &coreRegister{
		selector: selector,
	}
}

func NewReadOnlyRegister(getter types.SelectorGetter) types.Register {
	return &readOnlyRegister{getter}
}
