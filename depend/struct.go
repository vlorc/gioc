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
	typ           reflect.Type
	dep           []*DependencyDescription
	injectFactory func(types.DependencyScan, reflect.Value) types.DependencyInject
}

type CoreDependencyScan struct {
	dep []*DependencyDescription
	pos int
}

type CoreDependencyInject struct {
	types.DependencyScan
	data reflect.Value
	set  func(types.DependencyScan, reflect.Value, reflect.Value)
	get  func(types.DependencyScan, reflect.Value) reflect.Value
}

type CoreDependencyFactory struct {
	mux  sync.RWMutex
	pool map[reflect.Type]types.Dependency
}
