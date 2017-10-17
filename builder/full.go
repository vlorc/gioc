// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package builder

import (
	"github.com/vlorc/gioc/types"
)

var fullProcess []func(*types.BuildContext)

func init() {
	fullProcess = []func(*types.BuildContext){
		FullInstance,
		FullExtends,
		FullLazyInstance,
		FullLazyExtends,
	}
}

func FullAllInstance(ctx *types.BuildContext) {
	for ctx.Inject.Next() {
		ctx.Descriptor = ctx.Inject
		if ctx.FullBefore(ctx) {
			fullProcess[ctx.Descriptor.Flags()&3](ctx)
			ctx.FullAfter(ctx)
		}
	}
}

func FullExtends(ctx *types.BuildContext) {
	buildDefault(ctx.Provider, ctx.Inject.SubInject(ctx.Provider))
}

func FullInstance(ctx *types.BuildContext) {
	instance, err := ctx.Provider.Resolve(
		ctx.Descriptor.Type(),
		ctx.Descriptor.Name(),
	)
	if nil == err {
		ctx.Inject.SetInterface(instance)
		return
	}
	if 0 != ctx.Descriptor.Flags()&types.DEPENDENCY_FLAG_DEFAULT {
		ctx.Inject.SetValue(ctx.Descriptor.Default())
		return
	}
	if 0 != ctx.Descriptor.Flags()&types.DEPENDENCY_FLAG_OPTIONAL {
		return
	}
	panic(err)
}

func FullLazyInstance(ctx *types.BuildContext) {
	MakeLazyInstance(ctx.Inject.AsValue(), ctx.Provider, ctx.Inject.AsDescriptorGetter())
}

func FullLazyExtends(ctx *types.BuildContext) {
	MakeLazyExtends(ctx.Inject.AsValue(), ctx.Provider, ctx.Inject.AsDescriptorGetter())
}
