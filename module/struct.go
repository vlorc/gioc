// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package module

import "github.com/vlorc/gioc/types"

type CoreModuleFactory struct {
	table []ModuleInitHandle
}

type CoreModuleForFactory struct {
	table []ModuleInitHandle
}

type CoreModule struct {
	parent    func() types.Container
	container func() types.Container
}

type ModuleInitContext struct {
	Bootstrap []interface{}
	Parent    func() types.Container
	Container func() types.Container
}

type ModuleInitHandle func(*ModuleInitContext)
