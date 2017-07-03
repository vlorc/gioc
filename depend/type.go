// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package depend

import (
	"github.com/vlorc/gioc/types"
	"reflect"
)

func (df *CoreDependencyFactory) resolveArray(typ reflect.Type,val reflect.Value) (dep types.Dependency, err error) {
	if val.Len() <= 0{
		return
	}

	arr := []*types.DependencyDescription{}
	for i, n := 0, val.Len(); i < n; i++ {
		arr = df.appendValue(arr,val.Index(i),i)
	}

	dep = NewArrayDependency(typ, arr)
	return
}

func (df *CoreDependencyFactory) appendValue(dep []*types.DependencyDescription,val reflect.Value,index int)  []*types.DependencyDescription {
	typ := val.Type()
	if reflect.Ptr == typ.Kind() {
		typ = typ.Elem()
	}

	return append(dep, &types.DependencyDescription{
		Index: index,
		Type:  typ,
	})
}

func (df *CoreDependencyFactory) resolveFunc(typ reflect.Type,val reflect.Value) (dep types.Dependency, err error) {
	if typ.NumIn() <= 0 {
		return
	}

	arr := []*types.DependencyDescription{}
	for i, n := 0, typ.NumIn(); i < n; i++ {
		arr = df.appendParam(arr,typ.In(i),i)
	}

	dep = NewFuncDependency(typ, arr)
	return
}

func (df *CoreDependencyFactory) appendParam(dep []*types.DependencyDescription,typ reflect.Type,index int)  []*types.DependencyDescription {
	return append(dep, &types.DependencyDescription{
		Index: index,
		Type:  typ,
	})
}

func (df *CoreDependencyFactory) appendKey(dep []*types.DependencyDescription,m reflect.Value,k reflect.Value)  []*types.DependencyDescription{
	if reflect.String != k.Kind(){
		return dep
	}

	typ := m.MapIndex(k).Type()
	if reflect.Ptr == typ.Kind() {
		typ = typ.Elem()
	}

	return append(dep, &types.DependencyDescription{
		Name:k.String(),
		Type: typ,
	})
}

func (df *CoreDependencyFactory) resolveMap(typ reflect.Type,val reflect.Value) (dep types.Dependency, err error) {
	if val.Len() <= 0 {
		return
	}

	arr := []*types.DependencyDescription{}
	for _,k := range val.MapKeys() {
		arr = df.appendKey(arr,val,k)
	}

	if len(arr) > 0 {
		dep = NewMapDependency(typ,arr)
	}
	return
}

func (df *CoreDependencyFactory) resolveStruct(typ reflect.Type,_ reflect.Value) (dep types.Dependency, err error) {
	df.lock.RLock()
	dep = df.pool[typ]
	df.lock.RUnlock()

	if nil != dep {
		return
	}

	arr := []*types.DependencyDescription{}
	for i, n := 0, typ.NumField(); i < n; i++ {
		arr = df.appendField(arr,typ.Field(i),i)
	}

	if len(arr) > 0 {
		dep = NewStructDependency(typ, arr)

		df.lock.Lock()
		df.pool[typ] = dep
		df.lock.Unlock()
	}
	return
}

func (df *CoreDependencyFactory) appendField(dep []*types.DependencyDescription,field reflect.StructField, index int)  []*types.DependencyDescription {
	if uint(field.Name[0])-uint(65) >= uint(26) {
		return dep
	}

	tag := field.Tag.Get("inject")
	if "" == tag {
		return dep
	}

	if "-" != tag {
		temp := &types.DependencyDescription{
			Name:field.Name,
			Index:index,
			Type:field.Type,
		}
		df.tagParser.Resolve(df,tag, NewDescriptor(temp))
		dep = append(dep, temp)
	}

	return dep
}