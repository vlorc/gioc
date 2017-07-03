// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package selector

import (
	"sync"
	"reflect"
	"github.com/vlorc/gioc/types"
)

type TypeSelector struct {
	lock    sync.RWMutex
	table   map[reflect.Type]types.Binder
	factory types.BinderFactory
}

type TypeName struct {
	Type reflect.Type
	Name string
}

type TypeNameSelector struct {
	lock    sync.RWMutex
	table   map[TypeName]types.BeanFactory
}

type CoreSelectorFactory struct {}