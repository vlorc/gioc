// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package utils

import "bufio"

type EmptyLock struct{}

type Token int

type TokenScan struct {
	state int
	offset int
	position int
	input *bufio.Reader
	dump func(Token,int,int) bool
}
