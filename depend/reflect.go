// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package depend

import (
	"github.com/vlorc/gioc/types"
	"reflect"
	"unsafe"
)

func NewParamReflect(value reflect.Value) (ref types.Reflect) {
	if value.Len() > 0 {
		ref = ParamReflect(value.Interface().([]reflect.Value))
	}
	return
}

func NewStructReflect(value reflect.Value) (ref types.Reflect) {
	if reflect.Struct == value.Kind() {
		ref = StructReflect(value)
	}
	return
}

func NewArrayReflect(value reflect.Value) (ref types.Reflect) {
	if value.Len() > 0 {
		ref = ArrayReflect(value)
	}
	return
}

func NewMapReflect(value reflect.Value) (ref types.Reflect) {
	if reflect.Map == value.Kind() {
		ref = MapReflect(value)
	}
	return
}

func (pr ParamReflect) Set(des types.DescriptorGetter, val reflect.Value) {
	pr[des.Index()] = val
}

func (pr ParamReflect) Get(des types.DescriptorGetter) reflect.Value {
	val := pr[des.Index()]
	if !val.IsValid() {
		val = reflect.New(des.Type()).Elem()
		pr[des.Index()] = val
	}
	return val
}

func (sr StructReflect) Set(des types.DescriptorGetter, val reflect.Value) {
	sr.Get(des).Set(val)
}

func (sr StructReflect) Get(des types.DescriptorGetter) reflect.Value {
	val := reflect.Value(sr).Field(des.Index())
	if 0 != types.DEPENDENCY_FLAG_UNSAFE&des.Flags() {
		val = reflect.NewAt(val.Type(), unsafe.Pointer(val.UnsafeAddr()))
		val = val.Elem()
	}
	return val

}

func (ar ArrayReflect) Set(des types.DescriptorGetter, val reflect.Value) {
	reflect.Value(ar).Index(des.Index()).Set(val)
}

func (ar ArrayReflect) Get(des types.DescriptorGetter) reflect.Value {
	return reflect.Value(ar).Index(des.Index())
}

func (mr MapReflect) Set(des types.DescriptorGetter, val reflect.Value) {
	reflect.Value(mr).SetMapIndex(reflect.ValueOf(des.Name()), val)
}

func (mr MapReflect) Get(des types.DescriptorGetter) reflect.Value {
	return reflect.Value(mr).MapIndex(reflect.ValueOf(des.Name()))
}
