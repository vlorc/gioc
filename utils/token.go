// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package utils

import (
	"bufio"
	"errors"
	"io"
)

var table = []byte{
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	13, 1, 3, 8, 13, 14, 15, 16, 17, 18, 19, 19,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	3, 3, 4, 3, 5, 3, 3, 3, 3, 3, 3, 3,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6,
	6, 6, 7, 6, 5, 6, 6, 6, 6, 6, 6, 6,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	8, 8, 8, 9, 10, 8, 8, 8, 8, 8, 8, 8,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11,
	11, 11, 11, 12, 10, 11, 11, 11, 11, 11, 11, 11,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	13, 0, 0, 0, 13, 0, 0, 0, 13, 13, 13, 13,
	0, 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 2, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 2, 0, 0, 0, 0,
	13, 0, 0, 0, 0, 0, 0, 0, 13, 13, 19, 19,
	13, 0, 0, 0, 0, 0, 0, 0, 13, 13, 19, 19,
	2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 19, 19,
}

var token = []Token{
	4:  TOKEN_CHART,
	7:  TOKEN_ECHART,
	9:  TOKEN_STRING,
	12: TOKEN_ESTRING,
	13: TOKEN_IDENT,
	14: TOKEN_LPAREN,
	15: TOKEN_RPAREN,
	16: TOKEN_COMMA,
	19: TOKEN_NUMBER,
}

var index [128]byte

var transfer []func(*TokenScan, int, int, int)

func init() {
	for i, v := range ` '"\(),-+` {
		index[v] = byte(i + 1)
	}
	for i := 0; i <= 32; i++ {
		index[i] = 1
	}
	for i := 48; i <= 57; i++ {
		index[i] = 10
	}

	transfer = make([]func(*TokenScan, int, int, int), len(token))
	for i := range transfer {
		transfer[i] = func(d *TokenScan, n int, _ int, l int) {
			d.position += l
			d.state = n
		}
	}

	transfer[2] = func(d *TokenScan, _ int, _ int, _ int) {
		panic(errors.New("Illegal char"))
	}

	transfer[1] = func(d *TokenScan, n int, _ int, l int) {
		d.position += l
		d.offset += l
		d.state = n
	}

	transfer[0] = func(d *TokenScan, n int, i int, l int) {
		ok := d.reset()
		d.Transfer(i, l)
		if ok {
			panic(nil)
		}
	}
}

func NewTokenScan() *TokenScan {
	return &TokenScan{}
}

func (ts *TokenScan) Init(steam io.Reader) {
	input := bufio.NewReader(steam)
	ts.input = input
}

func (ts *TokenScan) Offset() int {
	return ts.offset
}

func (ts *TokenScan) Position() int {
	return ts.position
}

func (ts *TokenScan) Begin() {
	ts.position = 0
	ts.Reset()
}

func (ts *TokenScan) Reset() {
	ts.state = 1
	ts.offset = ts.position
}

func (ts *TokenScan) Next() bool {
	c, l, err := ts.input.ReadRune()
	if nil != err {
		return false
	}
	ts.Transfer(int(index[c-(127-c)*((127-c)>>31)]), l)
	return true
}

func (ts *TokenScan) Scan() (token Token, offset int, position int) {
	defer func() {
		if r := recover(); nil != r {
			panic(r)
		}
	}()

	ts.dump = func(t Token, o int, l int) bool {
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

func (ts *TokenScan) Transfer(i int, l int) {
	n := int(table[ts.state*12+i])
	transfer[n](ts, n, i, l)
}

func (ts *TokenScan) reset() (ok bool) {
	if token[ts.state] > 0 {
		ok = ts.dump(token[ts.state], ts.offset, ts.position)
	}
	ts.Reset()
	return
}

func (ts *TokenScan) End() bool {
	ok := ts.offset != ts.position
	if ok {
		ok = ts.reset()
	}
	return ok
}
