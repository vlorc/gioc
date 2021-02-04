// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package selector

import (
	"github.com/vlorc/gioc/types"
	"reflect"
	"sync"
)

type generalSelector struct {
	mux sync.RWMutex

	pool []types.GeneralFactory

	primary map[typeName]int

	types map[reflect.Type][]int

	name map[string]int
}

type readOnlySelector struct {
	getter types.SelectorGetter
}

type beanFactory struct {
	factory types.BeanFactory
	typ     reflect.Type
	name    string
}

type typeName struct {
	typ  reflect.Type
	name string
}

type coreSelectorFactory struct{}

var _ types.GeneralFactory = &beanFactory{}

func (f *beanFactory) Type() reflect.Type {
	return f.typ
}
func (f *beanFactory) Name() string {
	return f.name
}
func (f *beanFactory) Instance(provider types.Provider) (interface{}, error) {
	return f.factory.Instance(provider)
}
