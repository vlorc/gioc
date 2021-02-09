// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package dependency

import (
	"github.com/vlorc/gioc/types"
	"reflect"
)

func NewDependencyScan(dep []types.DependencyDescriptor) types.DependencyScan {
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

func (ds *CoreDependencyScan) Flags() types.DependencyFlag {
	return ds.dep[ds.pos].Flags
}

func (ds *CoreDependencyScan) Index() types.Indexer {
	return ds.dep[ds.pos].Index
}

func (ds *CoreDependencyScan) Type() reflect.Type {
	return ds.dep[ds.pos].Type
}

func (ds *CoreDependencyScan) Factory(provider types.Provider) types.BeanFactory {
	return ds.dep[ds.pos].Factory
}
