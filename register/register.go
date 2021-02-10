// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package register

import (
	"github.com/vlorc/gioc/factory"
	"github.com/vlorc/gioc/types"
	"github.com/vlorc/gioc/utils"
	"reflect"
)

func (r *coreRegister) Interface(instance interface{}, args ...string) {
	val := utils.IndirectValue(utils.ValueOf(instance))
	if reflect.Interface != val.Kind() || val.IsNil() {
		return
	}
	name := ""
	if len(args) > 0 {
		name = args[0]
	}
	r.selector.Add(val.Type(), name, factory.NewValueFactory(val.Interface()))
}

func (r *coreRegister) Set(instance interface{}, args ...string) {
	t := utils.TypeOf(instance)
	b := factory.NewValueFactory(instance)
	if len(args) > 0 {
		r.selector.Set(t, args[0], b)
	} else {
		r.selector.Set(t, "", b)
	}
}

func (r *coreRegister) Add(instance interface{}, args ...string) {
	t := utils.TypeOf(instance)
	b := factory.NewValueFactory(instance)
	if len(args) > 0 {
		r.selector.Add(t, args[0], b)
	} else {
		r.selector.Add(t, "", b)
	}
}

func (r *coreRegister) Put(instance interface{}, args ...string) bool {
	t := utils.TypeOf(instance)
	b := factory.NewValueFactory(instance)
	if len(args) > 0 {
		return r.selector.Put(t, args[0], b)
	} else {
		return r.selector.Put(t, "", b)
	}
}

func (r *coreRegister) Factory(factory types.BeanFactory, impType interface{}, args ...string) {
	name := ""
	if len(args) > 0 {
		name = args[0]
	}
	r.selector.Add(utils.TypeOf(impType), name, factory)
}

func (r *coreRegister) Selector() types.Selector {
	return r.selector
}
