// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package container

import (
	"github.com/vlorc/gioc/types"
)

func (c *CoreContainer) AsRegister() types.Register {
	return c.Register
}

func (c *CoreContainer) AsProvider() types.Provider {
	return c.Provider
}

func (c *CoreContainer) Seal() types.Container {
	return c
}

func (c *CoreContainer) Readonly() types.Container {
	return c
}

func (c *CoreContainer) Parent() types.Container {
	return c.parent()
}

func (c *CoreContainer) Child() types.Container {
	child := NewChildContainer(c, c,0)
	if nil != child {
		c.getChild()[child] = true
	}
	return child
}
