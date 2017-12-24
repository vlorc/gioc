package main

import (
	. "github.com/vlorc/gioc/module"
	. "github.com/vlorc/gioc/module/operation"
)

var Module2 = NewModuleFactory(
	Declare(
		Method(getId(0)), Name("id"), Export(),
	),
)

func getId(id int64) func() int64 {
	return func() int64 {
		id++
		return id
	}
}
