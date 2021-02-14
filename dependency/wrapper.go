// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package dependency

import (
	"github.com/vlorc/gioc/factory"
	"github.com/vlorc/gioc/types"
	"reflect"
)

func lazyWrapper(typ reflect.Type) func(types.BeanFactory) types.BeanFactory {
	return func(b types.BeanFactory) types.BeanFactory {
		return factory.NewLazyFactory(typ, b)
	}
}

func extendSliceWrapper(typ reflect.Type, name ...types.StringFactory) func(types.BeanFactory) types.BeanFactory {
	return func(b types.BeanFactory) types.BeanFactory {
		return factory.NewSliceFactory(typ, name...)
	}
}

func extendStructWrapper(dependency types.Dependency, typ reflect.Type) func(types.BeanFactory) types.BeanFactory {
	return func(b types.BeanFactory) types.BeanFactory {
		return factory.NewDependencyFactory(factory.NewTypeFactory(dependency.Type()), dependency, __elem(reflect.PtrTo(dependency.Type()), typ))
	}
}

func defaultWrapper(value interface{}) func(types.BeanFactory) types.BeanFactory {
	v := factory.NewValueFactory(value)
	return func(b types.BeanFactory) types.BeanFactory {
		if nil != b {
			return factory.NewChainFactory(b, v)
		}
		return v
	}
}

func newWrapper(typ reflect.Type) func(types.BeanFactory) types.BeanFactory {
	return func(b types.BeanFactory) types.BeanFactory {
		return factory.NewTypeFactory(typ)
	}
}

func requestWrapper(typ reflect.Type) func(types.BeanFactory) types.BeanFactory {
	return func(b types.BeanFactory) types.BeanFactory {
		return factory.NewRequestFactory(typ, b)
	}
}
