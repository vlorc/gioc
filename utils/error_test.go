// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package utils

import (
	"errors"
	"runtime"
	"testing"
)

func Test_Recover(t *testing.T) {
	var dst error
	src := errors.New("error test")
	go func() {
		defer Recover(&dst)
		panic(src)
	}()

	if runtime.Gosched(); src != dst {
		t.Error("can't matching error")
	}
}
