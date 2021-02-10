// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package dependency

import (
	"github.com/vlorc/gioc/types"
	"reflect"
	"sync"
)

type ResolveHandle func(reflect.Type, reflect.Value) types.Dependency

type CoreDependency struct {
	typ     reflect.Type
	dep     []types.DependencyDescriptor
	factory func(reflect.Value) types.Reflect
}

type CoreDependencyScan struct {
	dep []types.DependencyDescriptor
	pos int
}

type CoreDependencyInject struct {
	types.DependencyScan
	types.Reflect
}

type CoreDependencyFactory struct {
	lock    sync.RWMutex
	resolve map[reflect.Kind]ResolveHandle
	pool    map[reflect.Type]types.Dependency
	parser  types.TextParser
	tag     string
}

type ParamReflect []reflect.Value
type StructReflect reflect.Value
type ArrayReflect reflect.Value
type MapReflect reflect.Value
