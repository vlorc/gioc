package main

import (
	"fmt"
	. "github.com/vlorc/gioc"
	. "github.com/vlorc/gioc/module"
	. "github.com/vlorc/gioc/module/operation"
)

type Personal struct {
	Name   string `inject:"'name'"`
	Age    int    `inject:"'age' optional"`
	Gender int    `inject:"'gender' optional"`
	Email  string `inject:"'email' optional"`
}

type User struct {
	Id        int64 `inject:"'id'"`
	*Personal `inject:"extends"`
}

// config.go
var ConfigModule = NewModuleFactory(
	Declare(
		Mapping(map[string]interface{}{
			"id":      new(int64),
			"name":    "admin_001",
			"gender":  1,
			"email":   "xxx@163.com",
			"Version": "1.0.1",
		}),
	),
	Export(
		Method(func(param struct {
			id       *int64
			Name     string    `inject:"lower"`
			personal *Personal `inject:"extends"`
		}) *User {
			*param.id++
			return &User{
				Id:       *param.id,
				Personal: param.personal,
			}
		}), Singleton(),
	),
)

// main.go
func main() {
	NewRootModule(
		Import(ConfigModule),
		Bootstrap(func(user *User) {
			fmt.Println("id:", user.Id, "personal:", user.Personal)
		}),
		Bootstrap(func(user *User) {
			fmt.Println("id:", user.Id, "personal:", user.Personal)
		}),
	)
}
