// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package invoker

import (
	"github.com/vlorc/gioc/types"
	"github.com/vlorc/gioc/utils"
	"reflect"
)

func (si SimpleInvoker) Apply(args ...interface{}) ([]reflect.Value, error) {
	return si.ApplyWith(nil, args...)
}

func (si SimpleInvoker) ApplyWith(_ types.Provider, args ...interface{}) (result []reflect.Value, err error) {
	defer utils.Recover(&err)

	val := reflect.Value(si)
	param := make([]reflect.Value, val.Type().NumIn())
	for i := range param {
		param[i] = reflect.ValueOf(args[i])
	}
	result =  val.Call(param)

	return
}
