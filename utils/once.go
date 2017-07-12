// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package utils

import (
	"reflect"
	"sync"
)

func Once(src interface{}, init interface{}) {
	srcVal := DirectlyValue(ValueOf(src))
	initVal := ValueOf(init)
	once := sync.Once{}

	dstVal := reflect.MakeFunc(srcVal.Type(), func(args []reflect.Value) []reflect.Value {
		once.Do(func() {
			srcVal.Set(reflect.ValueOf(initVal.Call(args)[0].Interface()))
		})
		return srcVal.Call(args)
	})
	srcVal.Set(dstVal)
}
