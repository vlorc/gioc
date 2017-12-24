// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package operation

import (
	"github.com/vlorc/gioc/module"
	"github.com/vlorc/gioc/types"
)

func Import(factory ...types.ModuleFactory) module.ModuleInitHandle {
	return func(ctx *module.ModuleInitContext) {
		for _, v := range factory {
			_, err := v.Instance(ctx.Container)
			if nil != err {
				panic(err)
			}
		}
	}
}

func Bootstrap(fn ...interface{}) module.ModuleInitHandle {
	return func(ctx *module.ModuleInitContext) {
		ctx.Bootstrap = append(ctx.Bootstrap, fn...)
	}
}
