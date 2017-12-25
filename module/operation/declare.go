// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package operation

import (
	"github.com/vlorc/gioc/module"
)

type DeclareHandle func(*DeclareContext)

func Declare(handle ...DeclareHandle) module.ModuleInitHandle {
	return declare(toRegistered,handle)
}

func Export(handle ...DeclareHandle) module.ModuleInitHandle {
	return declare(toExport,handle)
}

func declare(done func(*DeclareContext),handle []DeclareHandle) module.ModuleInitHandle {
	return func(ctx *module.ModuleInitContext) {
		dc := &DeclareContext{done: done,Context: ctx}
		for _, v := range handle {
			v(dc)
		}
		dc.Reset()
	}
}