// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package binder

import (
	"github.com/vlorc/gioc/types"
)

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

