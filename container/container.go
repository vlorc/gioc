// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package container

import (
	"fmt"
	"github.com/vlorc/gioc/register"
	"github.com/vlorc/gioc/types"
	"strings"
	"sync/atomic"
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
		create: func(types.Provider, ...string) types.Container {
			return nil
		},
	}
}

func (c *CoreContainer) Readonly() types.Container {
	return &CoreContainer{
		register: register.NewReadOnlyRegister(c.provider.Selector()),
		provider: c.provider,
		create:   c.create,
	}
}

func (c *CoreContainer) NewChild(names ...string) types.Container {
	if "" == c.name {
		return c.create(c.AsProvider())
	}
	var name string
	if len(names) > 0 {
		name = names[0]
	} else {
		name = fmt.Sprintf("child-%d", atomic.AddUint32(&c.count, 1))
	}
	return c.create(c.AsProvider(), strings.Join([]string{c.name, name}, "::"))
}

func (c *CoreContainer) Name() string {
	return c.name
}
