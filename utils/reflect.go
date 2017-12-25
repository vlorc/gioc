// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package utils

import (
	"reflect"
)

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

// skip all pointer of reflect.Type
func DirectlyType(t reflect.Type) reflect.Type {
	for reflect.Ptr == t.Kind() {
		t = t.Elem()
	}
	return t
}

// skip all pointer of reflect.Value
func DirectlyValue(v reflect.Value) reflect.Value {
	for reflect.Ptr == v.Kind() {
		v = v.Elem()
	}
	return v
}

// get an interface type
func InterfaceOf(t reflect.Type) reflect.Type {
	if t = DirectlyType(t); reflect.Interface != t.Kind() {
		t = nil
	}
	return t
}

// get a can set value
func NewOf(src reflect.Value) reflect.Value {
	for reflect.Ptr == src.Kind() {
		tmp := reflect.New(src.Type().Elem())
		src.Set(tmp)
		src = tmp.Elem()
	}
	return src
}

func Convert(src reflect.Value, typ reflect.Type) reflect.Value {
	if src.IsValid() {
		if src.Type() != typ {
			src = src.Convert(typ)
	}
	} else {
		src = reflect.Zero(typ)
	}
	return src
}
