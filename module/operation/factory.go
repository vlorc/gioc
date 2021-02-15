// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package operation

import (
	"github.com/vlorc/gioc/factory"
	"github.com/vlorc/gioc/types"
	"github.com/vlorc/gioc/utils"
	"os"
	"reflect"
)

func Singleton() DeclareHandle {
	return func(ctx *DeclareContext) {
		if nil != ctx.Factory {
			ctx.Factory = factory.NewSingleFactory(ctx.Factory)
		}
	}
}

func Mutex() DeclareHandle {
	return func(ctx *DeclareContext) {
		if nil != ctx.Factory {
			ctx.Factory = factory.NewMutexFactory(ctx.Factory)
		}
	}
}

func Convert(val interface{}) DeclareHandle {
	return func(ctx *DeclareContext) {
		if nil != ctx.Factory {
			typ := utils.TypeOf(val)
			ctx.Factory = factory.NewConvertFactory(ctx.Factory, typ)
			ctx.Type = typ
		}
	}
}

func Func(val func(types.Provider) (interface{}, error)) DeclareHandle {
	return func(ctx *DeclareContext) {
		ctx.Reset()
		ctx.Factory = factory.NewFuncFactory(val)
	}
}

func Method(val interface{}, index ...int) DeclareHandle {
	return Method1(val, index...)
}

func Method1(val interface{}, index ...int) DeclareHandle {
	return func(ctx *DeclareContext) {
		ctx.Reset()

		typ := reflect.ValueOf(val).Type()
		if reflect.Func != typ.Kind() || typ.NumOut() <= 0 {
			return
		}
		if typ.NumIn() <= 0 {
			ctx.Factory, ctx.Type, _ = factory.NewMethodFactory(val, nil, index...)
		} else if toDependency(ctx, val) {
			toMethodFactory(ctx, val, index...)
		}
	}
}

func Method2(val interface{}, index ...int) DeclareHandle {
	return func(ctx *DeclareContext) {
		Method1(val)(ctx)

		instance, err := ctx.Factory.Instance(ctx.Context.Container().AsProvider())
		if nil != err {
			ctx.Factory = factory.NewValueFactory(nil, err)
			return
		}
		if v := reflect.ValueOf(instance); !v.IsValid() || (v.IsNil() && (reflect.Func == v.Kind() || reflect.Ptr == v.Kind() || reflect.Interface == v.Kind())) {
			ctx.Factory, ctx.Type = nil, nil
			return
		}

		if typ := utils.TypeOf(ctx.Type); reflect.Func == typ.Kind() {
			ctx.Factory, ctx.Type = nil, nil
			Method1(instance, index...)(ctx)
		} else {
			ctx.Factory = factory.NewValueFactory(instance)
		}
	}
}

func Instance(val interface{}) DeclareHandle {
	return func(ctx *DeclareContext) {
		ctx.Reset()
		ctx.Factory = factory.NewValueFactory(val)
		ctx.Type = reflect.TypeOf(val)
	}
}

func Env(keys ...string) DeclareHandle {
	if len(keys) <= 0 {
		return func(c *DeclareContext) {

		}
	}

	if len(keys) == 1 {
		k := keys[0]
		return func(ctx *DeclareContext) {
			ctx.Reset()

			ctx.Name = k
			ctx.Factory = factory.NewValueFactory(os.Getenv(k))
			ctx.Type = types.StringType
		}
	}

	return func(ctx *DeclareContext) {
		ctx.Reset()

		for _, k := range keys {
			c := *ctx
			c.Name = k
			ctx.Factory = factory.NewValueFactory(os.Getenv(k))
			ctx.Type = types.StringType
			c.done(&c)
		}
	}
}

func New(val interface{}) DeclareHandle {
	return func(ctx *DeclareContext) {
		ctx.Reset()
		typ := utils.TypeOf(val)
		ctx.Factory = factory.NewTypeFactory(typ)
		ctx.Type = reflect.PtrTo(typ)
	}
}

func Pointer(val interface{}) DeclareHandle {
	return func(ctx *DeclareContext) {
		ctx.Reset()
		if src := reflect.ValueOf(val); reflect.Ptr == src.Kind() && !src.IsNil() {
			ctx.Factory = factory.NewPointerFactory(src)
			ctx.Type = src.Type().Elem()
		}
	}
}

func toMethodFactory(ctx *DeclareContext, val interface{}, index ...int) {
	param := factory.NewDependencyFactory(factory.NewParamFactory(ctx.Dependency.Length()), ctx.Dependency)
	f, typ, err := factory.NewMethodFactory(val, param, index...)
	if nil != err {
		utils.Panic(err)
	}
	ctx.Type = typ
	ctx.Factory = f
}

func toRegistered(ctx *DeclareContext) {
	c := ctx.Context.Container()
	r := c.AsRegister()
	r.Factory(ctx.Factory, ctx.Type, ctx.Name)
}

func toExport(ctx *DeclareContext) {
	var bean types.BeanFactory
	if nil != ctx.Dependency {
		bean = factory.NewExportFactory(ctx.Factory, lazyProvider(ctx.Context.Container))
	} else {
		bean = ctx.Factory
	}
	ctx.Context.Parent().AsRegister().Factory(bean, ctx.Type, ctx.Name)
}

func toPrimary(ctx *DeclareContext) {
	var bean types.BeanFactory
	if nil != ctx.Dependency {
		bean = factory.NewExportFactory(ctx.Factory, lazyProvider(ctx.Context.Container))
	} else {
		bean = ctx.Factory
	}
	ctx.Context.Parent().AsRegister().Selector().Set(utils.TypeOf(ctx.Type), ctx.Name, bean)
}

func toDependency(ctx *DeclareContext, val interface{}) bool {
	var dependFactory types.DependencyFactory
	ctx.Context.Parent().AsProvider().Load(&dependFactory)
	dep, err := dependFactory.Instance(val)
	if nil == err {
		ctx.Dependency = dep
		return true
	}
	// add check error

	return false
}
