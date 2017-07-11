// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package builder

import (
	"github.com/vlorc/gioc/types"
)

func (b *CoreBuilder) AsFactory() types.BeanFactory {
	return b
}

func (b *CoreBuilder) Build(provider types.Provider,option ...func(*types.BuildContext)) (interface{}, error) {
	return BuildInstance(provider, b.factory, b.depend,option...)
}

func (b *CoreBuilder) Instance(provider types.Provider) (interface{}, error) {
	return BuildInstance(provider, b.factory, b.depend)
}

