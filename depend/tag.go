// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package depend

import (
	"fmt"
	"github.com/vlorc/gioc/types"
	"github.com/vlorc/gioc/utils"
	"reflect"
	"strings"
)

func NewTagParser() *TagParser {
	obj := &TagParser{}
	obj.tagHandle = map[string][]TagHandle{
		"optional": []TagHandle{
			flagsHandle(types.DEPENDENCY_FLAG_OPTIONAL),
		},
		"extends": []TagHandle{
			flagsHandle(types.DEPENDENCY_FLAG_EXTENDS),
			extendsHandle,
		},
		"default": []TagHandle{
			flagsHandle(types.DEPENDENCY_FLAG_DEFAULT),
			defaultHandle,
		},
		"lazy": []TagHandle{
			flagsHandle(types.DEPENDENCY_FLAG_LAZY),
			lazyHandle,
		},
		"id":    []TagHandle{nameHandle},
		"name":  []TagHandle{nameHandle},
		"lower": []TagHandle{lowerCaseHandle},
		"upper": []TagHandle{upperCaseHandle},
	}
	return obj
}

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
	val := reflect.Zero(ctx.Descriptor.Type())
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

func (tp *TagParser) Resolve(ctx *TagContext) {
	defer func() {
		if r := recover(); nil != r {
			panic(r)
		}
	}()

	ctx.TokenScan.Init(strings.NewReader(ctx.Tag))
	ctx.TokenScan.Begin()
	for {
		tp.nextToken(ctx)
	}
}

func (tp *TagParser) pop(ctx *TagContext,handle ...TagHandle) (utils.Token,string) {
	token, offset, length := ctx.TokenScan.Scan()
	if utils.TOKEN_EOF == token {
		tp.invoke(ctx,handle)
		panic(nil)
	}
	return token,ctx.Tag[offset:length]
}

func (tp *TagParser) nextToken(ctx *TagContext) {
	token, key := tp.pop(ctx)
	tp.dispatch(ctx, token, key)
}

func (tp *TagParser) getParam(ctx *TagContext) {
	for tp.nextParam(ctx) {

	}
}

func (tp *TagParser) nextParam(ctx *TagContext) (ok bool) {
	ok = true
	token, key := tp.pop(ctx)
	switch token {
	case utils.TOKEN_RPAREN:
		ok = false
	case utils.TOKEN_CHART, utils.TOKEN_STRING:
		ctx.Params = append(ctx.Params, NewParamString(key))
	case utils.TOKEN_NUMBER:
		ctx.Params = append(ctx.Params, NewParamNumber(key))
	}
	return
}

func (tp *TagParser) dispatch(ctx *TagContext, token utils.Token, key string) {
	switch token {
	case utils.TOKEN_IDENT:
		tp.Invoke(ctx, key)
	case utils.TOKEN_CHART, utils.TOKEN_STRING:
		ctx.Descriptor.SetName(key[1 : len(key)-1])
	}
}

func (tp *TagParser) invoke(ctx *TagContext, handle []TagHandle) {
	for _, f := range handle {
		if err := f(ctx); nil != err {
			panic(err)
		}
	}
}

func (tp *TagParser) Invoke(ctx *TagContext, key string) {
	handle, ok := tp.tagHandle[key]
	if !ok {
		panic(fmt.Errorf("can't find token '%s'", key))
	}

	ctx.Params = nil
	token, key := tp.pop(ctx,handle...)
	if utils.TOKEN_LPAREN != token {
		defer tp.dispatch(ctx, token, key)
	} else {
		tp.getParam(ctx)
	}
	tp.invoke(ctx,handle)
}
