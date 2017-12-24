// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package operation

import (
	"github.com/vlorc/gioc/module"
	"github.com/vlorc/gioc/types"
)

type DeclareContext struct {
	Name    string
	Value   interface{}
	Type    interface{}
	Factory types.BeanFactory
	Depend  types.Dependency
	Context *module.ModuleInitContext
}

type DeclareHandle func(*DeclareContext)

func Declare(handle ...DeclareHandle) module.ModuleInitHandle {
	return func(ctx *module.ModuleInitContext) {
		c := &DeclareContext{
			Context: ctx,
		}
		for _, v := range handle {
			v(c)
		}
		if nil != c.Factory && nil != c.Type {
			ctx.Container().AsRegister().RegisterFactory(c.Factory, c.Type, c.Name)
		}
	}
}
