// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package dependency

import (
	"fmt"
	"github.com/vlorc/gioc/types"
	"github.com/vlorc/gioc/utils"
	"reflect"
	"strings"
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

func (d *CoreDependency) AsReflect(imp interface{}) types.Reflect {
	val := utils.IndirectValue(utils.ValueOf(imp))

	return d.factory(val)
}

func (d *CoreDependency) String() string {
	b := strings.Builder{}

	fmt.Fprintf(&b, "type(%s)", d.typ.String())

	for i := range d.dep {
		fmt.Fprintf(&b, " descriptor(%v)", d.dep[i].Factory)
	}

	return b.String()
}
