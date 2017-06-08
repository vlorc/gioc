// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package register

import (
	"github.com/vlorc/gioc/binder"
	"github.com/vlorc/gioc/factory"
	"github.com/vlorc/gioc/types"
	"github.com/vlorc/gioc/utils"
	"reflect"
)

func (r *CoreRegister) MapperOf(typ reflect.Type) (m types.Mapper) {
	if b := r.BinderOf(typ); nil != b {
		m = b.AsMapper()
	}
	return
}

func (r *CoreRegister) BinderOf(t reflect.Type) (bind types.Binder) {
	r.m.RLock()
	bind = r.table[t]
	r.m.RUnlock()

	return bind
}

func (r *CoreRegister) AsBinder(t reflect.Type) types.Binder {
	r.m.RLock()
	bind, ok := r.table[t]
	r.m.RUnlock()

	if !ok {
		bind, _ = r.factory.Instance(t)
		r.m.Lock()
		r.table[t] = bind
		r.m.Unlock()
	}
	return bind
}

func (r *CoreRegister) AsMapper(typ reflect.Type) (m types.Mapper) {
	if b := r.AsBinder(typ); nil != b {
		m = b.AsMapper()
	}
	return
}

func (r *CoreRegister) RegisterMapper(mapping types.Mapper, impType interface{}) error {
	return r.RegisterMapper(binder.NewProxyBinder(mapping, nil), impType)
}

func (r *CoreRegister) RegisterBinder(bind types.Binder, impType interface{}) error {
	typ := utils.TypeOf(impType)

	r.m.Lock()
	r.table[typ] = bind
	r.m.Unlock()
	return nil
}

func (r *CoreRegister) RegisterInterface(instance interface{}, args ...string) error {
	val := utils.DirectlyValue(utils.ValueOf(instance))
	if reflect.Interface != val.Kind() {
		return types.NewError(types.ErrTypeNotInterface, instance)
	}

	return r.registerFactory(
		factory.NewValueFactory(val.Interface()),
		val.Type(),
		args...)
}

func (r *CoreRegister) RegisterInstance(instance interface{}, args ...string) error {
	return r.registerFactory(
		factory.NewValueFactory(instance),
		reflect.TypeOf(instance),
		args...)
}

func (r *CoreRegister) RegisterPointer(pointer interface{}, args ...string) error {
	srcValue := reflect.ValueOf(pointer)
	if reflect.Ptr != srcValue.Kind() {
		return types.NewError(types.ErrTypeNotPointer, srcValue, args...)
	}

	srcValue = srcValue.Elem()

	return r.registerFactory(
		factory.NewPointerFactory(srcValue),
		srcValue.Type(),
		args...)
}

func (r *CoreRegister) registerFactory(beanFactory types.BeanFactory, impType reflect.Type, args ...string) error {
	b := r.AsBinder(impType)
	var name string = types.DEFAULT_NAME
	if len(args) > 0 {
		name = args[0]
	}
	return b.Bind(name, beanFactory)
}

func (r *CoreRegister) RegisterFactory(beanFactory types.BeanFactory, impType interface{}, args ...string) error {
	return r.registerFactory(beanFactory, utils.TypeOf(impType), args...)
}

func (r *CoreRegister) RegisterMethod(paramFactory types.BeanFactory, funcType interface{}, impType interface{}, args ...string) error {
	beanFactory, srcType, err := factory.NewMethodFactory(funcType, paramFactory)
	if nil != err {
		return err
	}

	dstType := srcType
	if nil != impType {
		dstType = utils.TypeOf(impType)
		if reflect.Interface == dstType.Kind() && !srcType.Implements(dstType) {
			return types.NewError(types.ErrTypeImplements, srcType, args...)
		}
		if srcType != dstType {
			return types.NewError(types.ErrTypeNotMatch, srcType, args...)
		}
	}

	return r.registerFactory(
		beanFactory,
		dstType,
		args...)
}
