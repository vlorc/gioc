// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package operation

import (
	"github.com/vlorc/gioc/factory"
	"github.com/vlorc/gioc/module"
	"github.com/vlorc/gioc/types"
	"github.com/vlorc/gioc/utils"
)

type DeclareHandle func(*DeclareContext)

func Declare(handle ...DeclareHandle) module.ModuleInitHandle {
	return declare(currentRegister, toDeclare, handle)
}

func Export(handle ...DeclareHandle) module.ModuleInitHandle {
	return declare(parentRegister, toExport, handle)
}

func Primary(handle ...DeclareHandle) module.ModuleInitHandle {
	return PrimaryWith(parentRegister, handle...)
}

func Miss(handle ...DeclareHandle) module.ModuleInitHandle {
	return MissWith(parentRegister, handle...)
}

func PrimaryWith(register func(*module.ModuleInitContext) types.Register, handle ...DeclareHandle) module.ModuleInitHandle {
	return declare(register, toPrimary, handle)
}

func MissWith(register func(*module.ModuleInitContext) types.Register, handle ...DeclareHandle) module.ModuleInitHandle {
	return declare(register, toMiss, handle)
}

func declare(register func(*module.ModuleInitContext) types.Register, done func(*DeclareContext), handle []DeclareHandle) module.ModuleInitHandle {
	return func(ctx *module.ModuleInitContext) {
		dc := &DeclareContext{done: done, register: register, Context: ctx}
		for _, v := range handle {
			v(dc)
		}
		dc.Reset()
	}
}

func currentRegister(ctx *module.ModuleInitContext) types.Register {
	return ctx.Container().AsRegister()
}

func parentRegister(ctx *module.ModuleInitContext) types.Register {
	return ctx.Parent().AsRegister()
}

func toDeclare(ctx *DeclareContext) {
	ctx.register(ctx.Context).Factory(ctx.Factory, ctx.Type, ctx.Name)
}

func toExport(ctx *DeclareContext) {
	var bean types.BeanFactory
	if nil != ctx.Dependency {
		bean = factory.NewExportFactory(ctx.Factory, lazyProvider(ctx.Context.Container))
	} else {
		bean = ctx.Factory
	}

	ctx.register(ctx.Context).Factory(bean, ctx.Type, ctx.Name)
}

func toPrimary(ctx *DeclareContext) {
	var bean types.BeanFactory
	if nil != ctx.Dependency {
		bean = factory.NewExportFactory(ctx.Factory, lazyProvider(ctx.Context.Container))
	} else {
		bean = ctx.Factory
	}
	ctx.register(ctx.Context).Selector().Set(utils.TypeOf(ctx.Type), ctx.Name, bean)
}

func toMiss(ctx *DeclareContext) {
	var bean types.BeanFactory
	if nil != ctx.Dependency {
		bean = factory.NewExportFactory(ctx.Factory, lazyProvider(ctx.Context.Container))
	} else {
		bean = ctx.Factory
	}
	ctx.register(ctx.Context).Selector().Put(utils.TypeOf(ctx.Type), ctx.Name, bean)
}
