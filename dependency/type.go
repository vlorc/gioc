// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package dependency

import (
	"github.com/vlorc/gioc/factory"
	"github.com/vlorc/gioc/text"
	"github.com/vlorc/gioc/types"
	"github.com/vlorc/gioc/types/indexer"
	"github.com/vlorc/gioc/utils"
	"reflect"
	"strings"
)

func (df *CoreDependencyFactory) resolveArray(typ reflect.Type, val reflect.Value) types.Dependency {
	if val.Len() <= 0 {
		return nil
	}

	var dep []types.DependencyDescriptor
	for i, n := 0, val.Len(); i < n; i++ {
		dep = df.appendValue(dep, val.Index(i), i)
	}
	if len(dep) > 0 {
		return NewArrayDependency(typ, dep)
	}
	return nil
}

func (df *CoreDependencyFactory) appendValue(dep []types.DependencyDescriptor, val reflect.Value, index int) []types.DependencyDescriptor {
	typ := val.Type()
	if reflect.Ptr == typ.Kind() {
		typ = typ.Elem()
	}

	return append(dep, types.DependencyDescriptor{
		Factory: factory.NewResolveFactory(typ),
		Index:   indexer.Int(index),
		Flags:   0,
		Type:    typ,
	})
}

func (df *CoreDependencyFactory) resolveFunc(typ reflect.Type, val reflect.Value) types.Dependency {
	if typ.NumIn() <= 0 {
		return nil
	}

	var dep []types.DependencyDescriptor
	for i, n := 0, typ.NumIn(); i < n; i++ {
		dep = df.appendParam(dep, typ.In(i), i)
	}
	if len(dep) > 0 {
		return NewFuncDependency(typ, dep)
	}
	return nil
}

func (df *CoreDependencyFactory) appendParam(dep []types.DependencyDescriptor, typ reflect.Type, index int) []types.DependencyDescriptor {
	if b := df.checkAnonymous(typ); nil != b {
		dep = append(dep, types.DependencyDescriptor{
			Factory: b,
			Index:   indexer.Int(index),
			Flags:   types.DEPENDENCY_FLAG_EXTENDS,
			Type:    typ,
		})
	} else {
		dep = append(dep, types.DependencyDescriptor{
			Factory: factory.NewResolveFactory(typ),
			Index:   indexer.Int(index),
			Flags:   0,
			Type:    typ,
		})
	}
	return dep
}

func (df *CoreDependencyFactory) appendKey(dep []types.DependencyDescriptor, m reflect.Value, k reflect.Value) []types.DependencyDescriptor {
	if reflect.String != k.Kind() {
		return dep
	}

	typ := m.MapIndex(k).Type()
	if reflect.Ptr == typ.Kind() {
		typ = typ.Elem()
	}

	name := k.String()
	return append(dep, types.DependencyDescriptor{
		Factory: factory.NewResolveFactory(typ, types.RawStringFactory(name)),
		Index:   indexer.String(name),
		Flags:   0,
		Type:    typ,
	})
}

func (df *CoreDependencyFactory) resolveMap(typ reflect.Type, val reflect.Value) types.Dependency {
	if val.Len() <= 0 {
		return nil
	}

	var dep []types.DependencyDescriptor
	for _, k := range val.MapKeys() {
		dep = df.appendKey(dep, val, k)
	}
	if len(dep) > 0 {
		return NewMapDependency(typ, dep)
	}
	return nil
}

func (df *CoreDependencyFactory) resolveStruct(typ reflect.Type, _ reflect.Value) types.Dependency {
	if "" == typ.Name() {
		return df.anonymousToDependency(typ)
	}

	df.lock.RLock()
	dep := df.pool[typ]
	df.lock.RUnlock()
	if nil != dep {
		return dep
	}

	if dep = df.namedToDependency(typ); nil != dep {
		df.lock.Lock()
		df.pool[typ] = dep
		df.lock.Unlock()
	}
	return dep
}

func (df *CoreDependencyFactory) structToDependency(typ reflect.Type, skip func(string) bool) types.Dependency {
	var dep []types.DependencyDescriptor
	ctx := &types.ParseContext{
		Factory: df,
		Scan:    text.NewTokenScan(),
	}
	ctx.Dependency.Origin.Target = typ

	for i, n := 0, typ.NumField(); i < n; i++ {
		dep = df.appendField(dep, typ.Field(i), ctx, skip)
		ctx.Dependency.Reset()
	}
	if len(dep) > 0 {
		return NewStructDependency(typ, dep)
	}
	return nil
}

func (df *CoreDependencyFactory) namedToDependency(typ reflect.Type) types.Dependency {
	return df.structToDependency(typ, func(tag string) bool {
		return "" != tag && "-" != tag
	})
}

func (df *CoreDependencyFactory) anonymousToDependency(typ reflect.Type) (dep types.Dependency) {
	return df.structToDependency(typ, func(tag string) bool {
		return "-" != tag
	})
}

func (df *CoreDependencyFactory) appendField(
	dep []types.DependencyDescriptor,
	field reflect.StructField,
	ctx *types.ParseContext,
	skip func(string) bool) []types.DependencyDescriptor {

	str := field.Tag.Get(df.tag)
	if !skip(str) {
		return dep
	}

	if len(field.Index) > 1 {
		utils.Panic(types.NewWithError(types.ErrIndexNotSupport, field.Type))
	}

	if uint(field.Name[0]-65) >= uint(26) {
		ctx.Dependency.Flags |= types.DEPENDENCY_FLAG_UNSAFE
	}
	ctx.Dependency.Origin.Type = field.Type
	ctx.Dependency.Origin.Name = field.Name
	ctx.Dependency.Origin.Index = indexer.Int(field.Index[0])
	ctx.Dependency.Type = field.Type

	if "" != str {
		ctx.Scan.SetInput(strings.NewReader(str))
		ctx.Dump = func(i1 int, i2 int) string {
			return str[i1:i2]
		}

		if err := df.parser.Resolve(ctx); nil != err {
			utils.Panic(err)
		}
	}

	return df.appendContext(dep, &ctx.Dependency)
}

func (df *CoreDependencyFactory) checkAnonymous(typ reflect.Type) types.BeanFactory {
	if t := utils.IndirectType(typ); reflect.Struct == t.Kind() && "" == t.Name() {
		dep := df.anonymousToDependency(t)
		return factory.NewDependencyFactory(factory.NewTypeFactory(t), dep, __elem(dep.Type(), typ))
	}
	return nil
}

func (df *CoreDependencyFactory) appendContext(dep []types.DependencyDescriptor, ctx *types.DependencyContext) []types.DependencyDescriptor {
	for i := range ctx.Before {
		if err := ctx.Before[i](ctx); nil != err {
			ctx.Error = err
			return dep
		}
	}

	if nil == ctx.Name {
		ctx.Name = []types.StringFactory{types.RawStringFactory(ctx.Origin.Name)}
	}
	if nil != ctx.Default {
		ctx.Factory = ctx.Default(ctx)
	} else {
		ctx.Factory = factory.NewResolveFactory(ctx.Type, ctx.Name...)
	}

	if ctx.Factory, ctx.Error = ctx.Wrapper.Instance(ctx.Factory); nil != ctx.Error {
		return dep
	}

	for i := range ctx.After {
		if err := ctx.After[i](ctx); nil != err {
			ctx.Error = err
			return dep
		}
	}

	dep = append(dep, types.DependencyDescriptor{
		Factory: ctx.Factory,
		Index:   ctx.Origin.Index,
		Flags:   ctx.Flags,
		Type:    ctx.Origin.Type,
	})
	return dep
}

func __elem(src, dst reflect.Type) func(interface{}) interface{} {
	if reflect.Ptr == dst.Kind() || reflect.PtrTo(src) == dst {
		return func(v interface{}) interface{} {
			return v
		}
	}
	return func(v interface{}) interface{} {
		return utils.Elem(v, dst)
	}
}
