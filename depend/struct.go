// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package depend

import (
	"github.com/vlorc/gioc/types"
	"reflect"
	"sync"
)

type ResolveHandle func(reflect.Type, reflect.Value) types.Dependency

type CoreDependency struct {
	typ     reflect.Type
	dep     []*types.DependencyDescription
	factory func(reflect.Value) types.Reflect
}

type CoreDependencyScan struct {
	dep []*types.DependencyDescription
	pos int
}

type CoreDependencyInject struct {
	types.DependencyScan
	types.Reflect
}

type CoreDependencyFactory struct {
	lock      sync.RWMutex
	resolve   map[reflect.Kind]ResolveHandle
	pool      map[reflect.Type]types.Dependency
	tagParser *TagParser
}

type DescriptorGetter struct {
	des *types.DependencyDescription
}

type DescriptorSetter struct {
	des *types.DependencyDescription
}

type Descriptor struct {
	types.DescriptorGetter
	types.DescriptorSetter
}

type ParamReflect []reflect.Value
type StructReflect reflect.Value
type ArrayReflect reflect.Value
type MapReflect reflect.Value
