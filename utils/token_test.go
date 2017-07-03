// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package utils

import (
	"testing"
)

func Test_TokenScan(t *testing.T) {
	line := `id(-10,"test\n") 'age' default    `
	scan := NewTokenScan()
	arr := []struct{
		token Token
		str string
	}{
		{TOKEN_IDENT,"id"},
		{TOKEN_LPAREN,"("},
		{TOKEN_NUMBER,"-10"},
		{TOKEN_COMMA,","},
		{TOKEN_ESTRING,`"test\n"`},
		{TOKEN_RPAREN,")"},
		{TOKEN_CHART,`'age'`},
		{TOKEN_IDENT,"default"},
	}

	scan.Init(line)

	scan.Begin()
	for i := 0;i < len(arr); i ++{
		tk,o,l := scan.Scan()
		if tk != arr[i].token{
			t.Errorf("can't matching token %d,must is %d", tk,arr[i].token)
		}
		if line[o:l] != arr[i].str {
			t.Errorf("can't matching str %s,must is %s", line[o:l],arr[i].str)
		}
	}
}
