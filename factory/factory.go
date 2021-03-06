// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package factory

import (
	"errors"
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
	return &newFactory{
		typ: utils.TypeOf(value),
	}
}

func NewSingleFactory(value types.BeanFactory) types.BeanFactory {
	return singleFactoryOf(value)
}

func NewPointerFactory(value reflect.Value) types.BeanFactory {
	return &PointerFactory{
		value: utils.IndirectValue(value),
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

func NewResolveFactory(typ reflect.Type, name ...types.StringFactory) types.BeanFactory {
	f := &resolveFactory{typ: typ}
	if len(name) > 0 {
		f.name = make([]types.StringFactory, len(name))
		copy(f.name, name)
	}
	return f
}

func NewChainFactory(factory ...types.BeanFactory) types.BeanFactory {
	if len(factory) == 0 {
		utils.Panic(errors.New("factory is empty"))
	}
	if len(factory) == 1 {
		return factory[0]
	}

	chain := make(chainFactory, len(factory))
	copy(chain, factory)
	return chain
}

func NewDependencyFactory(factory types.BeanFactory, dependency types.Dependency, after ...func(interface{}) interface{}) types.BeanFactory {
	f := &DependencyFactory{
		factory:    factory,
		dependency: dependency,
	}
	if len(after) > 0 {
		f.after = make([]func(interface{}) interface{}, len(after))
		copy(f.after, after)
	}
	return f
}

func NewSliceFactory(typ reflect.Type, name ...types.StringFactory) types.BeanFactory {
	if len(name) == 0 {
		return &resolveAnyFactory{typ: typ}
	}

	dump := make([]types.StringFactory, len(name))
	copy(dump, name)

	if types.BeanFactoryType == typ {
		return &resolveGeneralFactory{
			typ: typ,
			append: func(value reflect.Value, factory types.GeneralFactory) reflect.Value {
				return reflect.Append(value, reflect.ValueOf(factory.(types.BeanFactory)))
			},
		}
	}

	if types.BeanFactoryType == typ {
		return &resolveGeneralFactory{
			typ: typ,
			append: func(value reflect.Value, factory types.GeneralFactory) reflect.Value {
				return reflect.Append(value, reflect.ValueOf(factory))
			},
		}
	}

	return &resolveNamesFactory{typ: typ, name: dump}
}

func NewLazyFactory(typ reflect.Type, factory types.BeanFactory) types.BeanFactory {
	return newRequestFactory(typ, factory, utils.LazyProxy)
}

func NewRequestFactory(typ reflect.Type, factory types.BeanFactory) types.BeanFactory {
	return newRequestFactory(typ, factory, func(v func([]reflect.Value) []reflect.Value) func([]reflect.Value) []reflect.Value {
		return v
	})
}

func NewMakeFactory(typ reflect.Type, length ...int) types.BeanFactory {
	switch typ.Kind() {
	case reflect.Slice:
		if len(length) > 0 {
			return &makeSliceFactory{typ: typ, length: length[0]}
		}
		return &makeSliceFactory{typ: typ}
	case reflect.Chan:
		if len(length) > 0 {
			return &makeChanFactory{typ: typ, length: length[0]}
		}
		return &makeChanFactory{typ: typ}
	case reflect.Map:
		return &makeMapFactory{typ: typ}
	default:
		return NewValueFactory(nil, types.NewError(types.ErrTypeNotSupport, typ))
	}
}
