// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package operation

import (
	"github.com/vlorc/gioc/factory"
	"reflect"
	"github.com/vlorc/gioc/builder"
	"github.com/vlorc/gioc/utils"
)

func Interface(val interface{},typ ...interface{}) DeclareHandle {
	return func(ctx *DeclareContext) {
		src := reflect.ValueOf(val)
		if reflect.Interface == src.Kind(){
			toInterface(ctx,src)
			return
		}
		if reflect.Ptr != src.Kind(){
			 return
		}
		switch src.Type().Elem().Kind() {
		case reflect.Interface:
			if !src.IsNil() {
				toInterface(ctx,src.Elem())
			}
		case reflect.Struct:
			if len(typ) > 0 {
				toStruct(ctx,src,typ[0])
			}
		}
	}
}

func toInterface(ctx *DeclareContext,val reflect.Value) {
	if !val.IsNil() {
		ctx.Factory = factory.NewValueFactory(val.Interface())
		ctx.Type = val.Type()
	}
}

func toStruct(ctx *DeclareContext,val reflect.Value,typ interface{}) {
	if val.IsNil() {
		ctx.Factory = factory.NewTypeFactory(val.Type().Elem())
	} else {
		ctx.Factory = factory.NewValueFactory(val.Elem().Interface())
	}
	if toDependency(ctx,typ) {
		ctx.Factory = builder.NewBuilder(ctx.Factory,ctx.Depend).AsFactory()
	}
	ctx.Factory = factory.NewConvertFactory(ctx.Factory,utils.TypeOf(typ))
}
