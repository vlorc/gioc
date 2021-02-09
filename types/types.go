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

	Get(typ reflect.Type, name string, deep int) BeanFactory

	Range(callback func(GeneralFactory) bool, types ...reflect.Type) bool

	AsSelector() SelectorGetter
}

type Module interface{}

type ModuleFactory interface {
	Instance(func() Container) (Module, error)
}

type EventListener interface {
	On(string, interface{}) error
	Emit(string, ...interface{}) error
	OnWith(func() Provider, string, interface{}) error
	EmitWith(func() Provider, string, ...interface{}) error
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

type StringFactory interface {
	// get an instance by provider
	Instance(provider Provider) (string, error)
}

type BeanFactory interface {
	// get an instance by provider
	Instance(provider Provider) (interface{}, error)
}

type GeneralFactory interface {
	BeanFactory

	Type() reflect.Type

	Name() string
}

// selector write
type SelectorSetter interface {
	Set(typ reflect.Type, name string, factory BeanFactory)

	Add(typ reflect.Type, name string, factory BeanFactory)

	Put(typ reflect.Type, name string, factory BeanFactory) bool
}

/// selector read only
type SelectorGetter interface {
	Get(typ reflect.Type, name string) GeneralFactory

	Range(callback func(GeneralFactory) bool, types ...reflect.Type) bool
}

// selector
type Selector interface {
	SelectorSetter
	SelectorGetter
}

// provider factory
type ProviderFactory interface {
	// create a builder by selector and the parent provider
	Instance(parent Provider, getter SelectorGetter) (Provider, error)
}

// register factory
type RegisterFactory interface {
	// create a register
	Instance(setter Selector) (Register, error)
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
	Instance() (Selector, error)
}

// invoker factory
type InvokerFactory interface {
	// method and builder
	Instance(method interface{}, dependency Dependency) (Invoker, error)
}

// register
type Register interface {
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

	AsSelector() Selector
}

// the container is provider and register
type Container interface {
	AsProvider() Provider
	AsRegister() Register
	Seal() Container
	Readonly() Container
	NewChild() Container
	Name() string
}

// dependency scan
type DependencyScan interface {
	Reset()

	Next() bool

	Index() Indexer

	Flags() DependencyFlag

	Type() reflect.Type

	Factory(provider Provider) BeanFactory
}

// dependency
type Dependency interface {
	// get raw type
	Type() reflect.Type
	// get dependency scan length
	Length() int
	// convert to dependency scan
	AsScan() DependencyScan

	AsReflect(imp interface{}) Reflect
}

type Indexer interface {
	Value() int
	String() string
}

type ReflectFactory interface {
	Instance(imp interface{}) (Reflect, error)
}

// be use for setting instance
type Reflect interface {
	// set value by dependency descriptor
	Set(Indexer, reflect.Value)
	// get value by dependency descriptor
	Get(Indexer) reflect.Value
}

// invoker
type Invoker interface {
	// call method
	Apply(...interface{}) ([]reflect.Value, error)
	// call method and instance fill
	ApplyWith(Provider, ...interface{}) ([]reflect.Value, error)
}

var Type = reflect.TypeOf((*reflect.Type)(nil)).Elem()
var ErrorType = reflect.TypeOf((*error)(nil)).Elem()
var ProviderType = reflect.TypeOf((*Provider)(nil)).Elem()
var RegisterFactoryType = reflect.TypeOf((*RegisterFactory)(nil)).Elem()
