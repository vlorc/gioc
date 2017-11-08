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
	"github.com/vlorc/gioc/utils"
	"reflect"
)

// create a root container
func NewRootContainer() types.Container {
	registerFactory := register.NewRegisterFactory()
	binderFactory := binder.NewBinderFactory()
	dependFactory := depend.NewDependencyFactory()
	builderFactory := builder.NewBuilderFactory()
	selectorFactory := selector.NewSelectorFactory()
	invokerFactory := invoker.NewInvokerFactory()

	sel, err := selectorFactory.Instance(binderFactory)
	if nil != err {
		panic(err)
	}
	reg, err := registerFactory.Instance(sel)
	if nil != err {
		panic(err)
	}

	paramFactory, err := builderFactory.Instance(
		factory.ParamFactory(1),
		depend.NewFuncDependency(utils.TypeOf(selectorFactory.Instance), []*types.DependencyDescription{
			{Type: types.BinderFactoryType, Flags: types.DEPENDENCY_FLAG_DEFAULT, Default: reflect.Zero(types.BinderFactoryType)},
		}))
	if nil != err {
		panic(err)
	}
	reg.RegisterMethod(paramFactory.AsFactory(), selectorFactory.Instance, nil)


	reg.RegisterInterface(&registerFactory)
	reg.RegisterInterface(&binderFactory)
	reg.RegisterInterface(&dependFactory)
	reg.RegisterInterface(&builderFactory)
	reg.RegisterInterface(&selectorFactory)
	reg.RegisterInterface(&invokerFactory)

	return container.NewContainer(reg, nil, 0)
}
