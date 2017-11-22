// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package utils

import (
	"errors"
	"testing"
	"sync"
)

func Test_Recover(t *testing.T) {
	var dst error
	src := errors.New("error test")
	wait := sync.WaitGroup{}
	wait.Add(1)
	go func() {
		defer wait.Done()
		defer Recover(&dst)
		panic(src)
	}()

	if wait.Wait(); src != dst {
		t.Error("can't matching error")
	}
}
