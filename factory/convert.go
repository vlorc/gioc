// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package factory

import (
	"github.com/vlorc/gioc/types"
	"github.com/vlorc/gioc/utils"
	"reflect"
)

func (cf *ConvertFactory) Instance(provider types.Provider) (instance interface{},err error) {
	instance,err = cf.factory.Instance(provider)
	if nil != err {
		return
	}
	if val := reflect.ValueOf(instance); val.Type().ConvertibleTo(cf.typ) {
		instance = val.Convert(cf.typ).Interface()
	} else {
		err = types.NewWithError(types.ErrTypeNotConvert,val)
	}
	return
}
