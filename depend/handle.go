// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package depend

import (
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
		extendsHandle,
		flagsHandle(types.DEPENDENCY_FLAG_EXTENDS),
	},
	"default": {
		defaultHandle,
		flagsHandle(types.DEPENDENCY_FLAG_DEFAULT),
	},
	"lazy": {
		lazyHandle,
		flagsHandle(types.DEPENDENCY_FLAG_LAZY),
	},
	"id":    {nameHandle},
	"name":  {nameHandle},
	"lower": {lowerCaseHandle},
	"upper": {upperCaseHandle},
	"make": {
		makeHandle,
		flagsHandle(types.DEPENDENCY_FLAG_DEFAULT | types.DEPENDENCY_FLAG_SKIP),
	},
	"new": {
		newHandle,
		flagsHandle(types.DEPENDENCY_FLAG_DEFAULT | types.DEPENDENCY_FLAG_SKIP),
	},
	"request": {
		requestHandle,
		flagsHandle(types.DEPENDENCY_FLAG_DEFAULT | types.DEPENDENCY_FLAG_REQUEST),
	},
	"skip": {
		skipHandle,
		flagsHandle(types.DEPENDENCY_FLAG_SKIP),
	},
	"*": {
		anyHandle,
		flagsHandle(types.DEPENDENCY_FLAG_ANY),
	},
}

func lowerCaseHandle(ctx *types.ParseContext) error {
	ctx.Descriptor.SetName(strings.ToLower(ctx.Descriptor.Name()))
	return nil
}

func upperCaseHandle(ctx *types.ParseContext) error {
	ctx.Descriptor.SetName(strings.ToUpper(ctx.Descriptor.Name()))
	return nil
}

func nameHandle(ctx *types.ParseContext) error {
	ctx.Descriptor.SetName(ctx.Params[0].String())
	return nil
}

func getDefault(ctx *types.ParseContext) func() reflect.Value {
	var val reflect.Value
	if len(ctx.Params) > 0 {
		val = ctx.Params[0].Value()
	}
	val = utils.Convert(val, ctx.Descriptor.Type())
	return func() reflect.Value {
		return val
	}
}

func defaultHandle(ctx *types.ParseContext) error {
	ctx.Descriptor.SetDefault(getDefault(ctx))
	return nil
}

func lazyHandle(ctx *types.ParseContext) (err error) {
	if typ := ctx.Descriptor.Type(); reflect.Func != typ.Kind() || typ.NumOut() != 1 || typ.NumIn() > 0 {
		err = types.NewWithError(types.ErrTypeNotSupport, typ, ctx.Descriptor.Name())
	}
	return
}

func extendsHandle(ctx *types.ParseContext) error {
	typ := ctx.Descriptor.Type()
	if 0 != ctx.Descriptor.Flags()&types.DEPENDENCY_FLAG_LAZY {
		typ = utils.DirectlyType(ctx.Descriptor.Type()).Out(0)
	}
	dep, err := ctx.Factory.Instance(typ)
	ctx.Descriptor.SetDepend(dep)
	return err
}

func flagsHandle(flag types.DependencyFlag) types.IdentHandle {
	return func(ctx *types.ParseContext) error {
		ctx.Descriptor.SetFlags(ctx.Descriptor.Flags() | flag)
		return nil
	}
}

func newHandle(ctx *types.ParseContext) (err error) {
	if reflect.Ptr != ctx.Descriptor.Type().Kind() {
		err = types.NewWithError(types.ErrTypeNotSupport, ctx.Descriptor.Type(), ctx.Descriptor.Name())
	} else {
		ctx.Descriptor.SetDefault(newDefault(ctx.Descriptor.Type().Elem()))
	}
	return
}

func makeHandle(ctx *types.ParseContext) (err error) {
	ln, cn := 0, 0
	if len(ctx.Params) > 0 {
		if ln = int(ctx.Params[0].Number()); len(ctx.Params) > 1 {
			cn = int(ctx.Params[1].Number())
		} else {
			cn = ln
		}
	}
	switch ctx.Descriptor.Type().Kind() {
	case reflect.Chan:
		ctx.Descriptor.SetDefault(makeChanDefault(ctx.Descriptor.Type(), cn))
	case reflect.Map:
		ctx.Descriptor.SetDefault(makeMapDefault(ctx.Descriptor.Type(), cn))
	case reflect.Slice:
		ctx.Descriptor.SetDefault(makeSliceDefault(ctx.Descriptor.Type(), ln, cn))
	default:
		err = types.NewWithError(types.ErrTypeNotSupport, ctx.Descriptor.Type(), ctx.Descriptor.Name())
	}
	return
}

func makeSliceDefault(typ reflect.Type, ln, cn int) func() reflect.Value {
	return func() reflect.Value {
		return reflect.MakeSlice(typ, ln, cn)
	}
}

func makeMapDefault(typ reflect.Type, ln int) func() reflect.Value {
	if ln > 0 {
		return func() reflect.Value {
			return reflect.MakeMapWithSize(typ, ln)
		}
	}
	return func() reflect.Value {
		return reflect.MakeMap(typ)
	}
}

func makeChanDefault(typ reflect.Type, ln int) func() reflect.Value {
	return func() reflect.Value {
		return reflect.MakeChan(typ, ln)
	}
}

func newDefault(typ reflect.Type) func() reflect.Value {
	return func() reflect.Value {
		return reflect.New(typ)
	}
}

func anyHandle(ctx *types.ParseContext) error {
	if reflect.Slice != ctx.Descriptor.Type().Kind() ||
		0 == ctx.Descriptor.Flags()&types.DEPENDENCY_FLAG_LAZY ||
		reflect.Slice != ctx.Descriptor.Type().Out(0).Kind() {
		return types.NewWithError(types.ErrTypeNotSupport, ctx.Descriptor.Type(), ctx.Descriptor.Name())
	}
	return nil
}

func skipHandle(ctx *types.ParseContext) error {
	if 0 == ctx.Descriptor.Flags()&types.DEPENDENCY_FLAG_DEFAULT {
		return types.NewWithError(types.ErrTypeNotSupport, ctx.Descriptor.Type(), ctx.Descriptor.Name())
	}
	return nil
}

func requestHandle(ctx *types.ParseContext) error {
	if reflect.Func != ctx.Descriptor.Type().Kind() {
		return types.NewWithError(types.ErrTypeNotSupport, ctx.Descriptor.Type(), ctx.Descriptor.Name())
	}
	return nil
}
