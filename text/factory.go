// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package text

import (
	"github.com/vlorc/gioc/types"
)

func NewParamFactory() types.ParamFactory {
	return &CoreParamFactory{table: map[types.Token]func(string) types.Param{
		types.TOKEN_CHART:   NewParamString,
		types.TOKEN_STRING:  NewParamString,
		types.TOKEN_ECHART:  NewParamEscapeString,
		types.TOKEN_ESTRING: NewParamEscapeString,
		types.TOKEN_NUMBER:  NewParamNumber,
		types.TOKEN_FLOAT:   NewParamFloat,
		types.TOKEN_NULL:    NewParamNull,
	}}
}

func NewTextParserFactory() types.TextParserFactory {
	return &CoreTextParserFactory{}
}

func NewTokenScanFactory() types.TokenScanFactory {
	return &CoreTokenScanFactory{}
}

func (pf *CoreParamFactory) Instance(token types.Token, value string) (types.Param, error) {
	param, ok := pf.table[token]
	if ok {
		return param(value), nil
	}
	return nil, types.NewError(types.ErrTokenNotSupport, token)
}

func (tpf *CoreTextParserFactory) Instance(handle map[string][]types.IdentHandle, paramFactory types.ParamFactory) (types.TextParser, error) {
	return NewTextParser(handle, paramFactory), nil
}

func (tsf *CoreTokenScanFactory) Instance() (types.TokenScan, error) {
	return NewTokenScan(), nil
}
