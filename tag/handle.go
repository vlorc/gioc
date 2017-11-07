// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package tag

import (
	"github.com/vlorc/gioc/types"
	"github.com/vlorc/gioc/utils"
	"reflect"
	"strings"
)

func lowerCaseHandle(ctx *TagContext) error {
	ctx.Descriptor.SetName(strings.ToLower(ctx.Descriptor.Name()))
	return nil
}

func upperCaseHandle(ctx *TagContext) error {
	ctx.Descriptor.SetName(strings.ToUpper(ctx.Descriptor.Name()))
	return nil
}

func nameHandle(ctx *TagContext) error {
	ctx.Descriptor.SetName(ctx.Params[0].String())
	return nil
}

func defaultHandle(ctx *TagContext) error {
	var val reflect.Value
	if len(ctx.Params) > 0 {
		val = ctx.Params[0].Value()
	}

	val = utils.Convert(val,ctx.Descriptor.Type())
	ctx.Descriptor.SetDefault(val)

	return nil
}

func lazyHandle(ctx *TagContext) (err error) {
	typ := utils.DirectlyType(ctx.Descriptor.Type())
	if reflect.Func != typ.Kind() || typ.NumOut() != 1 || typ.NumIn() > 0 {
		err = types.NewError(types.ErrTypeNotSupport, typ, ctx.Descriptor.Name())
	}
	return
}

func extendsHandle(ctx *TagContext) error {
	typ := ctx.Descriptor.Type()
	if 0 != ctx.Descriptor.Flags()&types.DEPENDENCY_FLAG_LAZY {
		typ = utils.DirectlyType(ctx.Descriptor.Type()).Out(0)
	}
	dep, err := ctx.Factory.Instance(typ)
	ctx.Descriptor.SetDepend(dep)
	return err
}

func flagsHandle(flag types.DependencyFlag) TagHandle {
	return func(ctx *TagContext) error {
		ctx.Descriptor.SetFlags(ctx.Descriptor.Flags() | flag)
		return nil
	}
}
