// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package factory

import (
	"github.com/vlorc/gioc/types"
	"github.com/vlorc/gioc/utils"
	"reflect"
)

func (f *resolveAnyFactory) Instance(provider types.Provider) (interface{}, error) {
	val := reflect.MakeSlice(reflect.SliceOf(f.typ), 0, 8)

	provider.Range(func(factory types.GeneralFactory) bool {
		v, err := factory.Instance(provider)
		if nil == err {
			val = reflect.Append(val, utils.Convert(reflect.ValueOf(v), f.typ))
		}

		// add check error
		return true
	}, f.typ)

	return val.Interface(), nil
}

func (f *resolveNamesFactory) Instance(provider types.Provider) (interface{}, error) {
	val := reflect.MakeSlice(reflect.SliceOf(f.typ), 0, len(f.name))

	for _, b := range f.name {
		id, err := b.Instance(provider)
		if nil != err {
			// add check error
			continue
		}

		v, err := provider.Get(f.typ, id)
		if nil != err {
			// add check error
			continue
		}

		val = reflect.Append(val, utils.Convert(reflect.ValueOf(v), f.typ))
	}

	return val.Interface(), nil
}
