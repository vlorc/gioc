// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package register

import "github.com/vlorc/gioc/types"

var errReadonly = types.NewError(types.ErrNotRegister)

func (r ReadOnlyRegister) RegisterMapper(types.Mapper, interface{}) error {
	return errReadonly
}

func (r ReadOnlyRegister) RegisterBinder(types.Binder, interface{}) error {
	return errReadonly
}

func (r ReadOnlyRegister) RegisterInterface(interface{}, ...string) error {
	return errReadonly
}

func (r ReadOnlyRegister) RegisterInstance(interface{}, ...string) error {
	return errReadonly
}

func (r ReadOnlyRegister) RegisterPointer(interface{}, ...string) error {
	return errReadonly
}

func (r ReadOnlyRegister) RegisterFactory(types.BeanFactory, interface{}, ...string) error {
	return errReadonly
}

func (r ReadOnlyRegister) RegisterMethod(types.BeanFactory, interface{}, interface{}, ...string) error {
	return errReadonly
}
