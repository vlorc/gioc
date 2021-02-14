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

func (f *makeSliceFactory) Instance(types.Provider) (interface{}, error) {
	return reflect.MakeSlice(f.typ, f.length, f.length).Interface(), nil
}

func (f *makeSliceFactory) String() string {
	return fmt.Sprintf("type(%s) length(%d)", f.typ.String(), f.length)
}

func (f *makeMapFactory) Instance(types.Provider) (interface{}, error) {
	return reflect.MakeMap(f.typ).Interface(), nil
}

func (f *makeMapFactory) String() string {
	return fmt.Sprintf("type(%s)", f.typ.String())
}

func (f *makeChanFactory) Instance(types.Provider) (interface{}, error) {
	return reflect.MakeChan(f.typ, f.length).Interface(), nil
}

func (f *makeChanFactory) String() string {
	return fmt.Sprintf("type(%s) buffer(%d)", f.typ.String(), f.length)
}
