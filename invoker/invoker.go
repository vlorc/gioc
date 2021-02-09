// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package invoker

import (
	"github.com/vlorc/gioc/types"
	"github.com/vlorc/gioc/utils"
	"reflect"
)

func lazyDependency(val reflect.Value) func(types.Provider) types.Dependency {
	return utils.Lazy(func(provider types.Provider) types.Dependency {
		var dependFactory types.DependencyFactory
		provider.Assign(&dependFactory)
		dep, err := dependFactory.Instance(val)
		if nil != err {
			utils.Panic(err)
		}
		return dep
	}).(func(types.Provider) types.Dependency)
}

func NewInvoker(method interface{}, dependency types.Dependency) types.Invoker {
	val := utils.ValueOf(method)
	if reflect.Func != val.Kind() {
		utils.Panic(types.NewWithError(types.ErrTypeNotFunction, method))
	}
	if val.Type().NumIn() <= 0 {
		return NoParamInvoker(val)
	}
	return newInvoker(val, dependency)
}

func newInvoker(val reflect.Value, dependency types.Dependency) types.Invoker {
	if nil == dependency {
		return &CoreInvoker{
			method:     val,
			dependency: lazyDependency(val),
		}
	}
	return &CoreInvoker{
		method: val,
		dependency: func(types.Provider) types.Dependency {
			return dependency
		},
	}
}

func (ci *CoreInvoker) Apply(args ...interface{}) ([]reflect.Value, error) {
	return ci.ApplyWith(nil, args...)
}

func (ci *CoreInvoker) ApplyWith(provider types.Provider, args ...interface{}) ([]reflect.Value, error) {
	return ci.applyWith(provider, args...)
}

func (ci *CoreInvoker) applyWith(provider types.Provider, args ...interface{}) (result []reflect.Value, err error) {
	defer utils.Recover(&err)

	dep := ci.dependency(provider)

	params := make([]reflect.Value, dep.Length())

	for scan := dep.AsScan(); scan.Next(); {
		i := scan.Index().Value()
		if i < len(args) {
			params[i] = utils.Convert(reflect.ValueOf(args[i]), scan.Type())
			continue
		}
		b := scan.Factory(provider)
		instance, err := b.Instance(provider)
		if nil != err {
			return nil, err
		}
		params[i] = utils.Convert(reflect.ValueOf(instance), scan.Type())
	}

	result = ci.method.Call(params)
	return
}
