// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package register

import (
	"github.com/vlorc/gioc/factory"
	"github.com/vlorc/gioc/types"
	"github.com/vlorc/gioc/utils"
	"reflect"
)

func (r *coreRegister) RegisterInterface(instance interface{}, args ...string) error {
	val := utils.IndirectValue(utils.ValueOf(instance))
	if reflect.Interface != val.Kind() || val.IsNil() {
		return types.NewWithError(types.ErrTypeNotInterface, instance)
	}

	return r.registerFactory(
		factory.NewValueFactory(val.Interface()),
		val.Type(),
		args...)
}

func (r *coreRegister) RegisterInstance(instance interface{}, args ...string) error {
	return r.registerFactory(
		factory.NewValueFactory(instance),
		reflect.TypeOf(instance),
		args...)
}

func (r *coreRegister) RegisterPointer(pointer interface{}, args ...string) error {
	srcValue := reflect.ValueOf(pointer)
	if reflect.Ptr != srcValue.Kind() {
		return types.NewWithError(types.ErrTypeNotPointer, srcValue, args...)
	}

	srcValue = srcValue.Elem()
	return r.registerFactory(
		factory.NewPointerFactory(srcValue),
		srcValue.Type(),
		args...)
}

func (r *coreRegister) RegisterFactory(beanFactory types.BeanFactory, impType interface{}, args ...string) error {
	return r.registerFactory(
		beanFactory,
		utils.TypeOf(impType),
		args...)
}

func (r *coreRegister) RegisterMethod(paramFactory types.BeanFactory, method interface{}, impType interface{}, args ...string) error {
	beanFactory, srcType, err := factory.NewMethodFactory(method, paramFactory)
	if nil != err {
		return err
	}

	dstType := srcType
	if nil != impType {
		dstType = utils.TypeOf(impType)
		if reflect.Interface == dstType.Kind() && !srcType.Implements(dstType) {
			return types.NewWithError(types.ErrTypeImplements, srcType, args...)
		}
		if srcType != dstType {
			return types.NewWithError(types.ErrTypeNotMatch, srcType, args...)
		}
	}
	return r.registerFactory(
		beanFactory,
		dstType,
		args...)
}

func (r *coreRegister) registerFactory(beanFactory types.BeanFactory, impType reflect.Type, args ...string) error {
	var name string
	if len(args) > 0 {
		name = args[0]
	}

	r.selector.Set(impType, name, beanFactory)
	return nil
}

func (r *coreRegister) AsSelector() types.Selector {
	return r.selector
}
