// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package builder

import (
	"github.com/vlorc/gioc/types"
	"github.com/vlorc/gioc/utils"
	"reflect"
	"sync"
)

func MakeLazyLoad(dstVal reflect.Value, load func()) {
	once := sync.Once{}
	srcVal := reflect.MakeFunc(dstVal.Type(), func(args []reflect.Value) []reflect.Value {
		once.Do(load)
		return dstVal.Call(args)
	})
	dstVal.Set(srcVal)
}

func MakeLazyInstance(val reflect.Value, provider types.Provider, des types.DescriptorGetter) {
	MakeLazyLoad(val, func() {
		instance, err := provider.ResolveType(des.Type().Out(0), des.Name(), -1)
		if nil != err {
			panic(err)
		}
		results := []reflect.Value{reflect.ValueOf(instance)}
		val.Set(reflect.MakeFunc(val.Type(), func([]reflect.Value) []reflect.Value {
			return results
		}))
	})
}

func MakeLazyExtends(val reflect.Value, provider types.Provider, des types.DescriptorGetter) {
	MakeLazyLoad(val, func() {
		results := []reflect.Value{reflect.New(des.Type().Out(0)).Elem()}
		buildDefault(provider, des.Depend().AsInject(utils.NewOf(results[0])))
		val.Set(reflect.MakeFunc(val.Type(), func([]reflect.Value) []reflect.Value {
			return results
		}))
	})
}
