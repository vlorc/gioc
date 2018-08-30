// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package text

import (
	"fmt"
	"github.com/vlorc/gioc/types"
	"github.com/vlorc/gioc/utils"
)

func NewTextParser(handle map[string][]types.IdentHandle, paramFactory types.ParamFactory) types.TextParser {
	return &CoreTextParser{
		handle:       handle,
		paramFactory: paramFactory,
	}
}

func (tp *CoreTextParser) Resolve(ctx *types.ParseContext) (err error) {
	defer utils.Recover(&err)
	defer ctx.Scan.End()

	ctx.Scan.Begin()
	for {
		tp.nextToken(ctx)
	}
}

func (tp *CoreTextParser) pop(ctx *types.ParseContext, handle ...types.IdentHandle) (types.Token, string) {
	token, offset, length := ctx.Scan.Scan()
	if types.TOKEN_EOF == token {
		tp.invoke(ctx, handle)
		panic(nil)
	}
	return token, ctx.Dump(offset, length)
}

func (tp *CoreTextParser) nextToken(ctx *types.ParseContext) {
	token, key := tp.pop(ctx)
	tp.dispatch(ctx, token, key)
}

func (tp *CoreTextParser) getParam(ctx *types.ParseContext) {
	for tp.nextParam(ctx) {

	}
}

func (tp *CoreTextParser) nextParam(ctx *types.ParseContext) bool {
	token, key := tp.pop(ctx)
	if types.TOKEN_EOF == token {
		panic(fmt.Errorf("can't get param '%s'", key))
	}
	if types.TOKEN_RPAREN == token {
		return false
	}
	param, err := tp.paramFactory.Instance(token, key)
	if nil != err {
		panic(err)
	}
	ctx.Params = append(ctx.Params, param)
	return true
}

func (tp *CoreTextParser) dispatch(ctx *types.ParseContext, token types.Token, key string) {
	switch token {
	case types.TOKEN_IDENT:
		tp.Invoke(ctx, key)
	case types.TOKEN_CHART, types.TOKEN_STRING:
		ctx.Descriptor.SetName(key[1 : len(key)-1])
	}
}

func (tp *CoreTextParser) invoke(ctx *types.ParseContext, handle []types.IdentHandle) {
	for _, f := range handle {
		if err := f(ctx); nil != err {
			panic(err)
		}
	}
}

func (tp *CoreTextParser) Invoke(ctx *types.ParseContext, key string) {
	handle, ok := tp.handle[key]
	if !ok {
		panic(fmt.Errorf("can't find token '%s'", key))
	}

	ctx.Params = nil
	token, key := tp.pop(ctx, handle...)
	if types.TOKEN_LPAREN != token {
		defer tp.dispatch(ctx, token, key)
	} else {
		tp.getParam(ctx)
	}
	tp.invoke(ctx, handle)
}
