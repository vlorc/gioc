// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package binder

import (
	"github.com/vlorc/gioc/types"
)

func (sb *SelectorBind) AsMapper() types.Mapper {
	return sb.selector.MapperOf(sb.typ)
}

func (sb *SelectorBind) Bind(key string, factory types.BeanFactory) (err error) {
	return sb.selector.SetFactory(sb.typ, key, factory)
}

func (sb *SelectorBind) Resolve(key string) types.BeanFactory {
	return sb.selector.FactoryOf(sb.typ, key)
}
