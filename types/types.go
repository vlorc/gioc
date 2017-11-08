// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

/*
Package type provides functionality for the interface and error defined.

*/
package types

import "reflect"

type Provider interface {
	Resolve(interface{}, ...string) (interface{}, error)
	ResolveType(reflect.Type, string, int) (interface{}, error)
	ResolveNamed(interface{}, string, int) (interface{}, error)
	Assign(interface{}, ...string) error
	AssignType(reflect.Value, reflect.Type, string, int) error
	AssignNamed(interface{}, interface{}, string, int) error
}

type BeanFactory interface {
	Instance(Provider) (interface{}, error)
}

type Mapper interface {
	Resolve(string) BeanFactory
}

type Binder interface {
	Mapper
	AsMapper() Mapper
	Bind(string, BeanFactory) error
}

type SelectorSetter interface {
	AsMapper(reflect.Type) Mapper
	AsBinder(reflect.Type) Binder

	SetBinder(reflect.Type, Binder) error
	SetFactory(reflect.Type, string, BeanFactory) error
}

type SelectorGetter interface {
	MapperOf(reflect.Type) Mapper
	BinderOf(reflect.Type) Binder
	FactoryOf(reflect.Type, string) BeanFactory
}

type Selector interface {
	SelectorSetter
	SelectorGetter
}

type ProviderFactory interface {
	Instance(SelectorGetter,Provider) (Provider, error)
}

type BuilderFactory interface {
	Instance(BeanFactory, Dependency) (Builder, error)
}

type BinderFactory interface {
	Instance(reflect.Type) (Binder, error)
}

type RegisterFactory interface {
	Instance(SelectorSetter) (Register, error)
}

type DependencyFactory interface {
	Instance(interface{}) (Dependency, error)
}

type SelectorFactory interface {
	Instance(BinderFactory) (Selector, error)
}

type InvokerFactory interface {
	Instance(interface{}, Builder) (Invoker, error)
}

type Register interface {
	RegisterBinder(Binder, interface{}) error
	RegisterMapper(Mapper, interface{}) error
	RegisterPointer(interface{}, ...string) error
	RegisterInstance(interface{}, ...string) error
	RegisterInterface(interface{}, ...string) error
	RegisterFactory(BeanFactory, interface{}, ...string) error
	RegisterMethod(BeanFactory, interface{}, interface{}, ...string) error
}

type Container interface {
	AsProvider() Provider
	AsRegister() Register
	Seal() Container
	Readonly() Container
	NewChild() Container
}

type DescriptorGetter interface {
	Type() reflect.Type
	Name() string
	Default() reflect.Value
	Flags() DependencyFlag
	Index() int
	Depend() Dependency
}

type DescriptorSetter interface {
	SetType(reflect.Type)
	SetName(string)
	SetDefault(reflect.Value)
	SetFlags(DependencyFlag)
	SetIndex(int)
	SetDepend(Dependency)
}

type Descriptor interface {
	DescriptorSetter
	DescriptorGetter
}

type DependencyScan interface {
	DescriptorGetter
	Reset()
	Next() bool
	AsDescriptorGetter() DescriptorGetter
	AsDescriptorSetter() DescriptorSetter
	Test(interface{}) bool
}

type DependencyInject interface {
	DependencyScan
	AsValue() reflect.Value
	Convert(interface{})
	SetInterface(interface{})
	SetValue(reflect.Value)
	SubInject(Provider) DependencyInject
}

type Dependency interface {
	Type() reflect.Type
	Length() int
	AsScan() DependencyScan
	AsInject(interface{}) DependencyInject
}

type PropertySetter interface {
	Set(DescriptorGetter, reflect.Value)
}

type PropertyGetter interface {
	Get(DescriptorGetter) reflect.Value
}

type Reflect interface {
	PropertySetter
	PropertyGetter
}

type Builder interface {
	BeanFactory
	AsFactory() BeanFactory
	Build(Provider, ...func(*BuildContext)) (interface{}, error)
}

type Invoker interface {
	Apply(...interface{}) []reflect.Value
	ApplyWith(Provider, ...interface{}) []reflect.Value
}

var ErrorType = reflect.TypeOf((*error)(nil)).Elem()
var ProviderType = reflect.TypeOf((*Provider)(nil)).Elem()
var RegisterFactoryType = reflect.TypeOf((*RegisterFactory)(nil)).Elem()
var BinderFactoryType = reflect.TypeOf((*BinderFactory)(nil)).Elem()
