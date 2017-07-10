// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package selector

import (
	"github.com/vlorc/gioc/types"
	"reflect"
)

func (ts *TypeSelector) MapperOf(typ reflect.Type) (m types.Mapper) {
	if b := ts.BinderOf(typ); nil != b {
		m = b.AsMapper()
	}
	return
}

func (ts *TypeSelector) BinderOf(typ reflect.Type) (bind types.Binder) {
	ts.lock.RLock()
	bind = ts.table[typ]
	ts.lock.RUnlock()

	return bind
}

func (ts *TypeSelector) AsBinder(typ reflect.Type) types.Binder {
	ts.lock.RLock()
	bind, ok := ts.table[typ]
	ts.lock.RUnlock()

	if !ok {
		bind, _ = ts.factory.Instance(typ)
		ts.lock.Lock()
		ts.table[typ] = bind
		ts.lock.Unlock()
	}
	return bind
}

func (ts *TypeSelector) AsMapper(typ reflect.Type) (m types.Mapper) {
	if b := ts.AsBinder(typ); nil != b {
		m = b.AsMapper()
	}
	return
}

func (ts *TypeSelector) FactoryOf(typ reflect.Type, name string) (factory types.BeanFactory) {
	if mapper := ts.MapperOf(typ); nil != mapper {
		factory = mapper.Resolve(name)
	}
	return
}

func (tns *TypeSelector) SetBinder(typ reflect.Type, bind types.Binder) error {
	tns.lock.Lock()
	tns.table[typ] = bind
	tns.lock.Unlock()
	return nil
}

func (ts *TypeSelector) SetFactory(typ reflect.Type, name string, factory types.BeanFactory) (err error) {
	if bind := ts.AsBinder(typ); nil != bind {
		err = bind.Bind(name, factory)
	}
	return
}
