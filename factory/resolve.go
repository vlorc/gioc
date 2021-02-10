// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package factory

import (
	"fmt"
	"github.com/vlorc/gioc/types"
)

func (f *resolveFactory) Instance(provider types.Provider) (instance interface{}, err error) {
	if len(f.name) <= 0 {
		return provider.Get(f.typ)
	}

	var id string
	for _, b := range f.name {
		if id, err = b.Instance(provider); nil != err {
			break
		}
		if instance, err = provider.Get(f.typ, id); nil == err {
			return instance, nil
		}
	}

	return nil, err
}

func (f *resolveFactory) String() string {
	return fmt.Sprintf("type(%s) name(%v)", f.typ.String(), f.name)
}

func (f *typeResolveFactory) Instance(provider types.Provider) (interface{}, error) {
	return provider.Get(f.typ, f.name)
}

func (f *typeResolveFactory) String() string {
	return fmt.Sprintf("type(%s) name(%s)", f.typ.String(), f.name)
}
