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

func NewSelectorBinder(typ reflect.Type,selector types.Selector) types.Binder {
	return &SelectorBind{
		typ: typ,
		selector: selector,
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
