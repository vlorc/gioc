// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package register

import (
	"github.com/vlorc/gioc/types"
)

type coreRegister struct {
	selector types.Selector
}

type readOnlyRegister struct {
	getter types.SelectorGetter
}

type coreRegisterFactory struct{}
