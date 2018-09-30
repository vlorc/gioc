package main

import (
	. "github.com/vlorc/gioc"
	. "github.com/vlorc/gioc/module"
	. "github.com/vlorc/gioc/module/operation"
)

type User struct {
	Id       int64 `inject:"lower"`
	personal ****struct {
		name   string
		age    int    `inject:"default(99)"`
		gender int    `inject:"optional"`
		email  string `inject:"optional"`
	} `inject:"extends"`
}

// config.go
var ConfigModule = NewModuleFactory(
	Export(
		Mapping(map[string]interface{}{
			"id":     int64(123),
			"name":   "admin_001",
			"gender": 1,
			"email":  "xxx@163.com",
		}),
	),
)

// main.go
func main() {
	NewRootModule(
		Import(ConfigModule),
		Bootstrap(func(param struct{ user User }) {
			println("id: ", param.user.Id, " name: ", (***param.user.personal).name)
		}),
	)
}
