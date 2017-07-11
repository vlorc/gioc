// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package invoker

import (
	"reflect"
	"github.com/vlorc/gioc/types"
	"github.com/vlorc/gioc/utils"
	"github.com/vlorc/gioc/builder"
)

func NewInvoker(method interface{},builder types.Builder) types.Invoker {
	srcVal := utils.ValueOf(method)
	if reflect.Func != srcVal.Kind(){
		panic(types.NewError(types.ErrTypeNotFunction,method))
	}
	return &CoreInvoker{
		method: srcVal,
		builder: builder,
	}
}

func(i *CoreInvoker)Apply(args ...interface{}) []reflect.Value{
	return i.ApplyWith(nil,args...)
}

func(i *CoreInvoker)ApplyWith(provider types.Provider,args ...interface{}) []reflect.Value{
	temp,err := i.builder.Build(provider,func(ctx *types.BuildContext){
		ctx.FullBefore = builder.MakeIndexFullBefore(args)
	})
	if nil != err {
		return nil
	}
	param := temp.([]reflect.Value)
	return i.method.Call(param)
}