// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package mapper

import (
	"github.com/vlorc/gioc/types"
	"sync"
)

func NewNamedMapping(table map[string]types.BeanFactory, lock sync.Locker) types.Mapper {
	return &NamedMapping{
		lock:  lock,
		table: table,
	}
}

func (nm *NamedMapping) Resolve(key string) (factory types.BeanFactory) {
	nm.lock.Lock()
	factory = nm.table[key]
	nm.lock.Unlock()
	return
}
