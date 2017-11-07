// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package container

import (
	"github.com/vlorc/gioc/types"
	"github.com/vlorc/gioc/utils"
)

func NewWithContainer(provider types.Provider, parent types.Container, deep int) types.Container {
	var binderFactory types.BinderFactory
	var selectorFactory types.SelectorFactory
	var registerFactory types.RegisterFactory

	provider.Assign(&binderFactory)
	provider.Assign(&selectorFactory)
	provider.Assign(&registerFactory)

	if nil == selectorFactory || nil == registerFactory {
		return nil
	}
	sel, err := selectorFactory.Instance(binderFactory)
	if nil != err {
		panic(err)
	}
	reg, err := registerFactory.Instance(sel)
	if nil != err {
		panic(err)
	}
	return NewContainer(reg, parent, deep)
}

func NewChildContainer(provider types.Provider, parent types.Container, deep int) (c types.Container) {
	var reg types.Register
	if err := provider.Assign(&reg); nil != err {
		c = NewWithContainer(provider, parent, deep)
	} else {
		c = NewContainer(reg, parent, deep)
	}
	return
}

func NewContainer(register types.Register, parent types.Container, deep int) types.Container {
	c := &CoreContainer{
		Register: register,
		parent: func() types.Container {
			return parent
		},
	}

	utils.Once(&c.getChild, func() interface {}{
		pool := make(map[types.Container]bool)
		return func() map[types.Container]bool {
			return pool
		}
	})
	p := c.AsProvider()
	register.RegisterInstance(&p)
	return c
}

