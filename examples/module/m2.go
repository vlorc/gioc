package main

import (
	"sync/atomic"
	. "github.com/vlorc/gioc/module"
	. "github.com/vlorc/gioc/module/operation"
)

var Module2 = NewModuleFactory(
	Declare(
		Instance(new(int64)), Id("id"),
	),
	Export(
		Method(nextId), Id("id"),
	),
)

func nextId(param struct{id *int64}) int64 {
	return atomic.AddInt64(param.id,1)
}
