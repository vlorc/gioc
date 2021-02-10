// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package types

import (
	"fmt"
	"github.com/vlorc/gioc/utils"
)

func (e *Error) Error() string {
	return e.String()
}

func (e *Error) String() string {
	return e.format(e)
}

func formatter(e *Error) string {
	format := errFormatTable[e.Code]
	str := fmt.Sprintf(format, e.Args...)
	e.format = func(*Error) string {
		return str
	}
	return str
}

func NewWithError(code ErrorCode, impType interface{}, args ...string) error {
	err := &Error{
		Code:   code,
		format: formatter,
		Args:   []interface{}{"", ""},
	}
	if typ := utils.TypeOf(impType); "" != typ.Name() {
		err.Args[0] = typ.Name()
	} else {
		err.Args[0] = typ.String()
	}
	if len(args) > 0 {
		err.Args[1] = args[0]
	}
	return err
}

func NewError(code ErrorCode, args ...interface{}) error {
	return &Error{
		Code:   code,
		Args:   args,
		format: formatter,
	}
}
