package main

import (
	. "github.com/vlorc/gioc"
	. "github.com/vlorc/gioc/module"
	. "github.com/vlorc/gioc/module/operation"
)

// config.go
var ConfigModule = NewModuleFactory(
	Export(
		Mapping(map[string]interface{}{
			"id":   1,
			"name": "ioc",
		}),
	),
)

// main.go
func main() {
	NewRootModule(
		Import(ConfigModule),
		Bootstrap(func(param struct {
			id   int
			name string
		}) {
			println("id: ", param.id, " name: ", param.name)
		}),
	)
}
