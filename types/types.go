// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

/*
Package type provides functionality for the interface and error defined.

*/
package types

import (
	"io"
	"reflect"
)

type Param interface {
	String() string
	Number() int64
	Float() float64
	Boolean() bool
	Value() reflect.Value
	Kind() Kind
}

type TokenScan interface {
	Offset() int
	Position() int
	Reset()
	Begin()
	End() bool
	Next() bool
	Scan() (Token, int, int)
	SetInput(io.Reader) TokenScan
}

type TextParser interface {
	Resolve(*ParseContext) error
}

type Provider interface {
	// get instance by type pointer and name
	// impType must be a pointer
	Resolve(impType interface{}, name ...string) (interface{}, error)
	// get instance by type and name and deep
	ResolveType(typ reflect.Type, name string, deep int) (interface{}, error)
	// get instance by type pointer and name and deep
	// impType must be a pointer
	ResolveNamed(impType interface{}, name string, deep int) (interface{}, error)
	// assigned to instance by type pointer and name
	// dst must be a pointer
	Assign(dst interface{}, name ...string) error
	// assigned to instance by type and name and deep
	// dst must be can set
	// typ for nil using dst type
	AssignType(dst reflect.Value, typ reflect.Type, name string, deep int) error
	// assigned to instance by type pointer and name and deep
	// dst and impType must be a pointer
	// impType for nil using dst type
	AssignNamed(dst interface{}, impType interface{}, name string, deep int) error
}

type Module interface{}

type ModuleFactory interface {
	Instance(func() Container) (Module, error)
}

type ParamFactory interface {
	// create a Param by token and value
	Instance(Token, string) (Param, error)
}

type TokenScanFactory interface {
	Instance() (TokenScan, error)
}

type TextParserFactory interface {
	// create a textParser by table and param factory
	Instance(table map[string][]IdentHandle, factory ParamFactory) (TextParser, error)
}

type BeanFactory interface {
	// get an instance by provider
	Instance(provider Provider) (interface{}, error)
}

// a read only binder
type Mapper interface {
	// get an factory by name
	Resolve(name string) BeanFactory
}

// the relation between name and factory
type Binder interface {
	Mapper
	// convert to a read only mapper
	AsMapper() Mapper
	// set a factory by name
	Bind(name string, factory BeanFactory) error
}

// selector write
type SelectorSetter interface {
	// create and set a mapper by type
	AsMapper(reflect.Type) Mapper
	// create and set a binder by type
	AsBinder(reflect.Type) Binder
	// set a custom binder by type
	SetBinder(reflect.Type, Binder) error
	// set a custom factory by type
	SetFactory(reflect.Type, string, BeanFactory) error
}

/// selector read only
type SelectorGetter interface {
	// get mapper by type
	MapperOf(reflect.Type) Mapper
	// get binder by type
	BinderOf(reflect.Type) Binder
	// get factory by type and name
	FactoryOf(reflect.Type, string) BeanFactory
}

// selector
type Selector interface {
	SelectorSetter
	SelectorGetter
}

// provider factory
type ProviderFactory interface {
	// create a builder by selector and the parent provider
	Instance(getter SelectorGetter, parent Provider) (Provider, error)
}

// builder factory
type BuilderFactory interface {
	// create a builder by factory and depend
	Instance(factory BeanFactory, depend Dependency) (Builder, error)
}

// binder factory
type BinderFactory interface {
	// create a binder by typ
	Instance(typ reflect.Type) (Binder, error)
}

// register factory
type RegisterFactory interface {
	// create a register by setter
	Instance(setter SelectorSetter) (Register, error)
}

// dependency factory
type DependencyFactory interface {
	// dependency resolve
	// imp can be functions\struct pointer\arrays\maps
	Instance(imp interface{}) (Dependency, error)
}

// selector factory
type SelectorFactory interface {
	// create a selector by binder
	Instance(binder BinderFactory) (Selector, error)
}

// invoker factory
type InvokerFactory interface {
	// method and builder
	Instance(method interface{}, param Builder) (Invoker, error)
}

// register
type Register interface {
	// register a custom binder
	// impType must be a pointer
	RegisterBinder(binder Binder, impType interface{}) error
	// register a custom mapper
	// impType must be a pointer
	RegisterMapper(mapper Mapper, impType interface{}) error
	// register a pointer
	RegisterPointer(pointer interface{}, name ...string) error
	// register an instance,use the type of instance
	RegisterInstance(instance interface{}, name ...string) error
	// register an interface,use the type of interface
	RegisterInterface(instance interface{}, name ...string) error
	// register a custom factory
	// impType must be a pointer
	RegisterFactory(factory BeanFactory, impType interface{}, name ...string) error
	// register a custom method
	// impType must be a pointer,it's the return value type of method
	RegisterMethod(factory BeanFactory, method interface{}, impType interface{}, name ...string) error
}

// the container is provider and register
type Container interface {
	AsProvider() Provider
	AsRegister() Register
	Seal() Container
	Readonly() Container
	NewChild() Container
}

// dependency descriptor read only
type DescriptorGetter interface {
	Type() reflect.Type
	Name() string
	Default() reflect.Value
	Flags() DependencyFlag
	Index() int
	Depend() Dependency
}

// dependency descriptor write
type DescriptorSetter interface {
	SetType(reflect.Type)
	SetName(string)
	SetDefault(reflect.Value)
	SetFlags(DependencyFlag)
	SetIndex(int)
	SetDepend(Dependency)
}

// dependency descriptor
type Descriptor interface {
	DescriptorSetter
	DescriptorGetter
}

// dependency scan
type DependencyScan interface {
	DescriptorGetter
	// position reset
	Reset()
	// has next
	Next() bool
	// convert to a read only descriptor
	AsDescriptorGetter() DescriptorGetter
	// convert to a write descriptor
	AsDescriptorSetter() DescriptorSetter
	// the test can be set by impTyp
	// impTyp must be a pointer
	Test(impTyp interface{}) bool
}

// dependency inject
type DependencyInject interface {
	DependencyScan
	// convert to value
	AsValue() reflect.Value
	// convert and set interface
	Convert(interface{})
	// set interface
	SetInterface(interface{})
	// set value
	SetValue(reflect.Value)
	// get a inject by provider
	SubInject(provider Provider) DependencyInject
}

// dependency
type Dependency interface {
	// get raw type
	Type() reflect.Type
	// get dependency scan length
	Length() int
	// convert to dependency scan
	AsScan() DependencyScan
	// convert to dependency inject by instance
	AsInject(interface{}) DependencyInject
}

// property setter
type PropertySetter interface {
	// set value by dependency descriptor
	Set(DescriptorGetter, reflect.Value)
}

type PropertyGetter interface {
	// get value by dependency descriptor
	Get(DescriptorGetter) reflect.Value
}

// be use for setting instance
type Reflect interface {
	PropertySetter
	PropertyGetter
}

// builder is instance fill procedure
type Builder interface {
	AsFactory() BeanFactory
	Build(Provider, ...func(*BuildContext)) (interface{}, error)
}

// invoker
type Invoker interface {
	// call method
	Apply(...interface{}) []reflect.Value
	// call method and instance fill
	ApplyWith(Provider, ...interface{}) []reflect.Value
}

var ErrorType = reflect.TypeOf((*error)(nil)).Elem()
var ProviderType = reflect.TypeOf((*Provider)(nil)).Elem()
var RegisterFactoryType = reflect.TypeOf((*RegisterFactory)(nil)).Elem()
var BinderFactoryType = reflect.TypeOf((*BinderFactory)(nil)).Elem()
