// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package mapper

import (
	"github.com/vlorc/gioc/types"
	"sync"
)

type NamedMapping struct {
	m     sync.Locker
	table map[string]types.BeanFactory
}

func NewNamedMapping(t map[string]types.BeanFactory, m sync.Locker) types.Mapper {
	return &NamedMapping{
		m:     m,
		table: t,
	}
}

func (nm *NamedMapping) Resolve(v string) (factory types.BeanFactory, err error) {
	nm.m.Lock()
	factory = nm.table[v]
	nm.m.Unlock()
	return
}
