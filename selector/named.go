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

func (ns *NamedSelector) MapperOf(typ reflect.Type) types.Mapper {
	return mapper.NewSelectorMapping(typ, ns)
}

func (ns *NamedSelector) BinderOf(typ reflect.Type) (bind types.Binder) {
	return binder.NewSelectorBinder(typ, ns)
}

func (ns *NamedSelector) AsBinder(typ reflect.Type) types.Binder {
	return ns.BinderOf(typ)
}

func (ns *NamedSelector) AsMapper(typ reflect.Type) types.Mapper {
	return ns.MapperOf(typ)
}

func (ns *NamedSelector) FactoryOf(typ reflect.Type, name string) types.BeanFactory{
	ns.lock.RLock()
	factory := ns.selector.FactoryOf(typ,name)
	ns.lock.RUnlock()
	return factory
}

func (ns *NamedSelector) SetBinder(typ reflect.Type, binder types.Binder) error {
	return fmt.Errorf("can't support SetBinder")
}

func (ns *NamedSelector) SetFactory(typ reflect.Type, name string, factory types.BeanFactory) error {
	ns.lock.Lock()
	err := ns.selector.SetFactory(typ,name,factory)
	ns.lock.Unlock()
	return err
}

func (tns typeNameSelector) FactoryOf(typ reflect.Type, name string) types.BeanFactory {
	return tns[typeName{typ: typ, name: name}]
}

func (tns typeNameSelector) SetFactory(typ reflect.Type, name string, factory types.BeanFactory) error {
	tns[typeName{typ: typ, name: name}] = factory
	return nil
}

func (ns nameSelector) FactoryOf(typ reflect.Type, name string) types.BeanFactory {
	return ns[name].factory
}

func (ns nameSelector) SetFactory(typ reflect.Type, name string, factory types.BeanFactory) error {
	ns[name] = typeFactory{typ: typ,factory: factory}
	return nil
}