// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package depend

import (
	"github.com/vlorc/gioc/types"
	"github.com/vlorc/gioc/utils"
	"reflect"
)

func NewDependencyFactory() types.DependencyFactory {
	return  &CoreDependencyFactory{
		pool: make(map[reflect.Type]types.Dependency),
		tagParser:NewTagParser(),
	}
}

func (df *CoreDependencyFactory) Instance(impTyp interface{}) (types.Dependency, error) {
	return df.instance(utils.DirectlyType(utils.TypeOf(impTyp)),impTyp)
}

func (df *CoreDependencyFactory) instance(typ reflect.Type,val interface{}) (dep types.Dependency, err error) {
	defer utils.Recover(&err)

	switch typ.Kind() {
	case reflect.Func:
		dep, err = df.resolveFunc(typ)
	case reflect.Struct:
		dep, err = df.resolveStruct(typ)
	case reflect.Array:
		dep, err = df.resolveArray(reflect.ValueOf(val))
	case reflect.Map:
		dep, err = df.resolveMap(reflect.ValueOf(val))
	default:
		err = types.NewError(types.ErrTypeNotSupport, typ)
	}

	if nil == dep && nil == err {
		err = types.NewError(types.ErrDependencyNotNeed, typ)
	}

	return
}

func (df *CoreDependencyFactory) resolveArray(src reflect.Value) (dep types.Dependency, err error) {
	if src.Len() <= 0{
		return
	}

	des := []*DependencyDescription{}
	for i, n := 0, src.Len(); i < n; i++ {

		typ := src.Index(i).Type()
		if reflect.Ptr == typ.Kind() {
			typ = typ.Elem()
		}

		des = append(des, &DependencyDescription{
			Index: i,
			Type:  typ,
		})
	}

	dep = NewArrayDependency(src.Type(), des)
	return
}

func (df *CoreDependencyFactory) resolveFunc(typ reflect.Type) (dep types.Dependency, err error) {
	if typ.NumIn() <= 0 {
		return
	}

	des := []*DependencyDescription{}
	for i, n := 0, typ.NumIn(); i < n; i++ {
		des = append(des, &DependencyDescription{
			Index: i,
			Type:  typ.In(i),
		})
	}

	dep = NewFuncDependency(typ, des)
	return
}

func (df *CoreDependencyFactory) resolveMap(src reflect.Value) (dep types.Dependency, err error) {
	if src.Len() <= 0 {
		return
	}

	des := []*DependencyDescription{}
	for i,k := range src.MapKeys() {
		if reflect.String != k.Kind(){
			continue
		}

		typ := src.MapIndex(k).Type()
		if reflect.Ptr == typ.Kind() {
			typ = typ.Elem()
		}

		des = append(des, &DependencyDescription{
			Name:k.String(),
			Index: i,
			Type: typ,
		})
	}

	if len(des) > 0 {
		dep = NewMapDependency(src.Type(),des)
	}
	return
}

func (df *CoreDependencyFactory) resolveStruct(typ reflect.Type) (dep types.Dependency, err error) {
	df.lock.RLock()
	dep = df.pool[typ]
	df.lock.RUnlock()

	if nil != dep {
		return
	}

	des := df.resolveFields(typ, []*DependencyDescription{})
	if len(des) > 0 {
		dep = NewStructDependency(typ, des)

		df.lock.Lock()
		df.pool[typ] = dep
		df.lock.Unlock()
	}
	return
}

func (df *CoreDependencyFactory) resolveFields(typ reflect.Type, src []*DependencyDescription) (dst []*DependencyDescription) {
	for i, n := 0, typ.NumField(); i < n; i++ {
		field := typ.Field(i)

		if uint(field.Name[0])-uint(65) >= uint(26) {
			continue
		}

		tag := field.Tag.Get("inject")
		if "" == tag {
			continue
		}

		des := &DependencyDescription{
			Name:field.Name,
			Index:i,
			Type:field.Type,
		}

		if "-" != tag {
			df.tagParser.Resolve(df,tag, NewDescriptor(des))
			src = append(src, des)
		}
	}

	dst = src
	return
}

func NewDependency(typ reflect.Type, dep []*DependencyDescription, reflectFactory func(reflect.Value) types.Reflect) types.Dependency {
	return &CoreDependency{
		typ:            typ,
		dep:            dep,
		factory: reflectFactory,
	}
}

func NewStructDependency(typ reflect.Type, dep []*DependencyDescription) types.Dependency {
	return NewDependency(typ, dep, NewStructReflect)
}

func NewArrayDependency(typ reflect.Type, dep []*DependencyDescription) types.Dependency {
	return NewDependency(typ, dep, NewArrayReflect)
}

func NewMapDependency(typ reflect.Type, dep []*DependencyDescription) types.Dependency {
	return NewDependency(typ, dep, NewMapReflect)
}

func NewFuncDependency(typ reflect.Type, dep []*DependencyDescription) types.Dependency {
	return NewDependency(typ, dep, NewParamReflect)
}
