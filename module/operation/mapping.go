// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package operation

import (
	"reflect"
)

func Mapping(val interface{}) DeclareHandle {
	return func(ctx *DeclareContext) {
		switch r := val.(type) {
		case map[string]interface{}:
			mappingMap(ctx,r)
		case []interface{}:
			mappingArray(ctx,r)
		default:
			mappingStruct(ctx,val)
		}
	}
}

func mapping(ctx *DeclareContext,val interface{}) {
	if reflect.Func == reflect.TypeOf(val).Kind() {
		Method(val)(ctx)
	} else {
		Instance(val)(ctx)
	}
}

func mappingStruct(ctx *DeclareContext,val interface{}) {

}

func mappingArray(ctx *DeclareContext,array []interface{}) {
	for i,l := 0,len(array) >> 1; i < l;i++ {
		mapping(ctx,array[i * 2 + 1])
		ctx.Name = array[i * 2 + 0].(string)
	}
}

func mappingMap(ctx *DeclareContext,table map[string]interface{}) {
	for k,v := range table {
		mapping(ctx,v)
		ctx.Name = k
	}
}
