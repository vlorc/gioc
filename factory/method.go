// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package factory

import (
	"github.com/vlorc/gioc/types"
	"reflect"
)

func methodFactoryOf(
	impType interface{},
	paramFactory types.BeanFactory,
	index ...int,
) (factory types.BeanFactory, resultType reflect.Type, err error) {
	srcType := reflect.TypeOf(impType)
	if reflect.Func != srcType.Kind() {
		return
	}
	if srcType.NumOut() <= 0 {
		return
	}

	retIndex := 0
	if len(index) > 0 {
		if index[0] < 0 || index[0] > srcType.NumOut() {
			return
		}
		retIndex = index[0]
	}

	var makeParam func(types.Provider) ([]reflect.Value, error)
	var makeInstance func([]reflect.Value) (interface{}, error)

	errIndex := -1
	for i := srcType.NumOut() - 1; i >= 0; i-- {
		if srcType.Out(i).Implements(types.ErrorType) {
			errIndex = i
			break
		}
	}
	if errIndex > 0 {
		makeInstance = func(v []reflect.Value) (instance interface{}, err error) {
			instance = v[retIndex].Interface()
			if !v[errIndex].IsNil() {
				err = v[errIndex].Interface().(error)
			}
			return
		}
	} else {
		makeInstance = func(v []reflect.Value) (interface{}, error) {
			return v[retIndex].Interface(), nil
		}
	}

	if nil != paramFactory {
		makeParam = func(provider types.Provider) ([]reflect.Value, error) {
			instance, err := paramFactory.Instance(provider)
			return instance.([]reflect.Value), err
		}
	} else if 1 == srcType.NumIn() && srcType.In(0) == types.ProviderType {
		makeParam = func(provider types.Provider) ([]reflect.Value, error) {
			return []reflect.Value{reflect.ValueOf(&provider).Elem()}, nil
		}
	} else {
		makeParam = func(types.Provider) ([]reflect.Value, error) {
			return nil, nil
		}
	}

	factory = newCallFactory(
		reflect.ValueOf(impType),
		makeParam,
		makeInstance,
	)
	resultType = srcType.Out(retIndex)

	return
}

func newCallFactory(
	src reflect.Value,
	makeParam func(types.Provider) ([]reflect.Value, error),
	makeInstance func([]reflect.Value) (interface{}, error),
) types.BeanFactory {
	return NewFuncFactory(func(provider types.Provider) (instance interface{}, err error) {
		param, err := makeParam(provider)
		if nil == err {
			result := src.Call(param)
			instance, err = makeInstance(result)
		}
		return
	})
}
