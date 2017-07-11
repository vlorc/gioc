// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package builder

import "github.com/vlorc/gioc/types"

func MakeIndexFullBefore(param []interface{}) func(*types.BuildContext) bool {
	return func(ctx *types.BuildContext) bool{
		ok := ctx.Descriptor.Index() < len(param)
		if ok {
			ctx.Inject.Convert(param[ctx.Descriptor.Index()])
		}
		return !ok
	}
}

func MakeNameFullBefore(table map[string]interface{}) func(*types.BuildContext) bool {
	return func(ctx *types.BuildContext) bool{
		val,ok := table[ctx.Descriptor.Name()]
		if ok {
			ctx.Inject.Convert(val)
		}
		return !ok
	}
}

func MakeDescriptorFullBefore(table map[string]types.Descriptor) func(*types.BuildContext) bool {
	return func(ctx *types.BuildContext) bool{
		val,ok := table[ctx.Descriptor.Name()]
		if ok {
			ctx.Descriptor = val
		}
		return true
	}
}