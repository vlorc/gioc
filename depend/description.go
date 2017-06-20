package depend

import (
	"github.com/vlorc/gioc/types"
	"reflect"
)

func NewDescriptorGetter(des *DependencyDescription) types.PropertyDescriptorGetter{
	return &DescriptorGetter{des:des}
}

func NewDescriptorSetter(des *DependencyDescription) types.PropertyDescriptorSetter{
	return &DescriptorSetter{des:des}
}

func NewDescriptor(des *DependencyDescription) types.PropertyDescriptor{
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

func(dg *DescriptorGetter)Default() interface{}{
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

func(dg *DescriptorSetter)SetDefault(def interface{}){
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