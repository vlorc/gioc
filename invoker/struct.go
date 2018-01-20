// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package invoker

import (
	"github.com/vlorc/gioc/types"
	"reflect"
)

type CoreInvoker struct {
	method  reflect.Value
	param   []reflect.Value
	builder func(types.Provider) types.Builder
}

type NoParamInvoker reflect.Value
type SimpleInvoker reflect.Value

type CoreInvokerFactory struct {}

type WithInvoker struct {
	provider func() types.Provider
	invoker types.Invoker
}