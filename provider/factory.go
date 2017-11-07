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

func NewWithProvider(selector types.SelectorGetter) types.Provider {
	return &CoreProvider{
		deep:0,
		parent:nil,
		selector:selector,
	}
}