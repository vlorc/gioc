// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package types

import (
	"fmt"
	"github.com/vlorc/gioc/utils"
)

func (this *Error) Error() string {
	return this.Message
}

func (this *Error) String() string {
	return this.Message
}

func NewError(code ErrorCode, impType interface{}, args ...string) error {
	typ := utils.TypeOf(impType)
	name := ""
	format := errFormatTable[code]
	if len(args) > 0 {
		name = args[0]
	}

	err := &Error{
		Type:    typ,
		Name:    name,
		Code:    code,
		Message: fmt.Sprintf(format, typ.Name(), name),
	}

	return err
}
