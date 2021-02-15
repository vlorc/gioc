package main

import (
	. "github.com/vlorc/gioc"
	. "github.com/vlorc/gioc/module"
	. "github.com/vlorc/gioc/module/operation"
	"github.com/vlorc/gioc/types"
)

// config.go
var ConfigModule = NewModuleFactory(
	Condition(
		HavingValue(Equal("PRODUCT"), types.StringType, "ENV"),
		Export(
			Mapping(map[string]interface{}{
				"id":   1,
				"name": "product",
			}),
		), Bootstrap(func() {
			println("product")
		}),
	),
	Condition(
		Not(HavingValue(Equal("PRODUCT"), types.StringType, "ENV")),
		Export(
			Mapping(map[string]interface{}{
				"id":   2,
				"name": "debug",
			}),
		),
		Bootstrap(func() {
			println("debug")
		}),
	),
)

// main.go
func main() {
	NewRootModule(
		Declare(
			Instance("DEBUG"), Id("ENV"),
		),
		Import(ConfigModule),
		Bootstrap(func(param struct {
			id   int
			name string
		}) {
			println("id: ", param.id, " name: ", param.name)
		}),
	)
}
