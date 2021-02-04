// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package provider

import (
	"github.com/vlorc/gioc/types"
)

func NewWithProvider(parent types.Provider, selector types.SelectorGetter) types.Provider {
	return &coreProvider{
		parent:   parent,
		selector: selector,
	}
}

func NewProviderFactory() types.ProviderFactory {
	return &coreProviderFactory{}
}

func (*coreProviderFactory) Instance(parent types.Provider, selector types.SelectorGetter) (types.Provider, error) {
	return NewWithProvider(parent, selector), nil
}
