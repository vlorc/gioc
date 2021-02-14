package main

import (
	. "github.com/vlorc/gioc"
	. "github.com/vlorc/gioc/module"
	. "github.com/vlorc/gioc/module/operation"
)

// config.go
var ConfigModule = NewModuleFactory(
	Declare(
		Instance(new(int64)),
	),
	Export(
		Method(func(id *int64) int64 {
			return *id
		}), Id("id"),
		Method(func(id *int64) int64 {
			*id++
			return *id
		}), Id("inc"),
		Method(func(param struct{ inc int64 }) int64 {
			return param.inc
		}), Id("once"), Singleton(),
	),
)

// main.go
func main() {
	NewRootModule(
		Import(ConfigModule),
		Bootstrap(func(param struct {
			id   int64        `inject:"default(100)"`
			inc  func() int64 `inject:"lazy"`
			once int64
			req  func() int64 `inject:"request 'inc'"`
			ids  []int64      `inject:"make(10)"`
		}) {
			println("id: ", param.id, " inc: ", param.inc(), " once: ", param.once, "request: ", param.req(), param.req())
		}),
	)
}
