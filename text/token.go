// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package text

import (
	"bufio"
	"errors"
	"github.com/vlorc/gioc/types"
	"io"
)

var table = []byte{
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	13, 1, 3, 8, 13, 14, 15, 16, 17, 18, 19, 19, 20, 21, 13, 13,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	3, 3, 4, 3, 5, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6,
	6, 6, 7, 6, 5, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	8, 8, 8, 9, 10, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11,
	11, 11, 11, 12, 10, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	13, 0, 2, 2, 13, 0, 0, 0, 13, 13, 13, 13, 2, 13, 13, 13,
	0, 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0, 0,
	13, 0, 0, 0, 0, 0, 0, 0, 13, 13, 19, 19, 20, 2, 2, 2,
	13, 0, 0, 0, 0, 0, 0, 0, 13, 13, 19, 19, 20, 2, 2, 2,
	2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 19, 19, 20, 2, 2, 2,
	2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 20, 20, 2, 2, 2, 2,
	13, 0, 0, 0, 13, 0, 0, 0, 13, 13, 13, 13, 2, 13, 22, 13,
	13, 0, 0, 0, 13, 0, 0, 0, 13, 13, 13, 13, 2, 13, 13, 23,
	13, 0, 2, 2, 13, 0, 0, 0, 13, 13, 13, 13, 2, 13, 13, 13,
}

var token = []types.Token{
	4:  types.TOKEN_CHART,
	7:  types.TOKEN_ECHART,
	9:  types.TOKEN_STRING,
	12: types.TOKEN_ESTRING,
	13: types.TOKEN_IDENT,
	14: types.TOKEN_LPAREN,
	15: types.TOKEN_RPAREN,
	16: types.TOKEN_COMMA,
	19: types.TOKEN_NUMBER,
	20: types.TOKEN_FLOAT,
	23: types.TOKEN_NULL,
}

var index = [128]byte{
	0x30: 10,
	0x31: 11,
	0x32: 11,
	0x33: 11,
	0x34: 11,
	0x35: 11,
	0x36: 11,
	0x37: 11,
	0x38: 11,
	0x39: 11,
	0x2e: 12,
	0x6e: 13,
	0x69: 14,
	0x6c: 15,
}

var transfer []func(*CoreTokenScan, int, int, int)

func init() {
	for i, v := range ` '"\(),-+` {
		index[v] = byte(i + 1)
	}
	for i := 0; i <= 32; i++ {
		index[i] = 1
	}
	transfer = make([]func(*CoreTokenScan, int, int, int), len(token))
	for i := range transfer {
		transfer[i] = func(d *CoreTokenScan, n int, _ int, l int) {
			d.position += l
			d.state = n
		}
	}

	transfer[2] = func(d *CoreTokenScan, _ int, _ int, _ int) {
		panic(errors.New("illegal char"))
	}

	transfer[1] = func(d *CoreTokenScan, n int, _ int, l int) {
		d.position += l
		d.offset += l
		d.state = n
	}

	transfer[0] = func(d *CoreTokenScan, n int, i int, l int) {
		ok := d.reset()
		d.Transfer(i, l)
		if ok {
			panic(nil)
		}
	}
}

func NewTokenScan() types.TokenScan {
	return &CoreTokenScan{}
}

func (ts *CoreTokenScan) Offset() int {
	return ts.offset
}

func (ts *CoreTokenScan) Position() int {
	return ts.position
}

func (ts *CoreTokenScan) Begin() {
	ts.position = 0
	ts.Reset()
}

func (ts *CoreTokenScan) SetInput(stream io.Reader) types.TokenScan {
	ts.input = bufio.NewReader(stream)
	return ts
}
func (ts *CoreTokenScan) Reset() {
	ts.state = 1
	ts.offset = ts.position
}

func (ts *CoreTokenScan) Next() bool {
	c, l, err := ts.input.ReadRune()
	if nil != err {
		return false
	}
	ts.Transfer(int(index[c-(127-c)*((127-c)>>31)]), l)
	return true
}

func (ts *CoreTokenScan) Scan() (token types.Token, offset int, position int) {
	defer func() {
		if r := recover(); nil != r {
			panic(r)
		}
	}()

	ts.dump = func(t types.Token, o int, l int) bool {
		token = t
		offset = o
		position = l
		return true
	}
	for ts.Next() {

	}
	if !ts.End() {
		token = -1
	}
	return
}

func (ts *CoreTokenScan) Transfer(i int, l int) {
	n := int(table[ts.state*16+i])
	transfer[n](ts, n, i, l)
}

func (ts *CoreTokenScan) reset() (ok bool) {
	if token[ts.state] > 0 {
		ok = ts.dump(token[ts.state], ts.offset, ts.position)
	}
	ts.Reset()
	return
}

func (ts *CoreTokenScan) End() bool {
	ok := ts.offset != ts.position
	if ok {
		ok = ts.reset()
	}
	return ok
}
