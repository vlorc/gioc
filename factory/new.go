// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package factory

import (
	"fmt"
	"github.com/vlorc/gioc/types"
	"reflect"
)

func (f *newFactory) Instance(types.Provider) (interface{}, error) {
	return reflect.New(f.typ).Interface(), nil
}

func (f *newFactory) String() string {
	return fmt.Sprintf("type(%s)", f.typ.String())
}
