// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package mapper

import (
	"github.com/vlorc/gioc/types"
	"sync"
)

type NamedMapping struct {
	lock     sync.Locker
	table map[string]types.BeanFactory
}

func NewNamedMapping(t map[string]types.BeanFactory, lock sync.Locker) types.Mapper {
	return &NamedMapping{
		lock:    lock,
		table: t,
	}
}

func (nm *NamedMapping) Resolve(v string) (factory types.BeanFactory, err error) {
	nm.lock.Lock()
	factory = nm.table[v]
	nm.lock.Unlock()
	return
}
