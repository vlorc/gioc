// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package builder

import (
	"github.com/vlorc/gioc/types"
)

func (bf *CoreBuilderFactory) Instance(factory types.BeanFactory, depend types.Dependency) (types.Builder, error) {
	return &CoreBuilder{
		depend:  depend,
		factory: factory,
	}, nil
}

func NewBuilderFactory() types.BuilderFactory {
	return &CoreBuilderFactory{}
}
