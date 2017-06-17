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

	for i, v := range key {
		err := binder.Bind(v, intFactory(i))
		if nil != err {
			t.Errorf("can't bind key %s error : %s", v, err.Error())
		}
	}

	for i, v := range key {
		temp, err := binder.Resolve(v)
		if nil != err {
			t.Errorf("can't found key %s error : %s", v, err.Error())
		}
		if temp != interface{}(intFactory(i)) {
			t.Errorf("can't matching key %s,were modified", v)
		}
	}
}

func test_binderFactory(t *testing.T, factory types.BinderFactory) {
	if nil == factory {
		t.Errorf("can't allocate a factory")
	}

	binder, err := factory.Instance(reflect.TypeOf(0))
	if nil != err {
		t.Errorf("can't allocate a binder error : %s", err.Error())
	}
	test_binder(t, binder)
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
