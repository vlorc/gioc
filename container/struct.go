// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package container

import (
	"github.com/vlorc/gioc/types"
)

type CoreContainer struct {
	register types.Register
	provider types.Provider
	create   func(types.Provider, ...string) types.Container
	name     string
	count    uint32
}
