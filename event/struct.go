// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package event

import (
	"github.com/vlorc/gioc/types"
	"sync"
)

type CoreEventListener struct {
	lock sync.RWMutex
	provider func() types.Provider
	defaults []types.Invoker
	listeners  map[string][]types.Invoker
}