
# [Gioc](https://github.com/vlorc/gioc)

[简体中文](https://github.com/vlorc/gioc/blob/master/README_CN.md)

[![License](https://img.shields.io/:license-apache-blue.svg)](https://opensource.org/licenses/Apache-2.0)
[![codebeat badge](https://codebeat.co/badges/c41b426c-4121-4dc8-99c2-f1b60574be64)](https://codebeat.co/projects/github-com-vlorc-gioc-master)
[![Go Report Card](https://goreportcard.com/badge/github.com/vlorc/gioc)](https://goreportcard.com/report/github.com/vlorc/gioc)
[![GoDoc](https://godoc.org/github.com/vlorc/gioc?status.svg)](https://godoc.org/github.com/vlorc/gioc)
[![Build Status](https://travis-ci.org/vlorc/gioc.svg?branch=master)](https://travis-ci.org/vlorc/gioc?branch=master)
[![Coverage Status](https://coveralls.io/repos/github/vlorc/gioc/badge.svg?branch=master)](https://coveralls.io/github/vlorc/gioc?branch=master)

gioc is a lightweight Ioc framework,it provides register and factory and depend solution

## Features

* Dependency Resolve
* Dependency Inject
* Singleton/Transient Support
* Custom Tag
* Invoker Support
* [Lazy](https://github.com/vlorc/gioc/blob/master/examples/lazy/main.go) Load
* [Struct](https://github.com/vlorc/gioc/blob/master/examples/depend/main.go) Extends Support
* [Module](https://github.com/vlorc/gioc/blob/master/examples/module/main.go) Support

## Installing
	go get -u github.com/vlorc/gioc

## Quick Start

* Create Root Module
```golang
gioc.NewRootModule()
```

* Import Module
```golang
NewRootModule(
    Import(
        ConfigModule,
        ServerModule,
    )
)
```

* Declare Instance
```golang
NewRootModule(
    Declare(
        Instance(1), Id("id"),
        Instance("ioc"), Id("name"),
    ),
)
```

* Export Instance
```golang
NewModuleFactory(
    Export(
        Instance(1), Id("id"),
        Instance("ioc"), Id("name"),
    ),
)
```

## Examples

* Basic Module
```golang
import (
    . "github.com/vlorc/gioc"
    . "github.com/vlorc/gioc/module"
    . "github.com/vlorc/gioc/module/operation"
)

// config.go
var ConfigModule = NewModuleFactory(
    Export(
        Mapping(map[string]interface{}{
            "id": 1,
            "name": "ioc",
        }),
    ),
)

// main.go
func main() {
    NewRootModule(
        Import(ConfigModule),
        Bootstrap(func(param struct{ id int; name string }) {
            println("id: ", param.id," name: ",param.name)
        }),
    )
}
```

## License

This project is under the apache License. See the LICENSE file for the full license text.

## Interface

+ Provider
	+ provides Factory discovery
+ Factory
	+ responsible for generating Instance
	+ the basic plant has a value factory, method factory, agent factory, single factory, type factory
+ Mapper
	+ get the Factory by id
+ Binder
	+ the Factory is bound by id
	+ can be converted to read-only Mapper
+ Register
	+ as a connection to Factory and Selector
	+ provides the registration method, which eventually matches the Type to the Factory
	+ register custom Binder, Mapper, Factory
+ Dependency
	+ for target type dependency analysis, collection integration
	+ converted to an Injector by an instance
+ Injector
	+ and obtain the Instance padding based on the Dependency retrieval Provider
+ Builder
	+ is also a Factory
	+ use the Factory to get the instance and Injector to solve the Dependency
+ Container
	+ provides Register and Provider, and the parent container makes up traversal
	+ convert to read-only Provider
	+ convert to seal Container
+ Selector
	+ use type and id index Factory
	+ auto create Binder and Mapper
	+ index mode isolation
+ Module
    + import module
    + export factory
    + declare factory
    
# Roadmap
For details on planned features and future direction please refer to [roadmap](https://github.com/vlorc/gioc/blob/master/ROADMAP.md)

# Keyword

**dependency injection, inversion of control**

# Reference

