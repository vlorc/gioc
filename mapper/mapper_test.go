package mapper

import (
	"github.com/vlorc/gioc/types"
	"sync"
	"testing"
)

type intFactory int

func (i intFactory) Instance(types.Provider) (interface{}, error) {
	return i, nil
}

func test_mapper(t *testing.T, binder types.Binder) {
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

func test_namedMapping(t *testing.T, table map[string]types.BeanFactory) {
	mapper := NewNamedMapping(table, &sync.Mutex{})
	if nil == mapper {
		t.Errorf("can't allocate a mapper")
	}

	for k, v := range table {
		temp, err := mapper.Resolve(k)
		if nil != err {
			t.Errorf("can't found key %s error : %s", k, err.Error())
		}
		if v != temp {
			t.Errorf("can't matching key %s,were modified", k)
		}
	}
}

func Test_BinderFactory(t *testing.T) {
	table := map[string]types.BeanFactory{}
	for v, k := range []string{
		"a", "c", "b",
		"0", "1", "2",
	} {
		table[k] = intFactory(v)
	}
	test_namedMapping(t, table)
}
