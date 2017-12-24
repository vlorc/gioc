// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package selector

import (
	"github.com/vlorc/gioc/types"
	"reflect"
	"sync"
)

type FactorySelector interface {
	SetFactory(reflect.Type, string, types.BeanFactory) error
	FactoryOf(reflect.Type, string) types.BeanFactory
}

type TypeSelector struct {
	lock    sync.RWMutex
	table   map[reflect.Type]types.Binder
	factory types.BinderFactory
}

type NamedSelector struct {
	lock     sync.RWMutex
	selector FactorySelector
}

type typeNameSelector map[typeName]types.BeanFactory
type nameSelector map[string]typeFactory

type typeName struct {
	typ  reflect.Type
	name string
}

type typeFactory struct {
	typ     reflect.Type
	factory types.BeanFactory
}

type CoreSelectorFactory struct{}
