package builder

import (
	"reflect"
	"github.com/vlorc/gioc/types"
	"sync"
	"github.com/vlorc/gioc/utils"
)

func MakeLazyLoad(dstVal reflect.Value,load func()) {
	once := sync.Once{}
	srcVal := reflect.MakeFunc(dstVal.Type(),func(args []reflect.Value) []reflect.Value{
		once.Do(load)
		return dstVal.Call(args)
	})
	dstVal.Set(srcVal)
}

func MakeLazyInstance(val reflect.Value,provider types.Provider,des types.DescriptorGetter) {
	MakeLazyLoad(val, func() {
		instance, err := provider.ResolveType(val.Type().Out(0), des.Name(),-1)
		if nil != err {
			panic(err)
		}
		results := []reflect.Value{reflect.ValueOf(instance)}
		val.Set(reflect.MakeFunc(val.Type(),func([]reflect.Value) []reflect.Value{
			return results
		}))
	})
}

func MakeLazyExtends(val reflect.Value,provider types.Provider,des types.DescriptorGetter) {
	MakeLazyLoad(val, func() {
		dstVal := reflect.New(des.Type().Out(0)).Elem()
		FullAllInstance(&BuildContext{
			Provider: provider,
			Inject: des.Depend().AsInject(utils.NewOf(dstVal)),
		})
		results := []reflect.Value{dstVal}
		val.Set(reflect.MakeFunc(val.Type(),func([]reflect.Value) []reflect.Value{
			return results
		}))
	})
}