package main

import (
	"fmt"
	. "github.com/vlorc/gioc"
	. "github.com/vlorc/gioc/module"
	. "github.com/vlorc/gioc/module/operation"
)

type Personal struct {
	name   string `inject:"'name'"`
	age    *int   `inject:"'age' default"`
	gender int    `inject:"'gender' optional"`
	email  string `inject:"'email' optional"`
}

// config.go
var ConfigModule = NewModuleFactory(
	Export(
		Mapping(map[string]interface{}{
			"id":     int64(123),
			"name":   "admin",
			"gender": 1,
			"email":  "xxx@163.com",
		}),
	),
)

// main.go
func main() {
	NewRootModule(
		Import(ConfigModule),
		Bootstrap(func(param struct {
			Id       int64            `inject:"lower"`
			Strings  []string         `inject:"extends"`
			Personal func() *Personal `inject:"lazy extends"`
		}) {
			fmt.Println(param.Strings, "id:", param.Id, "name:", (*param.Personal()).name, *param.Personal())
		}),
	)
}
