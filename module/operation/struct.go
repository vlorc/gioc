// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package operation

import (
	"github.com/vlorc/gioc/module"
	"github.com/vlorc/gioc/types"
)

type DeclareContext struct {
	done       func(*DeclareContext)
	register   func(*module.ModuleInitContext) types.Register
	Name       string
	Type       interface{}
	Factory    types.BeanFactory
	Dependency types.Dependency
	Context    *module.ModuleInitContext
}

func (dc *DeclareContext) Reset() {
	if nil != dc.Factory && nil != dc.Type {
		dc.done(dc)
	}
	*dc = DeclareContext{done: dc.done, register: dc.register, Context: dc.Context}
}

type EventContext struct {
	*module.ModuleInitContext
	Listener types.EventListener
}
