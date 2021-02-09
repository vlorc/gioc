// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package factory

import (
	"fmt"
	"github.com/vlorc/gioc/types"
	"reflect"
)

func (f *DependencyFactory) Instance(provider types.Provider) (interface{}, error) {
	instance, err := f.factory.Instance(provider)
	if nil != err {
		return nil, err
	}

	if err = f.Inject(provider, instance); nil != err {
		return nil, err
	}

	for _, v := range f.after {
		instance = v(instance)
	}
	return instance, nil
}

func (f *DependencyFactory) Inject(provider types.Provider, instance interface{}) error {
	ref := f.dependency.AsReflect(instance)

	for scan := f.dependency.AsScan(); scan.Next(); {
		factory := scan.Factory(provider)

		val, err := factory.Instance(provider)
		if nil != err && 0 == scan.Flags()&types.DEPENDENCY_FLAG_OPTIONAL {
			return err
		}

		ref.Set(scan.Index(), reflect.ValueOf(val))
	}
	return nil
}

func (f *DependencyFactory) String() string {
	return fmt.Sprintf("factory(%v) dependency(%v)", f.factory, f.dependency)
}
