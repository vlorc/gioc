// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package operation

import (
	"github.com/vlorc/gioc/factory"
	"github.com/vlorc/gioc/types"
	"github.com/vlorc/gioc/utils"
	"reflect"
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
		ctx.Factory = factory.NewFuncFactory(val)
	}
}

func Method(val interface{}) DeclareHandle {
	return func(ctx *DeclareContext) {
		typ := reflect.ValueOf(val).Type()
		if reflect.Func != typ.Kind() || typ.NumOut() <= 0 {
			return
		}
		if typ.NumIn() <= 0 {
			ctx.Factory, ctx.Type, _ = factory.NewMethodFactory(val, nil)
			return
		}
		if Dependency(val)(ctx); nil != ctx.Depend {
			newMethodFactory(ctx)
		}
	}
}

func Instance(val interface{}) DeclareHandle {
	return func(ctx *DeclareContext) {
		ctx.Factory = factory.NewValueFactory(val)
		ctx.Type = val
	}
}

func Pointer(val interface{}) DeclareHandle {
	return func(ctx *DeclareContext) {
		tmp := reflect.ValueOf(val)
		ctx.Factory = factory.NewPointerFactory(tmp)
		ctx.Type = tmp.Elem().Type()
	}
}

func toFactory(ctx *DeclareContext) {
	if nil != ctx.Factory {
		return
	}
	if nil != ctx.Depend {
		toBuildFactory(ctx)
	} else if nil != ctx.Type {
		ctx.Factory = factory.NewTypeFactory(ctx.Type)
	}
}

func getBuild(ctx *DeclareContext, bean types.BeanFactory) types.Builder {
	var buildFactory types.BuilderFactory
	ctx.Context.Container().AsProvider().Assign(&buildFactory)
	build, _ := buildFactory.Instance(bean, ctx.Depend)
	return build
}

func toBuildFactory(ctx *DeclareContext) {
	typ := utils.TypeOf(ctx.Value)
	if reflect.Func == typ.Kind() {
		newMethodFactory(ctx)
		return
	}
	ctx.Factory = getBuild(ctx, factory.NewTypeFactory(ctx.Value)).AsFactory()
}

func newMethodFactory(ctx *DeclareContext) {
	param := getBuild(ctx, factory.NewParamFactory(ctx.Depend.Length()))
	bean, typ, err := factory.NewMethodFactory(ctx.Value, param.AsFactory())
	if nil != err {
		panic(err)
	}
	ctx.Type = typ
	ctx.Factory = bean
}

func toExport(ctx *DeclareContext) {
	ctx.Context.Parent().AsRegister().RegisterFactory(
		factory.NewExportFactory(ctx.Factory, lazyProvider(ctx.Context.Container)),
		ctx.Type,
		ctx.Name,
	)
}
