// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package invoker

import (
	"github.com/vlorc/gioc/builder"
	"github.com/vlorc/gioc/factory"
	"github.com/vlorc/gioc/types"
	"github.com/vlorc/gioc/utils"
	"reflect"
)

func lazyBuilder(val reflect.Value) func(types.Provider) types.Builder {
	return utils.Lazy(func(provider types.Provider) types.Builder {
		var dependFactory types.DependencyFactory
		provider.Assign(&dependFactory)
		dep, err := dependFactory.Instance(val)
		if nil != err {
			panic(err)
		}
		return builder.NewBuilder(factory.NewParamFactory(dep.Length()), dep)
	}).(func(types.Provider) types.Builder)
}

func NewInvoker(method interface{}, builder types.Builder) types.Invoker {
	val := utils.ValueOf(method)
	if reflect.Func != val.Kind() {
		panic(types.NewWithError(types.ErrTypeNotFunction, method))
	}
	if val.Type().NumIn() <= 0 {
		return NoParamInvoker(val)
	}
	return newInvoker(val, builder)
}

func newInvoker(val reflect.Value, builder types.Builder) types.Invoker {
	if nil == builder {
		return &CoreInvoker{
			method:  val,
			builder: lazyBuilder(val),
		}
	}
	return &CoreInvoker{
		method: val,
		builder: func(types.Provider) types.Builder {
			return builder
		},
	}
}

func (i *CoreInvoker) Apply(args ...interface{}) []reflect.Value {
	return i.ApplyWith(nil, args...)
}

func (i *CoreInvoker) ApplyWith(provider types.Provider, args ...interface{}) []reflect.Value {
	temp, err := i.builder(provider).Build(provider, func(ctx *types.BuildContext) {
		if len(args) > 0 {
			ctx.FullBefore = builder.IndexFullBefore(args)
		}
	})
	if nil != err {
		panic(err)
	}
	return i.method.Call(temp.([]reflect.Value))
}
