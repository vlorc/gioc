// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package tag

import (
	"fmt"
	"github.com/vlorc/gioc/types"
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

func (tp *TagParser) pop(ctx *TagContext,handle ...TagHandle) (Token,string) {
	token, offset, length := ctx.TokenScan.Scan()
	if TOKEN_EOF == token {
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
	case TOKEN_RPAREN:
		ok = false
	case TOKEN_CHART, TOKEN_STRING:
		ctx.Params = append(ctx.Params, NewParamString(key))
	case TOKEN_ECHART,TOKEN_ESTRING:
		ctx.Params = append(ctx.Params, NewParamEscapeString(key))
	case TOKEN_NUMBER:
		ctx.Params = append(ctx.Params, NewParamNumber(key))
	case TOKEN_FLOAT:
		ctx.Params = append(ctx.Params, NewParamFloat(key))
	case TOKEN_NULL:
		ctx.Params = append(ctx.Params, NewParamNull(false))
	}
	return
}

func (tp *TagParser) dispatch(ctx *TagContext, token Token, key string) {
	switch token {
	case TOKEN_IDENT:
		tp.Invoke(ctx, key)
	case TOKEN_CHART, TOKEN_STRING:
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
	if TOKEN_LPAREN != token {
		defer tp.dispatch(ctx, token, key)
	} else {
		tp.getParam(ctx)
	}
	tp.invoke(ctx,handle)
}
