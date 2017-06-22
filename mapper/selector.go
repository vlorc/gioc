// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package mapper

import (
	"github.com/vlorc/gioc/types"
	"reflect"
)

func NewSelectorMapping(typ reflect.Type,selector types.SelectorGetter) types.Mapper {
	return &SelectorMapping{
		typ:  typ,
		selector: selector,
	}
}

func (sm *SelectorMapping) Resolve(key string) types.BeanFactory {
	return sm.selector.FactoryOf(sm.typ,key)
}
