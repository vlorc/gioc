// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package depend

import (
	"github.com/vlorc/gioc/types"
	"github.com/vlorc/gioc/utils"
	"reflect"
)

func NewDependencyInject(scan types.DependencyScan, ref types.Reflect) types.DependencyInject {
	return &CoreDependencyInject{
		scan,
		ref,
	}
}

func (di *CoreDependencyInject) AsValue() reflect.Value {
	src := di.Get(di.DependencyScan)
	return utils.NewOf(src)
}

func (di *CoreDependencyInject) SetValue(v reflect.Value) {
	di.Set(di.DependencyScan, v)
}

func (di *CoreDependencyInject) SetInterface(v interface{}) {
	di.Set(di.DependencyScan, reflect.ValueOf(v))
}

func (di *CoreDependencyInject) Convert(v interface{}) {
	val := reflect.ValueOf(v)
	if val.IsValid() {
		if val.Type() != di.Type() {
			val = val.Convert(di.Type())
		}
	} else {
		val = reflect.Zero(di.Type())
	}
	di.Set(di.DependencyScan, val)
}

func (di *CoreDependencyInject) SubInject(provider types.Provider) types.DependencyInject {
	dst := di.AsValue()
	return di.Depend().AsInject(dst)
}
