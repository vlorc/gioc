// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package utils

import (
	"errors"
	"fmt"
)

func Recover(err *error){
	r := recover()
	if r == nil {
		return
	}
	switch x := r.(type) {
	case error:
		*err = x
	case string:
		*err = errors.New(x)
	default:
		*err = fmt.Errorf("Unknown panic %s",r)
	}
	return
}