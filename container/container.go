// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package container

import (
	"github.com/vlorc/gioc/types"
)

func (c *CoreContainer) AsRegister() types.Register {
	return c.register
}

func (c *CoreContainer) AsProvider() types.Provider {
	return c.provider
}

func (c *CoreContainer) Seal() types.Container {
	return &CoreContainer{
		register: c.register,
		provider: c.provider,
		getChild: func() map[types.Container]bool {
			return nil
		},
	}
}

func (c *CoreContainer) Readonly() types.Container {
	return &CoreContainer{
		register: nil,
		provider: c.provider,
		getChild: func() map[types.Container]bool {
			return c.getChild()
		},
	}
}

func (c *CoreContainer) NewChild() types.Container {
	pool := c.getChild()
	if nil == pool {
		return nil
	}

	child := NewWithContainer(c.AsProvider())
	if nil != child {
		pool[child] = true
	}
	return child
}
