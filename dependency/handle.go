// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package dependency

import (
	"github.com/vlorc/gioc/factory"
	"github.com/vlorc/gioc/types"
	"github.com/vlorc/gioc/utils"
	"reflect"
	"strings"
)

var DefaultHandle = map[string][]types.IdentHandle{
	"optional": {
		flagsHandle(types.DEPENDENCY_FLAG_OPTIONAL),
	},
	"extends": {
		flagsHandle(types.DEPENDENCY_FLAG_EXTENDS),
		extendsHandle,
	},
	"default": {
		flagsHandle(types.DEPENDENCY_FLAG_DEFAULT),
		defaultHandle,
	},
	"lazy": {
		flagsHandle(types.DEPENDENCY_FLAG_LAZY),
		lazyHandle,
	},
	"id":    {idHandle},
	"name":  {nameHandle},
	"lower": {lowerCaseHandle},
	"upper": {upperCaseHandle},
	"make": {
		flagsHandle(types.DEPENDENCY_FLAG_MAKE),
		makeHandle,
	},
	"request": {
		flagsHandle(types.DEPENDENCY_FLAG_REQUEST),
		requestHandle,
	},
	"*": {
		flagsHandle(types.DEPENDENCY_FLAG_ANY),
		anyHandle,
	},
	"any": {
		flagsHandle(types.DEPENDENCY_FLAG_ANY),
		anyHandle,
	},
}

func __stringFactory(s string) types.StringFactory {
	if strings.HasPrefix(s, "${") && strings.HasSuffix(s, "}") {
		return types.NewNameFactory(s[2 : len(s)-1])
	}
	return types.NewStringFactory(s)
}

func lowerCaseHandle(ctx *types.ParseContext) error {
	ctx.Dependency.Name = append(ctx.Dependency.Name, types.RawStringFactory(strings.ToLower(ctx.Dependency.Origin.Name)))
	return nil
}

func upperCaseHandle(ctx *types.ParseContext) error {
	ctx.Dependency.Name = append(ctx.Dependency.Name, types.RawStringFactory(strings.ToUpper(ctx.Dependency.Origin.Name)))
	return nil
}

func makeHandle(ctx *types.ParseContext) error {
	typ := ctx.Dependency.Type
	if len(ctx.Params) > 0 {
		length := int(ctx.Params[0].Number())
		ctx.Dependency.Default = func(c *types.DependencyContext) types.BeanFactory {
			return factory.NewMakeFactory(typ, length)
		}
	} else {
		ctx.Dependency.Default = func(c *types.DependencyContext) types.BeanFactory {
			return factory.NewMakeFactory(typ)
		}
	}
	return nil
}

func idHandle(ctx *types.ParseContext) error {
	if len(ctx.Params) > 0 {
		ctx.Dependency.Name = []types.StringFactory{__stringFactory(ctx.Params[0].String())}
	}
	return nil
}

func nameHandle(ctx *types.ParseContext) error {
	for i := range ctx.Params {
		ctx.Dependency.Name = append(ctx.Dependency.Name, __stringFactory(ctx.Params[i].String()))
	}
	return nil
}

func getDefault(ctx *types.ParseContext) reflect.Value {
	var val reflect.Value
	if len(ctx.Params) > 0 {
		val = ctx.Params[0].Value()
	}
	val = utils.Convert(val, ctx.Dependency.Type)
	return val
}

func defaultHandle(ctx *types.ParseContext) error {
	instance := getDefault(ctx).Interface()

	ctx.Dependency.Wrapper.Append(65535, defaultWrapper(instance))
	return nil
}

func lazyHandle(ctx *types.ParseContext) error {
	typ := ctx.Dependency.Type
	if reflect.Func != typ.Kind() || typ.NumOut() != 1 || typ.NumIn() > 0 {
		return types.NewWithError(types.ErrTypeNotSupport, typ, ctx.Dependency.Origin.Name)
	}

	ctx.Dependency.Type = ctx.Dependency.Type.Out(0)
	ctx.Dependency.Wrapper.Append(0, lazyWrapper(typ))
	return nil
}

func extendsHandle(ctx *types.ParseContext) error {
	typ := utils.IndirectType(ctx.Dependency.Type)

	ctx.Dependency.Default = func(*types.DependencyContext) types.BeanFactory {
		return nil
	}

	if reflect.Slice == typ.Kind() {
		ctx.Dependency.Wrapper.Append(256, extendSliceWrapper(typ.Elem(), ctx.Dependency.Name...))
		return nil
	}
	if reflect.Struct == typ.Kind() {
		dep, err := ctx.Factory.Instance(typ)
		if nil != err {
			return err
		}
		ctx.Dependency.Wrapper.Append(256, extendStructWrapper(dep, ctx.Dependency.Type))
		return nil
	}

	return types.NewError(types.ErrExtendNotSupport, typ)
}

func flagsHandle(flag types.DependencyFlag) types.IdentHandle {
	return func(ctx *types.ParseContext) error {
		ctx.Dependency.Flags |= flag
		return nil
	}
}

func anyHandle(ctx *types.ParseContext) error {
	ctx.Dependency.Name = []types.StringFactory{}
	return nil
}

func requestHandle(ctx *types.ParseContext) error {
	typ := ctx.Dependency.Type
	if reflect.Func != typ.Kind() || typ.NumOut() != 1 || typ.NumIn() > 0 {
		return types.NewWithError(types.ErrTypeNotSupport, typ, ctx.Dependency.Origin.Name)
	}

	ctx.Dependency.Type = ctx.Dependency.Type.Out(0)
	ctx.Dependency.Wrapper.Append(0, requestWrapper(typ))
	return nil
}
