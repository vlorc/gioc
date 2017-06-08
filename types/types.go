// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

/*
Package type provides functionality for the interface and error defined.

*/
package types

import "reflect"

type DependencyFlag int

type ErrorCode int

type Error struct {
	Type    reflect.Type
	Name    string
	Code    ErrorCode
	Message string
}

type Provider interface {
	Resolve(interface{}, ...string) (interface{}, error)
	ResolveType(reflect.Type, string, int) (interface{}, error)
	ResolveNamed(interface{}, string, int) (interface{}, error)
	Assign(interface{}, ...string) error
	AssignNamed(interface{}, interface{}, string, int) error
}

type BeanFactory interface {
	Instance(Provider) (interface{}, error)
}

type Mapper interface {
	Resolve(string) (BeanFactory, error)
}

type Binder interface {
	Mapper
	AsMapper() Mapper
	Bind(string, BeanFactory) error
}

type BuilderFactory interface {
	Instance(BeanFactory, Dependency) (Builder, error)
}

type BinderFactory interface {
	Instance(reflect.Type) (Binder, error)
}

type RegisterFactory interface {
	Instance(BinderFactory) (Register, error)
}

type DependencyFactory interface {
	Instance(interface{}) (Dependency, error)
}

/*
type Component interface {

}
*/

type Register interface {
	AsMapper(reflect.Type) Mapper
	AsBinder(reflect.Type) Binder
	MapperOf(reflect.Type) Mapper
	BinderOf(reflect.Type) Binder
	RegisterBinder(Binder, interface{}) error
	RegisterMapper(Mapper, interface{}) error
	RegisterPointer(interface{}, ...string) error
	RegisterInstance(interface{}, ...string) error
	RegisterInterface(interface{}, ...string) error
	RegisterFactory(BeanFactory, interface{}, ...string) error
	RegisterMethod(BeanFactory, interface{}, interface{}, ...string) error
	//RegisterComponent(interface{},...string)error
}

type Container interface {
	Provider
	AsProvider() Provider
	AsRegister() Register
	Seal() Container
	Readonly() Container
	Parent() Container
	Child() Container
}

type DependencyScan interface {
	Type() reflect.Type
	Name() string
	Default() interface{}
	Flags() DependencyFlag
	Next() bool
	Index() int
	Test(interface{}) bool
	Depend() Dependency
}

type DependencyInject interface {
	DependencyScan
	SetInterface(interface{}) error
	SubInject(Provider) DependencyInject
}

type Dependency interface {
	Type() reflect.Type
	Length() int
	AsScan() DependencyScan
	AsInject(interface{}) DependencyInject
}

type Builder interface {
	BeanFactory
	AsFactory() BeanFactory
	Build(Provider, BeanFactory) (interface{}, error)
}

var ErrorType = reflect.TypeOf((*error)(nil)).Elem()

var BuilderType = reflect.TypeOf((*Builder)(nil)).Elem()
var DependencyType = reflect.TypeOf((*Dependency)(nil)).Elem()
var DependencyInjectType = reflect.TypeOf((*DependencyInject)(nil)).Elem()
var DependencyScanType = reflect.TypeOf((*DependencyScan)(nil)).Elem()
var ContainerType = reflect.TypeOf((*Container)(nil)).Elem()
var RegisterType = reflect.TypeOf((*Register)(nil)).Elem()
var DependencyFactoryType = reflect.TypeOf((*DependencyFactory)(nil)).Elem()
var RegisterFactoryType = reflect.TypeOf((*RegisterFactory)(nil)).Elem()
var BinderFactoryType = reflect.TypeOf((*BinderFactory)(nil)).Elem()
var BuilderFactoryType = reflect.TypeOf((*BuilderFactory)(nil)).Elem()
var BinderType = reflect.TypeOf((*Binder)(nil)).Elem()
var MapperType = reflect.TypeOf((*Mapper)(nil)).Elem()
var BeanFactoryFactoryType = reflect.TypeOf((*BeanFactory)(nil)).Elem()
var ProviderType = reflect.TypeOf((*Provider)(nil)).Elem()
var ParamType = reflect.TypeOf(([]reflect.Value)(nil))
