// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package container

import (
	"github.com/vlorc/gioc/types"
	"github.com/vlorc/gioc/utils"
	"strings"
)

func NewWithContainer(provider types.Provider, names ...string) types.Container {
	var selectorFactory types.SelectorFactory
	var registerFactory types.RegisterFactory
	var providerFactory types.ProviderFactory
	provider.Load(&selectorFactory)
	provider.Load(&registerFactory)
	provider.Load(&providerFactory)

	if nil == selectorFactory || nil == registerFactory || nil == providerFactory {
		return nil
	}
	sel, err := selectorFactory.Instance()
	if nil != err {
		utils.Panic(err)
	}
	reg, err := registerFactory.Instance(sel)
	if nil != err {
		utils.Panic(err)
	}
	pro, err := providerFactory.Instance(provider, sel)
	if nil != err {
		utils.Panic(err)
	}
	return NewContainer(reg, pro, names...)
}

func NewContainer(register types.Register, provider types.Provider, names ...string) types.Container {
	c := &CoreContainer{
		register: register,
		provider: provider,
		create:   NewWithContainer,
	}

	register.Interface(&register)
	register.Interface(&provider)

	if len(names) == 0 {
		return c
	}

	c.name = names[0]
	r := register
	if pos := strings.Index(c.name, "::"); pos > 0 {
		provider.Load(&r, c.name[:pos])
	}
	if nil != r {
		r.Interface(&register, c.name)
		r.Interface(&provider, c.name)
	}

	return c
}
