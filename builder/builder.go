// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package builder

import (
	"github.com/vlorc/gioc/types"
	"github.com/vlorc/gioc/utils"
)

var fullProcess []func(types.Provider,types.DependencyInject)

func init(){
	fullProcess = []func(types.Provider,types.DependencyInject){
		fullInstance,
		fullExtends,
	}
}

func (b *CoreBuilder) AsFactory() types.BeanFactory {
	return b
}

func (b *CoreBuilder) Build(provider types.Provider, factory types.BeanFactory) (interface{}, error) {
	if nil == factory {
		factory = b.factory
	}
	return buildAllInstance(provider, factory, b.depend)
}

func (this *CoreBuilder) Instance(provider types.Provider) (interface{}, error) {
	return buildAllInstance(provider, this.factory, this.depend)
}

func buildAllInstance(provider types.Provider,factory types.BeanFactory, depend types.Dependency) (instance interface{}, err error) {
	defer utils.Recover(&err)

	if instance, err = factory.Instance(provider);nil != err {
		return
	}

	if nil == instance {
		err = types.NewError(types.ErrInstanceNotFound, instance)
		return
	}

	if inject := depend.AsInject(instance); nil != inject {
		fullAllInstance(provider, inject)
	} else {
		err = types.NewError(types.ErrInjectNotAllot, instance)
	}

	return
}

func fullAllInstance(provider types.Provider, inject types.DependencyInject) {
	for inject.Next() {
		fullProcess[inject.Flags() & types.DEPENDENCY_FLAG_EXTENDS](provider,inject)
	}
}

func fullExtends(provider types.Provider, inject types.DependencyInject) {
	fullAllInstance(provider,inject.SubInject(provider))
}

func fullInstance(provider types.Provider, inject types.DependencyInject) {
	instance, err := provider.Resolve(inject.Type(), inject.Name())
	if nil == err {
		inject.SetInterface(instance)
		return
	}

	if 0 != inject.Flags()&types.DEPENDENCY_FLAG_DEFAULT {
		inject.SetValue(inject.Default())
		return
	}

	if 0 != inject.Flags()&types.DEPENDENCY_FLAG_OPTIONAL {
		return
	}

	panic(err)
}
