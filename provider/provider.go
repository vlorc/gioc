// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package provider

import (
	"github.com/vlorc/gioc/types"
	"github.com/vlorc/gioc/utils"
	"reflect"
)

func (p *CoreProvider) Resolve(impType interface{}, args ...string) (interface{}, error) {
	var name string = types.DEFAULT_NAME
	if len(args) > 0 {
		name = args[0]
	}
	return p.ResolveNamed(impType, name, -1)
}

func (p *CoreProvider) ResolveNamed(impType interface{}, name string, deep int) (interface{}, error) {
	return p.ResolveType(utils.TypeOf(impType), name, deep)
}

func (p *CoreProvider) ResolveType(typ reflect.Type, name string, deep int) (instance interface{}, err error) {
	if factory := p.selector.FactoryOf(typ, name); nil != factory {
		instance, err = factory.Instance(p)
		return
	}
	if nil == p.parent || 0 == deep {
		err = types.NewWithError(types.ErrFactoryNotFound, typ, name)
	}

	instance, err = p.parent.ResolveType(typ, name, deep - 1)
	return
}

func (p *CoreProvider) Assign(dst interface{}, args ...string) error {
	var name string = types.DEFAULT_NAME
	if len(args) > 0 {
		name = args[0]
	}
	return p.AssignType(utils.ValueOf(dst), nil, name, -1)
}

func (p *CoreProvider) AssignNamed(dst interface{}, impType interface{}, name string, deep int) (err error) {
	defer utils.Recover(&err)

	dstValue := utils.ValueOf(dst)
	var srcType reflect.Type
	if nil != impType{
		srcType = utils.TypeOf(impType)
	}
	return p.AssignType(dstValue, srcType,name,deep)
}

func (p *CoreProvider) AssignType(dstValue reflect.Value, srcType reflect.Type, name string, deep int) (err error) {
	defer utils.Recover(&err)

	if !dstValue.CanSet() {
		if reflect.Ptr != dstValue.Kind() {
			err = types.NewWithError(types.ErrTypeNotSet, dstValue)
			return
		}
		dstValue = dstValue.Elem()
	}
	if nil == srcType{
		srcType = dstValue.Type()
	}

	instance, err := p.ResolveType(srcType, name, deep)
	if nil != err {
		return err
	}

	srcValue := utils.Convert(reflect.ValueOf(instance),srcType)
	dstValue.Set(srcValue)
	return
}
