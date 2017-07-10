// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package binder

import (
	"github.com/vlorc/gioc/types"
	"reflect"
	"sync"
)

type NamedBind struct {
	lock  sync.RWMutex
	table map[string]types.BeanFactory
}

type ProxyBind struct {
	read  types.Mapper
	write types.Binder
}

type SelectorBind struct {
	typ      reflect.Type
	selector types.Selector
}

type CoreBinderFactory struct{}
