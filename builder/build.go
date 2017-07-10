// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package builder

import (
	"github.com/vlorc/gioc/types"
	"github.com/vlorc/gioc/utils"
)

func BuildInstance(provider types.Provider, factory types.BeanFactory, depend types.Dependency) (instance interface{}, err error) {
	defer utils.Recover(&err)

	if instance, err = factory.Instance(provider); nil != err {
		return
	}
	if nil == instance {
		err = types.NewError(types.ErrInstanceNotFound, instance)
		return
	}
	if inject := depend.AsInject(instance); nil != inject {
		FullAllInstance(&BuildContext{Provider: provider, Inject: inject})
	} else {
		err = types.NewError(types.ErrInjectNotAllot, instance)
	}
	return
}
