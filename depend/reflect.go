// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package depend

import (
	"reflect"
	"github.com/vlorc/gioc/types"
)

func NewParamReflect(value reflect.Value) (ref types.Reflect){
	if value.Len() > 0 {
		ref =  CoreParamReflect(value.Interface().([]reflect.Value))
	}
	return
}

func NewStructReflect(value reflect.Value) (ref types.Reflect){
	if reflect.Struct == value.Kind() {
		ref = CoreStructReflect(value)
	}
	return
}

func NewArrayReflect(value reflect.Value) (ref types.Reflect){
	if value.Len() > 0 {
		ref =  CoreArrayReflect(value)
	}
	return
}

func NewMapReflect(value reflect.Value) (ref types.Reflect){
	if reflect.Map == value.Kind() {
		ref =  CoreMapReflect(value)
	}
	return
}

func(pr CoreParamReflect)Set(des types.PropertyDescriptorGetter,val reflect.Value){
	pr[des.Index()] = val
}

func(pr CoreParamReflect)Get(des types.PropertyDescriptorGetter)reflect.Value{
	return pr[des.Index()]
}

func(sr CoreStructReflect)Set(des types.PropertyDescriptorGetter,val reflect.Value){
	reflect.Value(sr).Field(des.Index()).Set(val)
}

func(sr CoreStructReflect)Get(des types.PropertyDescriptorGetter)reflect.Value{
	return reflect.Value(sr).Field(des.Index())
}

func(ar CoreArrayReflect)Set(des types.PropertyDescriptorGetter,val reflect.Value){
	reflect.Value(ar).Index(des.Index()).Set(val)
}

func(ar CoreArrayReflect)Get(des types.PropertyDescriptorGetter)reflect.Value{
	return reflect.Value(ar).Index(des.Index())
}

func(mr CoreMapReflect)Set(des types.PropertyDescriptorGetter,val reflect.Value){
	reflect.Value(mr).SetMapIndex(reflect.ValueOf(des.Name()),val)
}

func(mr CoreMapReflect)Get(des types.PropertyDescriptorGetter)reflect.Value{
	return reflect.Value(mr).MapIndex(reflect.ValueOf(des.Name()))
}