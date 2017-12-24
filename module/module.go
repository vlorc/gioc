// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package module

import (
	"github.com/vlorc/gioc/factory"
	"github.com/vlorc/gioc/types"
	"github.com/vlorc/gioc/utils"
	"reflect"
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
	moduleBootstrap(module, ctx.Bootstrap)
	return module, nil
}

func moduleBootstrap(module *CoreModule, fn []interface{}) {
	if len(fn) <= 0 {
		return
	}
	var dependFactory types.DependencyFactory
	var buildFactory types.BuilderFactory
	module.parent().AsProvider().Assign(&dependFactory)
	module.parent().AsProvider().Assign(&buildFactory)

	for _, v := range fn {
		moduleBootstrapApply(module, v, dependFactory, buildFactory)
	}
}

func moduleBootstrapApply(
	module *CoreModule,
	fn interface{},
	dependFactory types.DependencyFactory,
	buildFactory types.BuilderFactory) {
	val := reflect.ValueOf(fn)
	if reflect.Func != val.Kind() {
		return
	}
	if 0 == val.Type().NumIn() {
		val.Call(nil)
		return
	}
	dep, err := dependFactory.Instance(fn)
	if nil != err {
		return
	}

	build, err := buildFactory.Instance(factory.NewParamFactory(dep.Length()), dep)
	if nil != err {
		return
	}
	param, err := build.Build(module.container().AsProvider())
	if nil != err {
		return
	}
	val.Call(param.([]reflect.Value))
}
