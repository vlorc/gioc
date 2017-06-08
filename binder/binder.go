// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package binder

import (
	"github.com/vlorc/gioc/types"
)

func (nb *NamedBind) AsMapper() types.Mapper {
	//return NewNamedMapping(this.table,this.m.RLocker())
	return nb
}

func (nb *NamedBind) Bind(key string, factory types.BeanFactory) error {
	nb.m.Lock()
	nb.table[key] = factory
	nb.m.Unlock()
	return nil
}

func (nb *NamedBind) Resolve(key string) (factory types.BeanFactory, err error) {
	nb.m.RLock()
	factory = nb.table[key]
	nb.m.RUnlock()
	return
}

type ProxyBind struct {
	read  types.Mapper
	write types.Binder
}

func NewProxyBinder(read types.Mapper, write types.Binder) types.Binder {
	if nil == read {
		read = write.AsMapper()
	}
	return &ProxyBind{
		read:  read,
		write: write,
	}
}

func (pb *ProxyBind) AsMapper() types.Mapper {
	return pb.read
}

func (pb *ProxyBind) Bind(key string, factory types.BeanFactory) (err error) {
	if nil != pb.write {
		err = pb.write.Bind(key, factory)
	}
	return
}

func (pb *ProxyBind) Resolve(key string) (types.BeanFactory, error) {
	return pb.read.Resolve(key)
}
