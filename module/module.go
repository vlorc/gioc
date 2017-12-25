// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package module

import (
	"github.com/vlorc/gioc/types"
	"github.com/vlorc/gioc/utils"
	"github.com/vlorc/gioc/invoker"
)

func newModule(parent, container func() types.Container) *CoreModule {
	return &CoreModule{
		parent:    parent,
		container: container,
	}
}

func newWithModule(parent func() types.Container) *CoreModule {
	return &CoreModule{
		parent:    parent,
		container: lazyChildContainer(parent),
	}
}

func lazyChildContainer(parent func() types.Container) func() types.Container {
	return utils.Lazy(func() types.Container {
		return parent().NewChild()
	}).(func() types.Container)
}

func moduleInit(module *CoreModule, table ...ModuleInitHandle) (types.Module, error) {
	ctx := &ModuleInitContext{
		Parent:    module.parent,
		Container: module.container,
	}
	for _, v := range table {
		v(ctx)
	}
	moduleBootstrap(module,ctx.Bootstrap)
	return module, nil
}

func moduleBootstrap(module *CoreModule, fn []interface{}) {
	if len(fn) <= 0 {
		return
	}
	for _, v := range fn {
		invoker.NewInvoker(v,nil).ApplyWith(module.container().AsProvider())
	}
}

