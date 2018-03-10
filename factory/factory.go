// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package factory

import (
	"github.com/vlorc/gioc/types"
	"github.com/vlorc/gioc/utils"
	"reflect"
)

func NewValueFactory(value interface{}, err ...error) types.BeanFactory {
	obj := &ValueFactory{
		value: value,
	}
	if len(err) > 0 {
		obj.err = err[0]
	}
	return obj
}

func NewMethodFactory(impType interface{}, paramFactory types.BeanFactory, index ...int) (types.BeanFactory, reflect.Type, error) {
	return methodFactoryOf(impType, paramFactory, index...)
}

func NewProxyFactory(value types.BeanFactory) types.BeanFactory {
	return &ProxyFactory{
		factory: value,
	}
}

func NewMutexFactory(value types.BeanFactory) types.BeanFactory {
	return &MutexFactory{
		factory: value,
	}
}

func NewTypeFactory(value interface{}) types.BeanFactory {
	return &TypeFactory{
		typ: utils.TypeOf(value),
	}
}

func NewSingleFactory(value types.BeanFactory) types.BeanFactory {
	return singleFactoryOf(value)
}

func NewPointerFactory(value reflect.Value) types.BeanFactory {
	return &PointerFactory{
		value: utils.DirectlyValue(value),
	}
}

func NewFuncFactory(value func(types.Provider) (interface{}, error)) types.BeanFactory {
	return FuncFactory(value)
}

func NewParamFactory(length int) types.BeanFactory {
	return ParamFactory(length)
}

func NewExportFactory(factory types.BeanFactory, provider func() types.Provider) types.BeanFactory {
	return &ExportFactory{
		provider: provider,
		factory:  factory,
	}
}

func NewConvertFactory(factory types.BeanFactory, typ reflect.Type) types.BeanFactory {
	return &ConvertFactory{
		factory: factory,
		typ:     typ,
	}
}
