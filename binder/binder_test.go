// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package binder

import (
	"github.com/vlorc/gioc/types"
	"reflect"
	"testing"
)

type intFactory int

func (i intFactory) Instance(types.Provider) (interface{}, error) {
	return i, nil
}

func test_binder(t *testing.T, binder types.Binder) {
	if nil == binder {
		t.Errorf("can't allocate a binder")
	}
	key := []string{
		"a", "c", "b",
		"0", "1", "2",
	}

	for v, k := range key {
		err := binder.Bind(k, intFactory(v))
		if nil != err {
			t.Errorf("can't bind key %s error : %s", v, err.Error())
		}
	}

	for v, k := range key {
		temp := binder.Resolve(k)
		if nil == temp {
			t.Errorf("can't found key %s", k)
		}
		if temp != interface{}(intFactory(v)) {
			t.Errorf("can't matching key %s,were modified", k)
		}
	}
}

func test_binderFactory(t *testing.T, factory types.BinderFactory) {
	if nil == factory {
		t.Errorf("can't allocate a factory")
	}

	bind, err := factory.Instance(reflect.TypeOf(0))
	if nil != err {
		t.Errorf("can't allocate a binder error : %s", err.Error())
	}
	test_binder(t, bind)
}

func Test_NameBinder(t *testing.T) {
	test_binder(t, NewNameBinder())
}

func Test_ProxyBinder(t *testing.T) {
	test_binder(t, NewProxyBinder(nil, NewNameBinder()))
}

func Test_BinderFactory(t *testing.T) {
	test_binderFactory(t, NewBinderFactory())
}
