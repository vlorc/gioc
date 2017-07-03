// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package depend

import (
	"github.com/vlorc/gioc/types"
	"reflect"
)

func NewDescriptorGetter(des *types.DependencyDescription) types.DescriptorGetter{
	return &DescriptorGetter{des:des}
}

func NewDescriptorSetter(des *types.DependencyDescription) types.DescriptorSetter{
	return &DescriptorSetter{des:des}
}

func NewDescriptor(des *types.DependencyDescription) types.Descriptor{
	return &Descriptor{
		NewDescriptorGetter(des),
		NewDescriptorSetter(des),
	}
}

func(dg *DescriptorGetter)Type() reflect.Type{
	return dg.des.Type
}

func(dg *DescriptorGetter)Name() string{
	return dg.des.Name
}

func(dg *DescriptorGetter)Default() reflect.Value{
	return dg.des.Default
}

func(dg *DescriptorGetter)Flags() types.DependencyFlag{
	return dg.des.Flags
}

func(dg *DescriptorGetter)Index() int{
	return dg.des.Index
}

func(dg *DescriptorGetter)Depend() types.Dependency{
	return dg.des.Depend
}

func(dg *DescriptorSetter)SetType(typ reflect.Type) {
	dg.des.Type = typ
}

func(dg *DescriptorSetter)SetName(name string) {
	dg.des.Name = name
}

func(dg *DescriptorSetter)SetDefault(def reflect.Value){
	dg.des.Default = def
}

func(dg *DescriptorSetter)SetFlags(flags types.DependencyFlag) {
	dg.des.Flags = flags
}

func(dg *DescriptorSetter)SetIndex(index int) {
	dg.des.Index = index
}

func(dg *DescriptorSetter)SetDepend(depend types.Dependency) {
	dg.des.Depend = depend
}