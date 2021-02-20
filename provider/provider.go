// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package provider

import (
	"github.com/vlorc/gioc/types"
	"github.com/vlorc/gioc/utils"
	"reflect"
)

func (p *coreProvider) Get(impType interface{}, args ...string) (interface{}, error) {
	var name string
	if len(args) > 0 {
		name = args[0]
	}
	return p.Instance(utils.TypeOf(impType), name)
}

func (p *coreProvider) Instance(typ reflect.Type, name string) (interface{}, error) {
	if b := p.Factory(typ, name, -1); nil != b {
		return b.Instance(p)
	}
	return nil, types.NewWithError(types.ErrFactoryNotFound, typ, name)
}

func (p *coreProvider) Load(receive interface{}, args ...string) error {
	var name string
	if len(args) > 0 {
		name = args[0]
	}
	return p.load(receive, nil, name)
}

func (p *coreProvider) load(receive interface{}, typ reflect.Type, name string) (err error) {
	defer utils.Recover(&err)

	val := utils.ValueOf(receive)

	if !val.CanSet() {
		if reflect.Ptr != val.Kind() {
			err = types.NewWithError(types.ErrTypeNotSet, receive)
			return
		}
		val = val.Elem()
	}
	if nil == typ {
		typ = val.Type()
	}

	instance, err := p.Instance(typ, name)
	if nil == err {
		v := utils.Convert(reflect.ValueOf(instance), typ)
		val.Set(v)
	}
	return
}

func (p *coreProvider) Factory(typ reflect.Type, name string, deep int) types.GeneralFactory {
	if factory := p.selector.Get(typ, name); nil != factory {
		return factory
	}
	if nil != p.parent && 0 != deep {
		return p.parent.Factory(typ, name, deep-1)
	}
	return nil
}

func (p *coreProvider) Range(callback func(types.GeneralFactory) bool, typ ...reflect.Type) bool {
	if p.selector.Range(callback, typ...) {
		if nil != p.parent {
			return p.parent.Range(callback, typ...)
		}
		return true
	}
	return false
}

func (p *coreProvider) Selector() types.SelectorGetter {
	return p.selector
}
