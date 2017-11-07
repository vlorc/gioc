// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package tag

const (
	TOKEN_EOF     Token = -1
	TOKEN_ERROR         = 2
	TOKEN_CHART         = 4
	TOKEN_ECHART        = 7
	TOKEN_STRING        = 9
	TOKEN_ESTRING       = 12
	TOKEN_IDENT         = 13
	TOKEN_LPAREN        = 14
	TOKEN_RPAREN        = 15
	TOKEN_COMMA         = 16
	TOKEN_NUMBER        = 19
	TOKEN_FLOAT         = 20
	TOKEN_NULL           = 23
)

const (
	Invalid Kind = iota
	Int
	Float
	String
	Boolean
	Null
)
