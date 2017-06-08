// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package builder

import (
	"github.com/vlorc/gioc/types"
)

func (b *CoreBuilder) AsFactory() types.BeanFactory {
	return b
}

func (b *CoreBuilder) Build(provider types.Provider, factory types.BeanFactory) (interface{}, error) {
	if nil == factory {
		factory = b.factory
	}
	return buildInstance(factory, b.depend, provider)
}

func (this *CoreBuilder) Instance(provider types.Provider) (interface{}, error) {
	return buildInstance(this.factory, this.depend, provider)
}

func buildInstance(factory types.BeanFactory, depend types.Dependency, provider types.Provider) (instance interface{}, err error) {
	instance, err = factory.Instance(provider)
	if nil != err {
		return
	}
	if nil == instance {
		err = types.NewError(types.ErrInstanceNotFound, instance)
		return
	}

	inject := depend.AsInject(instance)
	if nil == inject {
		err = types.NewError(types.ErrInjectNotAllot, instance)
		return
	}

	err = fullInstance(inject, provider)

	return
}

func fullInstance(inject types.DependencyInject, provider types.Provider) (err error) {
	var v interface{}
	for inject.Next() {
		if 0 != inject.Flags()&types.DEPENDENCY_FLAG_EXTENDS {
			if err = fullInstance(inject.SubInject(provider), provider); nil != err {
				break
			}
			continue
		}

		if v, err = provider.Resolve(inject.Type(), inject.Name()); nil != err {
			if 0 != inject.Flags()&types.DEPENDENCY_FLAG_DEFAULT {
				v = inject.Default()
				err = nil
			} else if 0 != inject.Flags()&types.DEPENDENCY_FLAG_OPTIONAL {
				continue
			} else {
				err = types.NewError(types.ErrInstanceNotFound, inject.Type(), inject.Name())
				break
			}
		}
		err = inject.SetInterface(v)
	}
	return
}
