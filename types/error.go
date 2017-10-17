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

func formatError(e *Error) string {
	format := errFormatTable[e.Code]
	str := fmt.Sprintf(format, e.Type.Name(), e.Name)
	e.format = func(*Error) string {
		return str
	}
	return str
}

func NewError(code ErrorCode, impType interface{}, args ...string) error {
	err := &Error{
		Type:   utils.TypeOf(impType),
		Name:   DEFAULT_NAME,
		Code:   code,
		format: formatError,
	}
	if len(args) > 0 {
		err.Name = args[0]
	}
	return err
}
