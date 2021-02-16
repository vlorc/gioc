// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package utils

import (
	"reflect"
)

var __valueType = reflect.TypeOf((*reflect.Value)(nil)).Elem()

// reflect.Type -> reflect.Type
// reflect.Value -> reflect.Type
// (*int)(nil) -> reflect.Value	get int type, must be a pointer
// convert the reflect.Type type
func TypeOf(v interface{}) (t reflect.Type) {
	switch r := v.(type) {
	case reflect.Type:
		t = r
	case reflect.Value:
		t = r.Type()
	default:
		if t = reflect.TypeOf(v); reflect.Ptr == t.Kind() {
			t = t.Elem()
		}
	}
	return t
}

// convert the reflect.Value type
func ValueOf(v interface{}) (t reflect.Value) {
	t, ok := v.(reflect.Value)
	if !ok {
		t = reflect.ValueOf(v)
	}
	return t
}

// returns the type the interface t points of reflect.Type
func IndirectType(t reflect.Type) reflect.Type {
	for reflect.Ptr == t.Kind() {
		t = t.Elem()
	}
	return t
}

// returns the value that v points to
func IndirectValue(v reflect.Value) reflect.Value {
	for reflect.Ptr == v.Kind() {
		v = v.Elem()
	}
	return v
}

// returns an interface type
func InterfaceOf(t reflect.Type) reflect.Type {
	if t = IndirectType(t); reflect.Interface != t.Kind() {
		t = nil
	}
	return t
}

// get a can set value
func NewOf(v reflect.Value) reflect.Value {
	for reflect.Ptr == v.Kind() {
		tmp := reflect.New(v.Type().Elem())
		v.Set(tmp)
		v = tmp.Elem()
	}
	return v
}

// type conversion
func Convert(val reflect.Value, typ reflect.Type) reflect.Value {
	if !val.IsValid() {
		return reflect.Zero(typ)
	}
	if typ == val.Type() {
		return val
	}
	if __valueType == val.Type() {
		return Convert(val.Interface().(reflect.Value), typ)
	}
	return val.Convert(typ)
}

// type conversion imp to typ
func Elem(imp interface{}, typ reflect.Type) interface{} {
	val := ValueOf(imp)
	for ; val.Type() != typ; val = val.Elem() {
		if reflect.Ptr != val.Kind() || val.IsNil() {
			return nil
		}
	}
	return val.Interface()
}

func IsNil(val interface{}) bool {
	v := ValueOf(val)

	return v.IsValid() || ((v.Kind() >= reflect.Chan && v.Kind() <= reflect.Slice) && v.IsNil())
}
