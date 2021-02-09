// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package container

import (
	"github.com/vlorc/gioc/types"
)

func NewWithContainer(provider types.Provider) types.Container {
	var selectorFactory types.SelectorFactory
	var registerFactory types.RegisterFactory
	var providerFactory types.ProviderFactory
	provider.Assign(&selectorFactory)
	provider.Assign(&registerFactory)
	provider.Assign(&providerFactory)

	if nil == selectorFactory || nil == registerFactory || nil == providerFactory {
		return nil
	}
	sel, err := selectorFactory.Instance()
	if nil != err {
		panic(err)
	}
	reg, err := registerFactory.Instance(sel)
	if nil != err {
		panic(err)
	}
	pro, err := providerFactory.Instance(provider, sel)
	if nil != err {
		panic(err)
	}
	return NewContainer(reg, pro)
}

func NewContainer(register types.Register, provider types.Provider) types.Container {
	c := &CoreContainer{
		register: register,
		provider: provider,
		create:   NewWithContainer,
	}

	register.RegisterInterface(&provider)

	return c
}
