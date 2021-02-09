// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package indexer

import "strconv"

type Int int

type String string

func (i Int) Value() int {
	return int(i)
}

func (i Int) String() string {
	return strconv.FormatInt(int64(i), 10)
}

func (i String) Value() int {
	v, _ := strconv.ParseInt(string(i), 10, 63)
	return int(v)
}

func (i String) String() string {
	return string(i)
}
