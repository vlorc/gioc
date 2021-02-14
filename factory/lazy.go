package factory

import (
	"github.com/vlorc/gioc/types"
	"github.com/vlorc/gioc/utils"
	"reflect"
)

func newRequestFactory(typ reflect.Type, factory types.BeanFactory, create func(func([]reflect.Value) []reflect.Value) func([]reflect.Value) []reflect.Value) types.BeanFactory {
	return NewFuncFactory(func(provider types.Provider) (interface{}, error) {
		proxy := create(func([]reflect.Value) []reflect.Value {
			instance, err := factory.Instance(provider)
			if nil != err {
				utils.Panic(err)
			}
			return []reflect.Value{utils.Convert(reflect.ValueOf(instance), typ.Out(0))}
		})
		return reflect.MakeFunc(typ, proxy).Interface(), nil
	})
}
