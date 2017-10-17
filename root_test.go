// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package gioc

import (
	"fmt"
	"github.com/vlorc/gioc/binder"
	"github.com/vlorc/gioc/factory"
	"github.com/vlorc/gioc/register"
	"github.com/vlorc/gioc/selector"
	"github.com/vlorc/gioc/types"
	"reflect"
	"testing"
)

func test_register(t *testing.T, r types.Register) {
	if nil == r {
		t.Errorf("can't allocate a Register")
	}

	err := r.RegisterInstance(1, "id")
	if nil != err {
		t.Errorf("can't register a int error : %s", err.Error())
	}

	iface, err := r.AsSelector().FactoryOf(reflect.TypeOf((*int)(nil)).Elem(), "id").Instance(nil)
	if nil != err {
		t.Errorf("can't get a int error : %s", err.Error())
	}
	if iface != interface{}(1) {
		t.Errorf("can't matching instance,were modified")
	}
}

func test_registerFactory(t *testing.T, f types.RegisterFactory) {
	if nil == f {
		t.Errorf("can't allocate a RegisterFactory")
	}

	r, err := f.Instance(selector.NewTypeNameSelector())
	if nil != err {
		t.Errorf("can't allocate a Register error : %s", err.Error())
	}
	test_register(t, r)
}

func Test_Register(t *testing.T) {
	test_register(t, register.NewRegister(selector.NewTypeSelector(binder.NewBinderFactory())))
}

func Test_RegisterFactory(t *testing.T) {
	test_registerFactory(t, register.NewRegisterFactory())
}

func Test_Invoker(t *testing.T) {
	root := NewRootContainer()

	getKey := func(id int64, name *string) (r string) {
		if nil != name {
			r = fmt.Sprintf("id(%d) - name(%s)", id, *name)
		} else {
			r = fmt.Sprintf("id(%d)", id)
		}
		return
	}

	name := "angel"
	root.AsRegister().RegisterInstance(&name)
	var dependFactory types.DependencyFactory
	var builderFactory types.BuilderFactory
	var invokerFactory types.InvokerFactory
	root.Assign(&dependFactory)
	root.Assign(&builderFactory)
	root.Assign(&invokerFactory)

	dep, err := dependFactory.Instance(getKey)
	if nil != err {
		t.Errorf("can't allocate a depend error : %s", err.Error())
	}
	build, err := builderFactory.Instance(factory.NewParamFactory(dep.Length()), dep)
	if nil != err {
		t.Errorf("can't allocate a build error : %s", err.Error())
	}
	invoker, err := invokerFactory.Instance(getKey, build)
	if nil != err {
		t.Errorf("can't allocate a invoker error : %s", err.Error())
	}

	results := invoker.ApplyWith(root.AsProvider(), 1)
	t.Log("getKey", results[0].Interface())
	results = invoker.ApplyWith(root.AsProvider(), -2, nil)
	t.Log("getKey", results[0].Interface())
}
