// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package depend

import (
	"github.com/vlorc/gioc/text"
	"github.com/vlorc/gioc/types"
	"github.com/vlorc/gioc/utils"
	"reflect"
)

func NewDependencyFactory() types.DependencyFactory {
	obj := &CoreDependencyFactory{
		pool:   make(map[reflect.Type]types.Dependency),
		parser: text.NewTagParser(),
	}
	obj.resolve = map[reflect.Kind]ResolveHandle{
		reflect.Array:  obj.resolveArray,
		reflect.Map:    obj.resolveMap,
		reflect.Func:   obj.resolveFunc,
		reflect.Struct: obj.resolveStruct,
	}
	return obj
}

func (df *CoreDependencyFactory) Instance(impTyp interface{}) (types.Dependency, error) {
	return df.instance(utils.DirectlyType(utils.TypeOf(impTyp)), impTyp)
}

func (df *CoreDependencyFactory) instance(typ reflect.Type, val interface{}) (dep types.Dependency, err error) {
	defer utils.Recover(&err)

	resolve, ok := df.resolve[typ.Kind()]
	if !ok {
		err = types.NewWithError(types.ErrTypeNotSupport, typ)
		return
	}
	if dep = resolve(typ, reflect.ValueOf(val)); nil == dep {
		err = types.NewWithError(types.ErrDependencyNotNeed, typ)
	}
	return
}

func NewDependency(typ reflect.Type, dep []*types.DependencyDescription, reflectFactory func(reflect.Value) types.Reflect) types.Dependency {
	return &CoreDependency{
		typ:     typ,
		dep:     dep,
		factory: reflectFactory,
	}
}

func NewStructDependency(typ reflect.Type, dep []*types.DependencyDescription) types.Dependency {
	return NewDependency(typ, dep, NewStructReflect)
}

func NewArrayDependency(typ reflect.Type, dep []*types.DependencyDescription) types.Dependency {
	return NewDependency(typ, dep, NewArrayReflect)
}

func NewMapDependency(typ reflect.Type, dep []*types.DependencyDescription) types.Dependency {
	return NewDependency(typ, dep, NewMapReflect)
}

func NewFuncDependency(typ reflect.Type, dep []*types.DependencyDescription) types.Dependency {
	return NewDependency(typ, dep, NewParamReflect)
}
