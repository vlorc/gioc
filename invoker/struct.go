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
	builder types.Builder
}

type CoreInvokerFactory struct {
}
