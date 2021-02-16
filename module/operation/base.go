// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package operation

import (
	"github.com/vlorc/gioc/factory"
	"github.com/vlorc/gioc/module"
	"github.com/vlorc/gioc/types"
	"github.com/vlorc/gioc/utils"
)

func lazyProvider(con func() types.Container) func() types.Provider {
	return func() types.Provider {
		return con().AsProvider()
	}
}

func Dependency(val ...interface{}) DeclareHandle {
	return func(ctx *DeclareContext) {
		if len(val) == 0 {
			val = []interface{}{ctx.Type}
		}
		if toDependency(ctx, val[0]) && nil != ctx.Factory {
			ctx.Factory = factory.NewDependencyFactory(ctx.Factory, ctx.Dependency)
		}
	}
}

func Type(typ interface{}) DeclareHandle {
	return func(ctx *DeclareContext) {
		ctx.Type = typ
	}
}

func Id(id string, args ...string) DeclareHandle {
	return func(ctx *DeclareContext) {
		for _, v := range args {
			c := *ctx
			c.Name = v
			c.done(&c)
		}
		ctx.Name = id
	}
}

func Name(name string) DeclareHandle {
	return func(ctx *DeclareContext) {
		ctx.Name = name
	}
}

func Factory(factory types.BeanFactory) DeclareHandle {
	return func(ctx *DeclareContext) {
		ctx.Factory = factory
	}
}

func Parent() string {
	return "#parent"
}

func Register(names ...string) func(*module.ModuleInitContext) types.Register {
	if len(names) == 0 {
		return func(ctx *module.ModuleInitContext) types.Register {
			return ctx.Container().AsRegister()
		}
	}
	if Parent() == names[0] {
		return func(ctx *module.ModuleInitContext) types.Register {
			return ctx.Parent().AsRegister()
		}
	}
	name := names[0]
	return func(ctx *module.ModuleInitContext) (r types.Register) {
		if err := ctx.Container().AsProvider().Load(&r, name); nil != err {
			utils.Panic(err)
		}
		return
	}
}

func Provider(names ...string) func(*module.ModuleInitContext) types.Provider {
	if len(names) == 0 {
		return func(ctx *module.ModuleInitContext) types.Provider {
			return ctx.Container().AsProvider()
		}
	}
	if Parent() == names[0] {
		return func(ctx *module.ModuleInitContext) types.Provider {
			return ctx.Parent().AsProvider()
		}
	}
	name := names[0]
	return func(ctx *module.ModuleInitContext) (p types.Provider) {
		if err := ctx.Container().AsProvider().Load(&p, name); nil != err {
			utils.Panic(err)
		}
		return
	}
}
