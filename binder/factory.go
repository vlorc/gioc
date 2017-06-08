// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package binder

import (
	"github.com/vlorc/gioc/types"
	"reflect"
)

func (bf *CoreBinderFactory) Instance(p reflect.Type) (types.Binder, error) {
	obj := &NamedBind{
		table: make(map[string]types.BeanFactory),
	}
	//obj.mapper = NewNamedMapping(obj.table,obj.m.RLocker())
	return obj, nil
}

func NewBinderFactory() types.BinderFactory {
	return &CoreBinderFactory{}
}
