// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package factory

import (
	"github.com/vlorc/gioc/types"
	"reflect"
	"sync"
	"testing"
)

func test_factory_instance(t *testing.T, f types.BeanFactory, p types.Provider) interface{} {
	if nil == f {
		t.Errorf("can't allocate a factory")
	}

	dst, err := f.Instance(p)
	if nil != err {
		t.Errorf("can't allocate a instance error : %s", err.Error())
	}
	return dst
}

func test_factory(t *testing.T, f types.BeanFactory, p types.Provider, src interface{}) {
	dst := test_factory_instance(t, f, p)
	if dst != src {
		t.Errorf("can't matching instance,were modified, %v != %v", dst, src)
	}
}

func Test_ValueFactory(t *testing.T) {
	test_factory(t, NewValueFactory(1), nil, 1)
}

func Test_ProxyFactory(t *testing.T) {
	test_factory(t, NewProxyFactory(NewValueFactory(1)), nil, 1)
}

func Test_FuncFactory(t *testing.T) {
	test_factory(t, NewFuncFactory(func(types.Provider) (interface{}, error) {
		return 1, nil
	}), nil, 1)
}

func Test_TypeFactory(t *testing.T) {
	dst := test_factory_instance(t, NewTypeFactory((*****int)(nil)), nil)
	if _, ok := dst.(*****int); !ok {
		t.Errorf("can't matching or allocate instance")
	}
}

func Test_PointerFactory(t *testing.T) {
	i := 1

	test_factory(t, NewPointerFactory(reflect.ValueOf(&i)), nil, 1)

	i = 2
	test_factory(t, NewPointerFactory(reflect.ValueOf(&i)), nil, 2)
}

func Test_ParamFactory(t *testing.T) {
	num := 10
	dst := test_factory_instance(t, NewParamFactory(num), nil)
	if param, ok := dst.([]reflect.Value); !ok || len(param) != num {
		t.Errorf("can't matching or allocate instance")
	}
}

func Test_MutexFactory(t *testing.T) {
	table := map[string]int{}
	wait := sync.WaitGroup{}

	factory := NewMutexFactory(NewFuncFactory(func(types.Provider) (interface{}, error) {
		i := table["i"]
		i++
		table["i"] = i
		return i, nil
	}))

	num := 50
	wait.Add(num)

	for i := 0; i < num; i++ {
		go func() {
			for i := 0; i < num; i++ {
				test_factory_instance(t, factory, nil)
			}
			wait.Done()
		}()
	}

	wait.Wait()

	if num*num != table["i"] {
		t.Errorf("can't matching instance")
	}
}

func Test_SingleFactory(t *testing.T) {
	value := 1
	all := sync.WaitGroup{}
	wait := sync.WaitGroup{}

	factory := NewSingleFactory(NewFuncFactory(func(types.Provider) (interface{}, error) {
		value++
		return value, nil
	}))

	num := 250
	wait.Add(num)
	all.Add(num)

	for i := 0; i < num; i++ {
		go func() {
			all.Wait()
			if interface{}(2) != test_factory_instance(t, factory, nil) {
				t.Errorf("can't matching instance")
			}
			wait.Done()
		}()
	}

	all.Add(-num)
	wait.Wait()
}

func test_methodFactory(t *testing.T, f interface{}) {
	factory, resultType, err := NewMethodFactory(f, nil, 1)

	if nil != err {
		t.Errorf("can't allocate a factory error : %s", err.Error())
	}
	if nil == factory {
		t.Errorf("can't allocate a factory")
	}
	if nil == resultType {
		t.Errorf("can't matching result type")
	}

	instance := test_factory_instance(t, factory, nil)
	if it, ok := instance.(string); !ok || "test" != it {
		t.Errorf("can't matching instance")
	}
}

func Test_MethodFactory(t *testing.T) {
	test_methodFactory(t, func() (*int, string, error) {
		return nil, "test", nil
	})

	test_methodFactory(t, func(types.Provider) (*int, string, error) {
		return nil, "test", nil
	})

	test_methodFactory(t, func() (*int, string) {
		return nil, "test"
	})
}
