// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package builder

import (
	"github.com/vlorc/gioc/types"
	"github.com/vlorc/gioc/utils"
)

func BuildInstance(provider types.Provider, factory types.BeanFactory, depend types.Dependency, option ...func(*types.BuildContext)) (instance interface{}, err error) {
	defer utils.Recover(&err)
	if instance, err = factory.Instance(provider); nil != err {
		return
	}
	if nil == instance {
		err = types.NewError(types.ErrInstanceNotFound, instance)
		return
	}
	if inject := depend.AsInject(instance); nil != inject {
		buildDefault(provider, inject, option...)
	} else {
		err = types.NewError(types.ErrInjectNotAllot, instance)
	}
	return
}

func buildDefault(provider types.Provider, inject types.DependencyInject, option ...func(*types.BuildContext)) {
	ctx := &types.BuildContext{
		Provider: provider,
		Inject:   inject,
		FullBefore: func(*types.BuildContext) bool {
			return true
		},
		FullAfter: func(*types.BuildContext) {
		},
	}
	for _, f := range option {
		f(ctx)
	}
	FullAllInstance(ctx)
}

func buildBefore(provider types.Provider, inject types.DependencyInject, before func(*types.BuildContext) bool) {
	buildDefault(provider, inject, func(ctx *types.BuildContext) {
		ctx.FullBefore = before
	})
}
