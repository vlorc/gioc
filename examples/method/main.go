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
	Id   int64
	Name string
	*Personal
}

// config.go
var ConfigModule = NewModuleFactory(
	Declare(
		Mapping(map[string]interface{}{
			"id":      new(int64),
			"name":    "admin",
			"gender":  1,
			"email":   "xxx@163.com",
			"Version": "1.0.1",
			"1":       "admin_x001",
			"2":       "admin_x102",
			"3":       "admin_x103",
		}),
	),
	Export(
		Method(func(param struct {
			id       *int64
			Name     string    `inject:"id('${id}') optional"`
			personal *Personal `inject:"extends"`
		}) *User {
			*param.id++
			return &User{
				Id:       *param.id,
				Name:     param.Name,
				Personal: param.personal,
			}
		}),
	),
)

// main.go
func main() {
	NewRootModule(
		Import(ConfigModule),
		Bootstrap(func(user *User) {
			fmt.Println("id:", user.Id, "name:", user.Name, "personal:", user.Personal)
		}),
		Bootstrap(func(user *User) {
			fmt.Println("id:", user.Id, "name:", user.Name, "personal:", user.Personal)
		}),
	)
}
