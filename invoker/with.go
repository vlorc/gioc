// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package invoker

import (
	"github.com/vlorc/gioc/types"
	"reflect"
)

func (wi *WithInvoker) Apply(args ...interface{}) ([]reflect.Value, error) {
	return wi.invoker.ApplyWith(wi.provider(), args...)
}

func (wi *WithInvoker) ApplyWith(_ types.Provider, args ...interface{}) ([]reflect.Value, error) {
	return wi.Apply(args...)
}
