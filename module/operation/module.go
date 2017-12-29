// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package operation

import (
	"github.com/vlorc/gioc/module"
	"github.com/vlorc/gioc/types"
	"github.com/vlorc/gioc/utils"
)

func Import(factory ...types.ModuleFactory) module.ModuleInitHandle {
	return func(ctx *module.ModuleInitContext) {
		for _, v := range factory {
			if _, err := v.Instance(ctx.Container); nil != err {
				panic(err)
			}
		}
	}
}

func Join(handle ...module.ModuleInitHandle) module.ModuleInitHandle {
	return func(ctx *module.ModuleInitContext) {
		for _,v := range handle {
			v(ctx)
		}
	}
}

func Bootstrap(fn ...interface{}) module.ModuleInitHandle {
	return func(ctx *module.ModuleInitContext) {
		ctx.Bootstrap = append(ctx.Bootstrap, fn...)
	}
}

func Score(container func()types.Container) module.ModuleInitHandle {
	return func(ctx *module.ModuleInitContext) {
		ctx.Container = utils.Lazy(container).(func()types.Container)
	}
}