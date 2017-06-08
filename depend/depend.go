// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package depend

import (
	"github.com/vlorc/gioc/types"
	"github.com/vlorc/gioc/utils"
	"reflect"
)

func (d *CoreDependency) Type() reflect.Type {
	return d.typ
}

func (d *CoreDependency) Length() int {
	return len(d.dep)
}

func (d *CoreDependency) AsScan() types.DependencyScan {
	return NewDependencyScan(d.dep)
}

func (d *CoreDependency) AsInject(v interface{}) types.DependencyInject {
	return d.injectFactory(d.AsScan(), utils.DirectlyValue(utils.ValueOf(v)))
}
