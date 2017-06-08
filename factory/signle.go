// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package factory

import (
	"github.com/vlorc/gioc/types"
	"sync"
)

func singleFactoryOf(src types.BeanFactory) types.BeanFactory {
	dst := &ProxyFactory{}
	once := &sync.Once{}
	dst.factory = NewFuncFactory(func(provider types.Provider) (interface{}, error) {
		once.Do(func() {
			dst.factory = NewValueFactory(src.Instance(provider))
		})
		return dst.Instance(provider)
	})
	return dst
}
