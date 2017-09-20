// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package selector

import (
	"fmt"
	"github.com/vlorc/gioc/binder"
	"github.com/vlorc/gioc/mapper"
	"github.com/vlorc/gioc/types"
	"reflect"
)

func (tns *TypeNameSelector) MapperOf(typ reflect.Type) types.Mapper {
	return mapper.NewSelectorMapping(typ, tns)
}

func (tns *TypeNameSelector) BinderOf(typ reflect.Type) (bind types.Binder) {
	return binder.NewSelectorBinder(typ, tns)
}

func (tns *TypeNameSelector) AsBinder(typ reflect.Type) types.Binder {
	return tns.BinderOf(typ)
}

func (tns *TypeNameSelector) AsMapper(typ reflect.Type) types.Mapper {
	return tns.MapperOf(typ)
}

func (tns *TypeNameSelector) FactoryOf(typ reflect.Type, name string) (factory types.BeanFactory) {
	tns.lock.RLock()
	factory = tns.table[TypeName{Type: typ, Name: name}]
	tns.lock.RUnlock()
	return
}

func (tns *TypeNameSelector) SetBinder(typ reflect.Type, binder types.Binder) error {
	return fmt.Errorf("can't support SetBinder")
}

func (tns *TypeNameSelector) SetFactory(typ reflect.Type, name string, factory types.BeanFactory) error {
	tns.lock.Lock()
	tns.table[TypeName{Type: typ, Name: name}] = factory
	tns.lock.Unlock()
	return nil
}
