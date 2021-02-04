// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package selector

import (
	"github.com/vlorc/gioc/types"
	"reflect"
)

func (s *readOnlySelector) Get(typ reflect.Type, name string) types.GeneralFactory {
	if nil != s.getter {
		return s.getter.Get(typ, name)
	}
	return nil
}

func (s *readOnlySelector) Range(callback func(types.GeneralFactory) bool, types ...reflect.Type) {

}

func (s *readOnlySelector) Add(typ reflect.Type, name string, factory types.BeanFactory) {

}

func (s *readOnlySelector) Set(typ reflect.Type, name string, factory types.BeanFactory) {

}

func (s *readOnlySelector) Put(typ reflect.Type, name string, factory types.BeanFactory) bool {
	return true
}
