// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package selector

import (
	"github.com/vlorc/gioc/types"
	"reflect"
)

func (sf *CoreSelectorFactory) Instance(factory types.BinderFactory) (selector types.Selector, _ error) {
	if nil == factory {
		selector = NewTypeNameSelector()
	} else {
		selector = NewTypeSelector(factory)
	}
	return
}

func NewSelectorFactory() types.SelectorFactory {
	return &CoreSelectorFactory{}
}

func NewTypeSelector(factory types.BinderFactory) types.Selector {
	return &TypeSelector{
		factory: factory,
		table:   make(map[reflect.Type]types.Binder),
	}
}

func NewTypeNameSelector() types.Selector {
	return &NamedSelector{
		selector: make(typeNameSelector),
	}
}

func NewNamedSelector() types.Selector {
	return &NamedSelector{
		selector: make(nameSelector),
	}
}
