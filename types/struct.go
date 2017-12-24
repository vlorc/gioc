// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package types

import "reflect"

type Error struct {
	Code   ErrorCode
	Args   []interface{}
	format func(*Error) string
}

type DependencyDescription struct {
	Type    reflect.Type
	Name    string
	Index   int
	Flags   DependencyFlag
	Default reflect.Value
	Depend  Dependency
}

type BuildContext struct {
	Descriptor DescriptorGetter
	Inject     DependencyInject
	Provider   Provider
	FullBefore func(*BuildContext) bool
	FullAfter  func(*BuildContext)
}

type ParseContext struct {
	Factory    DependencyFactory
	Descriptor Descriptor
	Params     []Param
	Scan       TokenScan
	Dump       func(int, int) string
}

type IdentHandle func(*ParseContext) error
