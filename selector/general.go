// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package selector

import (
	"github.com/vlorc/gioc/types"
	"reflect"
)

func (s *generalSelector) getByType(typ reflect.Type) types.GeneralFactory {
	if nil == typ {
		return nil
	}

	s.mux.RLock()
	index := s.types[typ]
	s.mux.RUnlock()

	if len(index) > 0 {
		return s.pool[index[0]]
	}
	return nil
}

func (s *generalSelector) getByName(name string) types.GeneralFactory {
	s.mux.RLock()
	index, ok := s.name[name]
	s.mux.RUnlock()

	if ok {
		return s.pool[index]
	}
	return nil
}

func (s *generalSelector) Get(typ reflect.Type, name string) types.GeneralFactory {
	if "" == name {
		return s.getByType(typ)
	}
	if nil == typ {
		return s.getByName(name)
	}

	s.mux.RLock()
	index, ok := s.primary[typeName{typ, name}]
	s.mux.RUnlock()

	if ok {
		return s.pool[index]
	}
	return nil
}

func (s *generalSelector) rangeWithPool(callback func(types.GeneralFactory) bool) {
	pool := s.pool

	for _, b := range pool {
		if !callback(b) {
			break
		}
	}
}

func (s *generalSelector) rangeByType(callback func(types.GeneralFactory) bool, typ reflect.Type) bool {
	s.mux.RLock()
	index := s.types[typ]
	s.mux.RUnlock()

	pool := s.pool
	for _, i := range index {
		if !callback(pool[i]) {
			return false
		}
	}
	return true
}

func (s *generalSelector) Range(callback func(types.GeneralFactory) bool, types ...reflect.Type) {
	if len(types) <= 0 {
		s.rangeWithPool(callback)
		return
	}
	for _, typ := range types {
		if nil != typ && !s.rangeByType(callback, typ) {
			break
		}
	}
}

func (s *generalSelector) setFactory(factory types.GeneralFactory, name string, typ reflect.Type, insert bool) {
	index := len(s.pool)
	s.pool = append(s.pool, factory)

	s.primary[typeName{typ, name}] = index

	if idx, ok := s.types[typ]; ok {
		if insert {
			s.types[typ] = append([]int{index}, idx...)
		} else {
			s.types[typ] = append(idx, index)
		}
	} else {
		s.types[typ] = []int{index}
	}

	if "" != name {
		s.name[name] = index
	}
	return
}

func (s *generalSelector) Add(typ reflect.Type, name string, factory types.BeanFactory) {
	b := &beanFactory{factory, typ, name}

	s.mux.Lock()
	defer s.mux.Unlock()

	s.setFactory(b, name, typ, false)
}

func (s *generalSelector) Set(typ reflect.Type, name string, factory types.BeanFactory) {
	b := &beanFactory{factory, typ, name}

	s.mux.Lock()
	defer s.mux.Unlock()

	s.setFactory(b, name, typ, true)
}

func (s *generalSelector) Put(typ reflect.Type, name string, factory types.BeanFactory) bool {
	b := &beanFactory{factory, typ, name}

	s.mux.Lock()
	defer s.mux.Unlock()

	if _, ok := s.primary[typeName{typ, name}]; ok {
		return false
	}
	if "" != name {
		if _, ok := s.name[name]; ok {
			return false
		}
	}

	s.setFactory(b, name, typ, false)
	return true
}
