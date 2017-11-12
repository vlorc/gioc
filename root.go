// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package gioc

import (
	"github.com/vlorc/gioc/binder"
	"github.com/vlorc/gioc/builder"
	"github.com/vlorc/gioc/container"
	"github.com/vlorc/gioc/depend"
	"github.com/vlorc/gioc/factory"
	"github.com/vlorc/gioc/invoker"
	"github.com/vlorc/gioc/register"
	"github.com/vlorc/gioc/selector"
	"github.com/vlorc/gioc/types"
	"github.com/vlorc/gioc/provider"
)

// create a root container
func NewRootContainer() types.Container {
	sel := selector.NewTypeSelector(binder.NewBinderFactory())
	root := container.NewContainer(
		register.NewRegister(sel),
		provider.NewWithProvider(sel,nil),
	)

	init_root(root,
		depend.NewDependencyFactory,
		builder.NewBuilderFactory,
		selector.NewSelectorFactory,
		invoker.NewInvokerFactory,
		register.NewRegisterFactory,
		provider.NewProviderFactory,
	)
	return root
}

// init root container
func init_root(root types.Container,args ...interface{})  {
	for _,v := range args {
		f,typ,_ := factory.NewMethodFactory(v,nil)
		root.AsRegister().RegisterFactory(
			factory.NewSingleFactory(f),
			typ,
		)
	}
}
