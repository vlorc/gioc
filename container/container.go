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
		create: func(types.Provider) types.Container{
			return nil
		},
	}
}

func (c *CoreContainer) Readonly() types.Container {
	return &CoreContainer{
		register: nil,
		provider: c.provider,
		create: c.create,
	}
}

func (c *CoreContainer) NewChild() types.Container {
	return c.create(c.AsProvider())
}
