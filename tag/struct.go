// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package tag

import (
	"github.com/vlorc/gioc/types"
	"reflect"
	"bufio"
)

type Token int
type Kind uint
type ParamString string
type ParamNumber int64
type ParamFloat float64
type ParamNull bool

type TokenScan struct {
	state    int
	offset   int
	position int
	input    *bufio.Reader
	dump     func(Token, int, int) bool
}


type TagContext struct {
	Factory    types.DependencyFactory
	Descriptor types.Descriptor
	Params     []Param
	Skip       func(string) bool
	Tag        string
	TokenScan  *TokenScan
}

type TagHandle func(*TagContext) error

type TagParser struct {
	tagHandle map[string][]TagHandle
}

type Param interface {
	String() string
	Number() int64
	Float() float64
	Boolean() bool
	Value() reflect.Value
	Kind() Kind
}