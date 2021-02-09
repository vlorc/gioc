// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package utils

import (
	"reflect"
	"testing"
)

func Test_TypeOf(t *testing.T) {
	err := "can't matching type"
	if TypeOf((*int)(nil)) != reflect.TypeOf(0) {
		t.Error(err)
	}

	if TypeOf((*int)(nil)) != TypeOf(reflect.TypeOf(0)) {
		t.Error(err)
	}

	if TypeOf((*int)(nil)) != TypeOf(reflect.ValueOf(0)) {
		t.Error(err)
	}
}

func Test_ValueOf(t *testing.T) {
	err := "can't matching type"

	if v := ValueOf(1); v.Interface() != interface{}(1) {
		t.Error(err)
	}

	if v := ValueOf(reflect.ValueOf(1)); v.Interface() != interface{}(1) {
		t.Error(err)
	}
}

func Test_InterfaceOf(t *testing.T) {
	err := "must is interface type"

	if InterfaceOf(reflect.TypeOf(0)) != nil {
		t.Error(err)
	}

	if InterfaceOf(reflect.TypeOf(new(interface{}))) == nil {
		t.Error(err)
	}
}

func Test_IndirectType(t *testing.T) {
	if v := IndirectType(reflect.TypeOf((*int)(nil))); nil == v || reflect.Ptr == v.Kind() {
		t.Error("must not is pointer type")
	}
}

func Test_IndirectValue(t *testing.T) {
	if v := IndirectValue(reflect.ValueOf(new(int))); !v.CanSet() {
		t.Error("can't set value")
	}
}

func Test_NewOf(t *testing.T) {
	var src ****int

	dst := NewOf(reflect.ValueOf(&src).Elem())

	dst.SetInt(1)

	if ****src != 1 {
		t.Error("can't set")
	}
}
