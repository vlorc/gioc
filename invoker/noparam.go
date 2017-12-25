// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package invoker

import (
	"github.com/vlorc/gioc/types"
	"reflect"
)

func (npi NoParamInvoker) Apply( ...interface{}) []reflect.Value {
	return npi.ApplyWith(nil)
}

func (npi NoParamInvoker) ApplyWith(types.Provider,...interface{}) []reflect.Value {
	return reflect.Value(npi).Call(nil)
}
