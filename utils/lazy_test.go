// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package utils

import (
	"testing"
)

func Test_Lazy(t *testing.T) {
	i := 0
	get := Lazy(func() int {
		i++
		return i
	}).(func() int)

	for i := 0; i < 10; i++ {
		if get() != 1 {
			t.Error("can't matching value")
		}
	}
}
