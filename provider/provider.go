// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package provider

import (
	"github.com/vlorc/gioc/types"
	"github.com/vlorc/gioc/utils"
	"reflect"
)

func (p *coreProvider) Resolve(impType interface{}, args ...string) (interface{}, error) {
	var name string
	if len(args) > 0 {
		name = args[0]
	}
	return p.ResolveNamed(impType, name, -1)
}

func (p *coreProvider) ResolveNamed(impType interface{}, name string, deep int) (interface{}, error) {
	return p.ResolveType(utils.TypeOf(impType), name, deep)
}

func (p *coreProvider) ResolveType(typ reflect.Type, name string, deep int) (instance interface{}, err error) {
	if factory := p.selector.Get(typ, name); nil != factory {
		instance, err = factory.Instance(p)
		return
	}
	if nil == p.parent || 0 == deep {
		err = types.NewWithError(types.ErrFactoryNotFound, typ, name)
	} else {
		instance, err = p.parent.ResolveType(typ, name, deep-1)
	}
	return
}

func (p *coreProvider) Assign(dst interface{}, args ...string) error {
	var name string
	if len(args) > 0 {
		name = args[0]
	}
	return p.AssignType(utils.ValueOf(dst), nil, name, -1)
}

func (p *coreProvider) AssignNamed(dst interface{}, impType interface{}, name string, deep int) (err error) {
	dstValue := utils.ValueOf(dst)
	var srcType reflect.Type
	if nil != impType {
		srcType = utils.TypeOf(impType)
	}
	return p.AssignType(dstValue, srcType, name, deep)
}

func (p *coreProvider) AssignType(dstValue reflect.Value, srcType reflect.Type, name string, deep int) (err error) {
	defer utils.Recover(&err)

	if !dstValue.CanSet() {
		if reflect.Ptr != dstValue.Kind() {
			err = types.NewWithError(types.ErrTypeNotSet, dstValue)
			return
		}
		dstValue = dstValue.Elem()
	}
	if nil == srcType {
		srcType = dstValue.Type()
	}

	instance, err := p.ResolveType(srcType, name, deep)
	if nil != err {
		return err
	}

	srcValue := utils.Convert(reflect.ValueOf(instance), srcType)
	dstValue.Set(srcValue)
	return
}

func (p *coreProvider) FactoryOf(typ reflect.Type, name string, deep int) types.BeanFactory {
	if factory := p.selector.Get(typ, name); nil != factory {
		return factory
	}
	return p.parent.FactoryOf(typ, name, deep-1)
}

func (p *coreProvider) AsSelector() types.SelectorGetter {
	return p.selector
}
