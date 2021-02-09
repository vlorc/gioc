package dependency

import (
	"github.com/vlorc/gioc/factory"
	"github.com/vlorc/gioc/types"
	"github.com/vlorc/gioc/utils"
	"reflect"
)

func lazyWrapper(typ reflect.Type) func(types.BeanFactory) types.BeanFactory {
	return func(b types.BeanFactory) types.BeanFactory {
		return factory.NewFuncFactory(func(provider types.Provider) (interface{}, error) {
			proxy := utils.LazyProxy(func([]reflect.Value) []reflect.Value {
				instance, err := b.Instance(provider)
				if nil != err {
					panic(err)
				}
				return []reflect.Value{utils.Convert(reflect.ValueOf(instance), typ)}
			})
			return reflect.MakeFunc(typ, proxy).Interface(), nil
		})
	}
}

func extendWrapper(dependency types.Dependency, typ reflect.Type) func(types.BeanFactory) types.BeanFactory {
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

func resolveWrapper(typ reflect.Type, names ...types.StringFactory) func(types.BeanFactory) types.BeanFactory {
	return func(b types.BeanFactory) types.BeanFactory {
		return factory.NewResolveFactory(typ, names...)
	}
}
