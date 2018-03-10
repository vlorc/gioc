// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package event

import "github.com/vlorc/gioc/types"

func NewEventListener(provider func() types.Provider) types.EventListener {
	return &CoreEventListener{
		provider:  provider,
		listeners: make(map[string][]types.Invoker),
	}
}

func NewEventListenerWith(provider types.Provider) types.EventListener {
	return NewEventListener(func() types.Provider {
		return provider
	})
}
