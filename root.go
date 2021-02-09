// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package gioc

import (
	"github.com/vlorc/gioc/container"
	"github.com/vlorc/gioc/dependency"
	"github.com/vlorc/gioc/event"
	"github.com/vlorc/gioc/factory"
	"github.com/vlorc/gioc/invoker"
	"github.com/vlorc/gioc/module"
	. "github.com/vlorc/gioc/module/operation"
	"github.com/vlorc/gioc/provider"
	"github.com/vlorc/gioc/register"
	"github.com/vlorc/gioc/selector"
	"github.com/vlorc/gioc/types"
	"github.com/vlorc/gioc/utils"
)

// create a root container
func NewRootContainer() types.Container {
	general := selector.NewGeneralSelector()
	root := container.NewContainer(
		register.NewRegister(general),
		provider.NewWithProvider(nil, general),
	)

	for _, v := range []interface{}{
		dependency.NewDependencyFactory,
		invoker.NewInvokerFactory,
		provider.NewProviderFactory,
	} {
		f, typ, _ := factory.NewMethodFactory(v, nil)
		root.AsRegister().RegisterFactory(
			factory.NewSingleFactory(f),
			typ,
		)
	}
	return root
}

// create a root module
func NewRootModule(table ...module.ModuleInitHandle) types.Module {
	return module.NewModuleFor(
		utils.Lazy(NewRootContainer).(func() types.Container),
		Import(event.EventModuleFor("root", "parent")),
		Join(table...),
		Event(Emit("ready"), Emit("init")),
	)
}
