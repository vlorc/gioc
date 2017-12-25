// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package operation

import (
	"reflect"
	"github.com/vlorc/gioc/factory"
	"github.com/vlorc/gioc/types"
	"github.com/vlorc/gioc/utils"
	"github.com/vlorc/gioc/builder"
)

func Singleton() DeclareHandle {
	return func(ctx *DeclareContext) {
		ctx.Factory = factory.NewSingleFactory(ctx.Factory)
	}
}

func Mutex() DeclareHandle {
	return func(ctx *DeclareContext) {
		ctx.Factory = factory.NewMutexFactory(ctx.Factory)
	}
}

func Func(val func(types.Provider) (interface{}, error)) DeclareHandle {
	return func(ctx *DeclareContext) {
		ctx.Reset()
		ctx.Factory = factory.NewFuncFactory(val)
	}
}

func Method(val interface{},index ...int) DeclareHandle {
	return func(ctx *DeclareContext) {
		ctx.Reset()
		typ := reflect.ValueOf(val).Type()
		if reflect.Func != typ.Kind() || typ.NumOut() <= 0 {
			return
		}
		if typ.NumIn() <= 0 {
			ctx.Factory, ctx.Type, _ = factory.NewMethodFactory(val, nil,index...)
		}else if toDependency(ctx,val) {
			toMethodFactory(ctx,val,index...)
		}
	}
}

func Interface(val interface{},typ ...interface{}) DeclareHandle {
	return func(ctx *DeclareContext) {
		ctx.Reset()
		if len(typ) <= 0 {
			toInterface(ctx,val)
			return
		}
		ctx.Factory = factory.NewTypeFactory(val)
		ctx.Type = utils.TypeOf(typ[0])
	}
}

func Instance(val interface{}) DeclareHandle {
	return func(ctx *DeclareContext) {
		ctx.Reset()
		ctx.Factory = factory.NewValueFactory(val)
		ctx.Type = reflect.TypeOf(val)
	}
}

func Pointer(val interface{}) DeclareHandle {
	return func(ctx *DeclareContext) {
		ctx.Reset()
		if tmp := reflect.ValueOf(val); reflect.Ptr == tmp.Kind() {
			ctx.Factory = factory.NewPointerFactory(tmp)
			ctx.Type = tmp.Elem().Type()
		}
	}
}

func toMethodFactory(ctx *DeclareContext,val interface{},index ...int) {
	param := builder.NewBuilder(factory.NewParamFactory(ctx.Depend.Length()), ctx.Depend)
	bean, typ, err := factory.NewMethodFactory(val, param.AsFactory(),index...)
	if nil != err {
		panic(err)
	}
	ctx.Type = typ
	ctx.Factory = bean
}

func toRegistered(ctx *DeclareContext) {
	ctx.Context.Container().AsRegister().RegisterFactory(
		ctx.Factory,
		ctx.Type,
		ctx.Name)
}

func toExport(ctx *DeclareContext) {
	var bean types.BeanFactory
	if nil != ctx.Depend {
		bean = factory.NewExportFactory(ctx.Factory, lazyProvider(ctx.Context.Container))
	} else {
		bean = ctx.Factory
	}
	ctx.Context.Parent().AsRegister().RegisterFactory(
		bean,
		ctx.Type,
		ctx.Name,
	)
}

func toDependency(ctx *DeclareContext,val interface{}) (ok bool) {
	var dependFactory types.DependencyFactory
	ctx.Context.Parent().AsProvider().Assign(&dependFactory)
	dep, err := dependFactory.Instance(val)
	if nil == err {
		ok = true
		ctx.Depend = dep
	}
	return
}

func toInterface(ctx *DeclareContext,val interface{}) bool{
	src := utils.DirectlyValue(utils.ValueOf(val))
	ok := reflect.Interface == src.Kind() && !src.IsNil()
	if ok{
		ctx.Factory = factory.NewValueFactory(src.Interface())
		ctx.Type = src.Type()
	}
	return ok
}