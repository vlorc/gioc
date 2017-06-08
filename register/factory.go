// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package register

import (
	"github.com/vlorc/gioc/types"
	"reflect"
)

func (rf *CoreRegisterFactory) Instance(factory types.BinderFactory) (types.Register, error) {
	return &CoreRegister{
		table:   make(map[reflect.Type]types.Binder),
		factory: factory,
	}, nil
}

func NewRegisterFactory() types.RegisterFactory {
	return &CoreRegisterFactory{}
}
