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
	df.mux.RLock()
	dep = df.pool[typ]
	df.mux.RUnlock()

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
		df.mux.Lock()
		df.pool[typ] = dep
		df.mux.Unlock()
	}
	return
}

func (df *CoreDependencyFactory) resolveArray(src reflect.Value) (dep types.Dependency, err error) {
	des := []*DependencyDescription{}
	for i, n := 0, src.Len(); i < n; i++ {
		des = append(des, &DependencyDescription{
			Index: i,
			Type:  src.Index(i).Type(),
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

func NewDependency(typ reflect.Type, dep []*DependencyDescription, injectFactory func(types.DependencyScan, reflect.Value) types.DependencyInject) types.Dependency {
	return &CoreDependency{
		typ:           typ,
		dep:           dep,
		injectFactory: injectFactory,
	}
}

func NewBaseDependency(
	typ reflect.Type,
	dep []*DependencyDescription,
	setter func(types.DependencyScan, reflect.Value, reflect.Value),
	getter func(types.DependencyScan, reflect.Value) reflect.Value) types.Dependency {
	return NewDependency(
		typ,
		dep,
		func(scan types.DependencyScan, val reflect.Value) types.DependencyInject {
			return &CoreDependencyInject{
				scan,
				val,
				setter,
				getter,
			}
		},
	)
}

func NewStructDependency(typ reflect.Type, dep []*DependencyDescription) types.Dependency {
	return NewBaseDependency(typ, dep, structSetter, structGetter)
}

func NewArrayDependency(typ reflect.Type, dep []*DependencyDescription) types.Dependency {
	return NewBaseDependency(typ, dep, arraySetter, arrayGetter)
}

func NewMapDependency(typ reflect.Type, dep []*DependencyDescription) types.Dependency {
	return NewBaseDependency(typ, dep, mapSetter, mapGetter)
}

func NewFuncDependency(typ reflect.Type, dep []*DependencyDescription) types.Dependency {
	return NewBaseDependency(typ, dep, funcSetter, funcGetter)
}

func funcSetter(scan types.DependencyScan, src, val reflect.Value) {
	src.Index(scan.Index()).Set(reflect.ValueOf(val))
}

func structSetter(scan types.DependencyScan, src, val reflect.Value) {
	src.Field(scan.Index()).Set(val)
}

func mapSetter(scan types.DependencyScan, src, val reflect.Value) {
	src.SetMapIndex(reflect.ValueOf(scan.Name()), val)
}

func arraySetter(scan types.DependencyScan, src, val reflect.Value) {
	src.Index(scan.Index()).Set(val)
}

func funcGetter(scan types.DependencyScan, src reflect.Value) reflect.Value {
	return src.Index(scan.Index()).Interface().(reflect.Value)
}

func structGetter(scan types.DependencyScan, src reflect.Value) reflect.Value {
	return src.Field(scan.Index())
}

func mapGetter(scan types.DependencyScan, src reflect.Value) reflect.Value {
	return src.MapIndex(reflect.ValueOf(scan.Name()))
}

func arrayGetter(scan types.DependencyScan, src reflect.Value) reflect.Value {
	return src.Index(scan.Index())
}
