// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package text

import (
	"testing"
	"strings"
	"github.com/vlorc/gioc/types"
)

func Test_TokenScan(t *testing.T) {
	line := `id(-10,"test\n") 'age' default -0.1 nil`
	scan := NewTokenScan()
	arr := []struct {
		token types.Token
		str   string
	}{
		{types.TOKEN_IDENT, "id"},
		{types.TOKEN_LPAREN, "("},
		{types.TOKEN_NUMBER, "-10"},
		{types.TOKEN_COMMA, ","},
		{types.TOKEN_ESTRING, `"test\n"`},
		{types.TOKEN_RPAREN, ")"},
		{types.TOKEN_CHART, `'age'`},
		{types.TOKEN_IDENT, "default"},
		{types.TOKEN_FLOAT, "-0.1"},
		{types.TOKEN_NULL, "nil"},
	}

	scan.SetInput(strings.NewReader(line))

	scan.Begin()
	for i := 0; i < len(arr); i++ {
		tk, o, l := scan.Scan()
		if tk != arr[i].token {
			t.Errorf("can't matching token %d,must is %d", tk, arr[i].token)
		}
		if line[o:l] != arr[i].str {
			t.Errorf("can't matching str %s,must is %s", line[o:l], arr[i].str)
		}
	}
}
