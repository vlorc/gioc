// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package provider

import (
	"github.com/vlorc/gioc/types"
)

func NewProxyProvider(provider types.Provider) types.Provider {
	return &ProxyProvider{
		provider: provider,
	}
}

func NewWithProvider(selector types.SelectorGetter,provider types.Provider) types.Provider {
	return &CoreProvider{
		deep:0,
		parent:provider,
		selector:selector,
	}
}

func NewProviderFactory() types.ProviderFactory {
	return &CoreProviderFactory{}
}

func (fi *CoreProviderFactory) Instance(selector types.SelectorGetter,provider types.Provider) (types.Provider, error) {
	return NewWithProvider(selector,provider),nil
}
