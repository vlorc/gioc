// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package dependency

import (
	"reflect"
	"unsafe"

	"github.com/vlorc/gioc/types"
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

func (pr ParamReflect) Set(des types.Indexer, val reflect.Value) {
	pr[des.Value()] = val
}

func (pr ParamReflect) Get(des types.Indexer) reflect.Value {
	return pr[des.Value()]
}

func (sr StructReflect) Set(des types.Indexer, val reflect.Value) {
	sr.Get(des).Set(val)
}

func (sr StructReflect) Get(des types.Indexer) reflect.Value {
	target := reflect.Value(sr)

	field := target.Type().Field(des.Value())
	val := target.Field(des.Value())

	if field.Name[0] < 'A' || field.Name[0] > 'Z' {
		val = reflect.NewAt(val.Type(), unsafe.Pointer(val.UnsafeAddr()))
		val = val.Elem()
	}
	return val

}

func (ar ArrayReflect) Set(des types.Indexer, val reflect.Value) {
	reflect.Value(ar).Index(des.Value()).Set(val)
}

func (ar ArrayReflect) Get(des types.Indexer) reflect.Value {
	return reflect.Value(ar).Index(des.Value())
}

func (mr MapReflect) Set(des types.Indexer, val reflect.Value) {
	reflect.Value(mr).SetMapIndex(reflect.ValueOf(des.String()), val)
}

func (mr MapReflect) Get(des types.Indexer) reflect.Value {
	return reflect.Value(mr).MapIndex(reflect.ValueOf(des.String()))
}
