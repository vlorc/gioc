// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package depend

import (
	"github.com/vlorc/gioc/types"
	"github.com/vlorc/gioc/utils"
	"reflect"
)

func NewDependencyScan(dep []*types.DependencyDescription) types.DependencyScan {
	return &CoreDependencyScan{
		dep: dep,
		pos: len(dep),
	}
}

func (ds *CoreDependencyScan) Reset() {
	ds.pos = len(ds.dep)
}

func (ds *CoreDependencyScan) Next() bool {
	ds.pos -= 1
	return ds.pos >= 0
}

func (ds *CoreDependencyScan) Type() reflect.Type {
	return ds.dep[ds.pos].Type
}

func (ds *CoreDependencyScan) Default() reflect.Value {
	return ds.dep[ds.pos].Default()
}

func (ds *CoreDependencyScan) Name() string {
	return ds.dep[ds.pos].Name
}

func (ds *CoreDependencyScan) Flags() types.DependencyFlag {
	return ds.dep[ds.pos].Flags
}

func (ds *CoreDependencyScan) Index() int {
	return ds.dep[ds.pos].Index
}

func (ds *CoreDependencyScan) Depend() types.Dependency {
	return ds.dep[ds.pos].Depend
}

func (ds *CoreDependencyScan) Test(v interface{}) bool {
	srcType := utils.TypeOf(v)
	dstType := ds.dep[ds.pos].Type

	return dstType == srcType ||
		dstType.ConvertibleTo(srcType) ||
		(dstType.Kind() == reflect.Interface && srcType.Implements(dstType))
}

func (ds *CoreDependencyScan) AsDescriptorGetter() types.DescriptorGetter {
	return NewDescriptorGetter(ds.dep[ds.pos])
}

func (ds *CoreDependencyScan) AsDescriptorSetter() types.DescriptorSetter {
	return nil //NewDescriptorSetter(ds.dep[ds.pos])
}
