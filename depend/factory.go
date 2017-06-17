// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package depend

import (
	"github.com/vlorc/gioc/types"
	"github.com/vlorc/gioc/utils"
	"reflect"
	"strings"
	"text/scanner"
)

func NewDependencyFactory() types.DependencyFactory {
	return &CoreDependencyFactory{
		pool: make(map[reflect.Type]types.Dependency),
	}
}

func (df *CoreDependencyFactory) Instance(impTyp interface{}) (types.Dependency, error) {
	return df.instance(utils.DirectlyType(utils.TypeOf(impTyp)))
}

func (df *CoreDependencyFactory) instance(typ reflect.Type) (dep types.Dependency, err error) {
	df.lock.RLock()
	dep = df.pool[typ]
	df.lock.RUnlock()

	if nil != dep {
		return
	}

	switch typ.Kind() {
	case reflect.Func:
		dep, err = df.resolveFunc(typ)
	case reflect.Struct:
		dep, err = df.resolveStruct(typ)
	default:
		err = types.NewError(types.ErrTypeNotSupport, typ)
	}

	if nil != dep {
		df.lock.Lock()
		df.pool[typ] = dep
		df.lock.Unlock()
	}
	return
}

func (df *CoreDependencyFactory) resolveArray(src reflect.Value) (dep types.Dependency, err error) {
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
	return
}

func (df *CoreDependencyFactory) resolveStruct(typ reflect.Type) (dep types.Dependency, err error) {
	des, err := df.resolveStructFields(typ, []*DependencyDescription{})
	if len(des) <= 0 {
		if nil == err {
			err = types.NewError(types.ErrDependencyNotNeed, typ)
		}
	} else {
		dep = NewStructDependency(typ, des)
	}

	return
}

func (df *CoreDependencyFactory) resolveStructFields(typ reflect.Type, src []*DependencyDescription) (dst []*DependencyDescription, err error) {
	for i, n := 0, typ.NumField(); i < n; i++ {
		field := typ.Field(i)
		if uint(field.Name[0])-uint(65) >= uint(26) {
			continue
		}
		tag := field.Tag.Get("inject")
		if "" == tag {
			continue
		}

		des := &DependencyDescription{}
		des.Index = i
		des.Type = field.Type

		df.resolveTag(tag, des)
		if 0 != (des.Flags & types.DEPENDENCY_FLAG_EXTENDS) {
			err = df.resolveExtends(des)
			if nil != err {
				return
			}
		}
		src = append(src, des)
	}

	dst = src

	return
}

func (df *CoreDependencyFactory) resolveExtends(des *DependencyDescription) (err error) {
	des.Depend, err = df.Instance(des.Type)
	return
}

func (df *CoreDependencyFactory) resolveTag(tag string, des *DependencyDescription) (err error) {
	var sc scanner.Scanner
	sc.Init(strings.NewReader(tag))
	flags := map[string]types.DependencyFlag{
		"optional": types.DEPENDENCY_FLAG_OPTIONAL,
		"extends":  types.DEPENDENCY_FLAG_EXTENDS,
		"default":  types.DEPENDENCY_FLAG_DEFAULT,
	}
	for token := sc.Scan(); token != scanner.EOF; token = sc.Scan() {
		str := sc.TokenText()
		switch token {
		case scanner.Char, scanner.String:
			des.Name = str[1 : len(str)-1]
		case scanner.Ident:
			if v, ok := flags[str]; ok {
				des.Flags |= v
			}
		}
	}
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
