// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package binder

import (
	"github.com/vlorc/gioc/types"
	"reflect"
)

func (bf *CoreBinderFactory) Instance(p reflect.Type) (types.Binder, error) {
	return NewNameBinder(), nil
}

func NewBinderFactory() types.BinderFactory {
	return &CoreBinderFactory{}
}

func NewNameBinder() types.Binder {
	return &NamedBind{
		table: make(map[string]types.BeanFactory),
	}
}

func NewProxyBinder(read types.Mapper, write types.Binder) types.Binder {
	if nil == read {
		read = write.AsMapper()
	}
	return &ProxyBind{
		read:  read,
		write: write,
	}
}
