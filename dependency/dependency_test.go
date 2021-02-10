// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package dependency

import (
	"github.com/vlorc/gioc/types"
	"github.com/vlorc/gioc/utils"
	"reflect"
	"testing"
)

func Test_NewDependencyFactory(t *testing.T) {
	test_factory(t, NewDependencyFactory())
}

func test_factory(t *testing.T, factory types.DependencyFactory) {
	if nil == factory {
		t.Errorf("can't allocate a DependencyFactory")
	}
}

func test_struct(t *testing.T, typ reflect.Type) {
	factory := NewDependencyFactory()
	test_factory(t, factory)

	dep, err := factory.Instance(typ)
	if nil != err {
		t.Errorf("can't instance dependency error: %s", err.Error())
	}
	if typ != dep.Type() {
		t.Errorf("can't matching dependency type: %s", dep.Type().String())
	}
	if typ.NumField() != dep.Length() {
		t.Errorf("can't matching dependency length: %d", dep.Length())
	}
	for scan := dep.AsScan(); scan.Next(); {
		t.Logf("%v", scan.Factory(nil))
	}
}

func Test_AnonymousStruct(t *testing.T) {
	type identity struct {
		Username string `inject:"lower"`
		Password string `inject:"upper"`
	}
	type bean struct {
		Id       func() int64 `inject:"lazy"`
		Flags    uint         `inject:"default(10)"`
		Name     *string      `inject:"'alias' optional default(nil)"`
		Identity identity     `inject:"extends"`
	}
	b := bean{}
	typ := utils.IndirectType(reflect.TypeOf(b))
	test_struct(t, typ)
}
