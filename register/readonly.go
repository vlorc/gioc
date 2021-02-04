// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package register

import (
	"github.com/vlorc/gioc/selector"
	"github.com/vlorc/gioc/types"
)

var errReadonly = types.NewError(types.ErrNotRegister)

func (r *readOnlyRegister) RegisterInterface(interface{}, ...string) error {
	return errReadonly
}

func (r *readOnlyRegister) RegisterInstance(interface{}, ...string) error {
	return errReadonly
}

func (r *readOnlyRegister) RegisterPointer(interface{}, ...string) error {
	return errReadonly
}

func (r *readOnlyRegister) RegisterFactory(types.BeanFactory, interface{}, ...string) error {
	return errReadonly
}

func (r *readOnlyRegister) RegisterMethod(types.BeanFactory, interface{}, interface{}, ...string) error {
	return errReadonly
}

func (r *readOnlyRegister) AsSelector() types.Selector {
	return selector.NewReadOnlyFactory(r.getter)
}
