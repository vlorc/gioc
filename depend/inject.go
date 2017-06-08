// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package depend

import (
	"github.com/vlorc/gioc/types"
	"reflect"
)

func (di *CoreDependencyInject) SetInterface(v interface{}) (err error) {
	/*if !di.Test(v) {
		err = types.NewError(types.ErrTypeNotMatch, v)
		return
	}*/

	di.set(di.DependencyScan, di.data, reflect.ValueOf(v))
	return
}

func (di *CoreDependencyInject) SubInject(provider types.Provider) types.DependencyInject {
	src := di.get(di.DependencyScan, di.data)

	for reflect.Ptr == src.Kind() {
		t := reflect.New(src.Type().Elem())
		src.Set(t)
		src = t
	}
	return di.Depend().AsInject(src)
}
