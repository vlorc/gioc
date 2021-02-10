// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package module

import (
	"github.com/vlorc/gioc/types"
	"github.com/vlorc/gioc/utils"
)

func NewWithModule(parent func() types.Container, table ...ModuleInitHandle) types.Module {
	m, err := moduleInit(newWithModule(parent), table...)
	if nil != err {
		utils.Panic(err)
		return nil
	}
	return m
}

func NewModuleFor(container func() types.Container, table ...ModuleInitHandle) types.Module {
	m, err := moduleInit(newModule(container, container), table...)
	if nil != err {
		utils.Panic(err)
		return nil
	}
	return m
}

func NewModuleFactory(table ...ModuleInitHandle) types.ModuleFactory {
	return &CoreModuleFactory{
		table: table,
	}
}

func NewModuleForFactory(table ...ModuleInitHandle) types.ModuleFactory {
	return &CoreModuleForFactory{
		table: table,
	}
}

func (mf *CoreModuleFactory) Instance(parent func() types.Container) (types.Module, error) {
	return moduleInit(newWithModule(parent), mf.table...)
}

func (mff *CoreModuleForFactory) Instance(parent func() types.Container) (types.Module, error) {
	return moduleInit(newModule(parent, parent), mff.table...)
}
