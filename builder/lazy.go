// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package builder

import (
	"github.com/vlorc/gioc/types"
	"github.com/vlorc/gioc/utils"
	"reflect"
)

func buildLazyExtends(val reflect.Value, provider types.Provider, des types.DescriptorGetter) {
	proxy := utils.LazyProxy(func([]reflect.Value) []reflect.Value {
		results := []reflect.Value{reflect.New(des.Type().Out(0)).Elem()}
		buildDefault(provider, des.Depend().AsInject(utils.NewOf(results[0])))
		return results
	})
	val.Set(reflect.MakeFunc(val.Type(), proxy))
}

func buildLazyInstance(val reflect.Value, provider types.Provider, des types.DescriptorGetter) {
	proxy := utils.LazyProxy(func([]reflect.Value) []reflect.Value {
		typ := des.Type().Out(0)
		instance, err := provider.ResolveType(typ, des.Name(), -1)
		if nil != err {
			panic(err)
		}
		return []reflect.Value{utils.Convert(reflect.ValueOf(instance), typ)}
	})
	val.Set(reflect.MakeFunc(val.Type(), proxy))
}
