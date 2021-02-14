// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package factory

import (
	"github.com/vlorc/gioc/types"
	"reflect"
)

func makeInstance(retIndex, errIndex int) func([]reflect.Value) (interface{}, error) {
	if errIndex > 0 {
		return func(v []reflect.Value) (instance interface{}, err error) {
			instance = v[retIndex].Interface()
			if !v[errIndex].IsNil() {
				err = v[errIndex].Interface().(error)
			}
			return
		}
	}
	return func(v []reflect.Value) (interface{}, error) {
		return v[retIndex].Interface(), nil
	}
}

func makeParam(typ reflect.Type, paramFactory types.BeanFactory) func(types.Provider) ([]reflect.Value, error) {
	if nil != paramFactory {
		return func(provider types.Provider) ([]reflect.Value, error) {
			instance, err := paramFactory.Instance(provider)
			if nil != err {
				return nil, err
			}
			return instance.([]reflect.Value), nil
		}
	}
	if 1 == typ.NumIn() && typ.In(0) == types.ProviderType {
		return func(provider types.Provider) ([]reflect.Value, error) {
			return []reflect.Value{reflect.ValueOf(&provider).Elem()}, nil
		}
	}
	return func(types.Provider) ([]reflect.Value, error) {
		return nil, nil
	}
}

func methodFactoryOf(
	impType interface{},
	paramFactory types.BeanFactory,
	index ...int) (factory types.BeanFactory, resultType reflect.Type, err error) {
	srcType := reflect.TypeOf(impType)
	if reflect.Func != srcType.Kind() {
		err = types.NewWithError(types.ErrTypeNotFunction, srcType)
		return
	}
	if srcType.NumOut() <= 0 {
		err = types.NewWithError(types.ErrTypeNotInterface, srcType)
		return
	}

	retIndex := 0
	if len(index) > 0 {
		if index[0] < 0 || index[0] > srcType.NumOut() {
			err = types.NewWithError(types.ErrTypeNotInterface, srcType)
			return
		}
		retIndex = index[0]
	}

	errIndex := -1
	for i := srcType.NumOut() - 1; i >= 0; i-- {
		if srcType.Out(i) == types.ErrorType {
			errIndex = i
			break
		}
	}

	factory = newCallFactory(reflect.ValueOf(impType), makeParam(srcType, paramFactory), makeInstance(retIndex, errIndex))
	resultType = srcType.Out(retIndex)

	return
}

func newCallFactory(
	method reflect.Value,
	makeParam func(types.Provider) ([]reflect.Value, error),
	makeInstance func([]reflect.Value) (interface{}, error)) types.BeanFactory {
	return NewFuncFactory(func(provider types.Provider) (instance interface{}, err error) {
		param, err := makeParam(provider)
		if nil == err {
			result := method.Call(param)
			instance, err = makeInstance(result)
		}
		return
	})
}
