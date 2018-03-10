// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package operation

import (
	"github.com/vlorc/gioc/builder"
	"github.com/vlorc/gioc/types"
)

func lazyProvider(con func() types.Container) func() types.Provider {
	return func() types.Provider {
		return con().AsProvider()
	}
}

func Dependency(val ...interface{}) DeclareHandle {
	return func(ctx *DeclareContext) {
		if len(val) <= 0 {
			val = []interface{}{ctx.Type}
		}
		if toDependency(ctx, val[0]) && nil != ctx.Factory {
			ctx.Factory = builder.NewBuilder(ctx.Factory, ctx.Depend).AsFactory()
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
