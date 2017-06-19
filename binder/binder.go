// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package binder

import (
	"github.com/vlorc/gioc/types"
)

func (nb *NamedBind) AsMapper() types.Mapper {
	return nb
}

func (nb *NamedBind) Bind(key string, factory types.BeanFactory) error {
	nb.lock.Lock()
	nb.table[key] = factory
	nb.lock.Unlock()
	return nil
}

func (nb *NamedBind) Resolve(key string) (factory types.BeanFactory, err error) {
	nb.lock.RLock()
	factory = nb.table[key]
	nb.lock.RUnlock()
	return
}
