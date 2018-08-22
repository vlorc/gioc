// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package utils

import (
	"reflect"
	"sync"
)

func LazyProxy(init func([]reflect.Value) []reflect.Value) func([]reflect.Value) []reflect.Value {
	var proxy func([]reflect.Value) []reflect.Value
	var once sync.Once
	proxy = func(args []reflect.Value) []reflect.Value {
		once.Do(func() {
			defer func() {
				if err := recover(); nil != err {
					proxy = func(args []reflect.Value) []reflect.Value {
						panic(err)
					}
				}
			}()
			proxy = ApplyProxy(init(args))
		})
		return proxy(args)
	}
	return func(args []reflect.Value) []reflect.Value {
		return proxy(args)
	}
}

func ApplyProxy(fn interface{}) func([]reflect.Value) []reflect.Value {
	switch r := fn.(type) {
	case []reflect.Value:
		return func([]reflect.Value) []reflect.Value {
			return r
		}
	case func([]reflect.Value) []reflect.Value:
		return r
	case func() []reflect.Value:
		return func([]reflect.Value) []reflect.Value {
			return r()
		}
	case func():
		return func([]reflect.Value) []reflect.Value {
			r()
			return nil
		}
	default:
		return func(args []reflect.Value) []reflect.Value {
			return reflect.ValueOf(fn).Call(args)
		}
	}
}

func Lazy(fn interface{}) interface{} {
	return reflect.MakeFunc(
		reflect.TypeOf(fn),
		LazyProxy(ApplyProxy(fn)),
	).Interface()
}
