// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package depend

import (
	"github.com/vlorc/gioc/types"
	"github.com/vlorc/gioc/utils"
	"reflect"
)

func (df *CoreDependencyFactory) resolveArray(typ reflect.Type,val reflect.Value) (dep types.Dependency, _ error) {
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

func (df *CoreDependencyFactory) resolveFunc(typ reflect.Type,val reflect.Value) (dep types.Dependency, _ error) {
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
	des := &types.DependencyDescription{
		Index: index,
		Type:  typ,
	}

	if temp := utils.DirectlyType(typ);
		reflect.Struct == temp.Kind() && "" == temp.Name(){
		des.Flags |= types.DEPENDENCY_FLAG_EXTENDS
		des.Depend = df.anonymousToDependency(temp)
	}

	return append(dep,des)
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

func (df *CoreDependencyFactory) resolveMap(typ reflect.Type,val reflect.Value) (dep types.Dependency, _ error) {
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

func (df *CoreDependencyFactory) resolveStruct(typ reflect.Type,_ reflect.Value) (dep types.Dependency, _ error) {
	if "" == typ.Name() {
		dep = df.anonymousToDependency(typ)
		return
	}

	df.lock.RLock()
	dep = df.pool[typ]
	df.lock.RUnlock()

	if nil != dep {
		return
	}

	if dep = df.namedToDependency(typ); nil != dep{
		df.lock.Lock()
		df.pool[typ] = dep
		df.lock.Unlock()
	}
	return
}

func (df *CoreDependencyFactory) structToDependency(typ reflect.Type,skip func(reflect.StructField,string)bool) (dep types.Dependency) {
	arr := []*types.DependencyDescription{}
	for i, n := 0, typ.NumField(); i < n; i++ {
		arr = df.appendField(arr,typ.Field(i),i,skip)
	}
	if len(arr) > 0 {
		dep = NewStructDependency(typ, arr)
	}
	return
}

func (df *CoreDependencyFactory) namedToDependency(typ reflect.Type) types.Dependency{
	return df.structToDependency(typ,func(field reflect.StructField,tag string)bool{
		return uint(field.Name[0] - 65) < uint(26) && "" != tag && "-" != tag
	})
}

func (df *CoreDependencyFactory) anonymousToDependency(typ reflect.Type) (dep types.Dependency) {
	return df.structToDependency(typ,func(field reflect.StructField,tag string)bool{
		return uint(field.Name[0] - 65) < uint(26) && "-" != tag
	})
}

func (df *CoreDependencyFactory) appendField(
	dep []*types.DependencyDescription,
	field reflect.StructField,
	index int,
	skip func(reflect.StructField,string)bool) []*types.DependencyDescription {

	tag := field.Tag.Get("inject")
	if !skip(field,tag) {
		return dep
	}

	des := &types.DependencyDescription{
		Name:field.Name,
		Index:index,
		Type:field.Type,
	}

	if "" != tag {
		df.tagParser.Resolve(df,tag, NewDescriptor(des))
	}


	return append(dep, des)
}