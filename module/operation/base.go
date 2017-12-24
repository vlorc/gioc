// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package operation

import (
	"github.com/vlorc/gioc/types"
)

func lazyProvider(con func() types.Container) func() types.Provider {
	return func() types.Provider {
		return con().AsProvider()
	}
}

func Dependency(val interface{}) DeclareHandle {
	return func(ctx *DeclareContext) {
		var dependFactory types.DependencyFactory
		ctx.Context.Container().AsProvider().Assign(&dependFactory)
		dep, err := dependFactory.Instance(val)
		if nil == err {
			ctx.Depend = dep
			ctx.Value = val
		}
	}
}

func Value(val interface{}) DeclareHandle {
	return func(ctx *DeclareContext) {
		ctx.Value = val
	}
}

func Type(typ interface{}) DeclareHandle {
	return func(ctx *DeclareContext) {
		ctx.Type = typ
	}
}

func Name(name string) DeclareHandle {
	return func(ctx *DeclareContext) {
		ctx.Name = name
	}
}

func Factory() DeclareHandle {
	return toFactory
}

func Export() DeclareHandle {
	return toExport
}
