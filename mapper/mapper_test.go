// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

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

func test_mapping(t *testing.T, mapper types.Mapper, table map[string]types.BeanFactory) {

	if nil == mapper {
		t.Errorf("can't allocate a mapper")
	}

	for k, v := range table {
		temp := mapper.Resolve(k)
		if nil == temp {
			t.Errorf("can't found key %s",k)
		}
		if v != temp {
			t.Errorf("can't matching key %s,were modified", k)
		}
	}
}

func Test_NamedMapping(t *testing.T) {
	table := map[string]types.BeanFactory{}
	for v, k := range []string{
		"a", "c", "b",
		"0", "1", "2",
	} {
		table[k] = intFactory(v)
	}

	test_mapping(t,NewNamedMapping(table, &sync.Mutex{}), table)
}
