// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package binder

import (
	"github.com/vlorc/gioc/types"
	"sync"
)

type NamedBind struct {
	m      sync.RWMutex
	table  map[string]types.BeanFactory
	mapper types.Mapper
}

type CoreBinderFactory struct {
}
