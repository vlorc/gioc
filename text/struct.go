// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package text

import (
	"github.com/vlorc/gioc/types"
	"bufio"
)

type ParamString string
type ParamNumber int64
type ParamFloat float64
type ParamNull bool

type CoreTokenScan struct {
	state    int
	offset   int
	position int
	input    *bufio.Reader
	dump     func(types.Token, int, int) bool
}

type CoreTextParser struct {
	handle map[string][]types.IdentHandle
	paramFactory types.ParamFactory
}

type CoreParamFactory struct {
	table map[types.Token]func(string)types.Param
}
type CoreTokenScanFactory struct {}
type CoreTextParserFactory struct {}