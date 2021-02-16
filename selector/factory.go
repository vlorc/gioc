// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package selector

import (
	"github.com/vlorc/gioc/types"
	"reflect"
)

func (*coreSelectorFactory) Instance() (types.Selector, error) {
	return NewGeneralSelector(), nil
}

func NewSelectorFactory() types.SelectorFactory {
	return &coreSelectorFactory{}
}

func NewReadOnlyFactory(getter types.SelectorGetter) types.Selector {
	return &readOnlySelector{getter}
}

func NewGeneralSelector() types.Selector {
	return &generalSelector{
		pool: make([]types.GeneralFactory, 0, 24),

		primary: map[typeName]int{},

		types: map[reflect.Type][]int{},

		name: map[string]int{},
	}
}
