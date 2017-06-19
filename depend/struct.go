// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package depend

import (
	"github.com/vlorc/gioc/types"
	"reflect"
	"sync"
)

type DependencyDescription struct {
	Type    reflect.Type
	Name    string
	Index   int
	Default interface{}
	Depend  types.Dependency
	Flags   types.DependencyFlag
}

type CoreDependency struct {
	typ            reflect.Type
	dep            []*DependencyDescription
	factory        func(reflect.Value) types.Reflect
}

type CoreDependencyScan struct {
	dep []*DependencyDescription
	pos int
}

type CoreDependencyInject struct {
	types.DependencyScan
	types.Reflect
}

type CoreDependencyFactory struct {
	lock  sync.RWMutex
	pool map[reflect.Type]types.Dependency
}

type CoreParamReflect []reflect.Value
type CoreStructReflect reflect.Value
type CoreArrayReflect reflect.Value
type CoreMapReflect reflect.Value
