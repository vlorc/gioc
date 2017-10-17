// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package builder

import (
	"github.com/vlorc/gioc/types"
)

type CoreBuilderFactory struct{}

type CoreBuilder struct {
	depend  types.Dependency
	factory types.BeanFactory
}
