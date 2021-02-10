// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package types

import "reflect"

type Error struct {
	Code   ErrorCode
	Args   []interface{}
	format func(*Error) string
}

type DependencyOrigin struct {
	Target reflect.Type
	Type   reflect.Type
	Name   string
	Index  Indexer
}

type DependencyDescriptor struct {
	Factory BeanFactory
	Index   Indexer
	Flags   DependencyFlag
	Type    reflect.Type
}

type DependencyContext struct {
	Origin DependencyOrigin

	Type  reflect.Type
	Name  []StringFactory
	Flags DependencyFlag

	Default func(*DependencyContext) BeanFactory

	Before []func(*DependencyContext) error

	After []func(*DependencyContext) error

	Wrapper WrapperFactory

	Factory BeanFactory

	Error error
}

type ParseContext struct {
	Factory DependencyFactory

	Dependency DependencyContext

	Params []Param
	Scan   TokenScan
	Dump   func(int, int) string
}

type IdentHandle func(*ParseContext) error

func (c *DependencyContext) Reset() {
	c.Type = nil
	c.Name = nil
	c.Flags = 0
	c.Factory = nil

	c.Origin.Target = nil
	c.Origin.Type = nil
	c.Origin.Name = ""
	c.Origin.Index = nil

	c.Default = nil
	c.After = nil
	c.Wrapper.Reset()
}
