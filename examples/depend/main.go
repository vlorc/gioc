package main

import (
	"fmt"
	"github.com/vlorc/gioc"
	"github.com/vlorc/gioc/factory"
	"github.com/vlorc/gioc/types"
)

type Personal struct {
	Name   string `inject:"'name'"`
	Age    int    `inject:"'age' optional"`
	Gender int    `inject:"'gender' optional"`
	Email  string `inject:"'email' optional"`
}

type Identity struct {
	Username string `inject:"'username'"`
	Password string `inject:"'password'"`
}

type User struct {
	id0       int64
	_id0       int64
	Id1       int64 `json:"'id'"`
	Id2       int64 `inject`
	Id3       int64 `inject:""`
	Id       int64 `inject:"'id'"`
	Identity `inject:"extends"`
	Personal `inject:"extends"`
}

func main() {
	container := gioc.NewRootContainer()

	for k, v := range map[string]interface{}{
		"id":       int64(123),
		"username": "admin",
		"password": "admin",
		"name":     "admin_001",
		"gender":   1,
		"email":    "xxx@163.com",
	} {
		container.AsRegister().RegisterInstance(v, k)
	}

	child := container.Child()

	var info *User
	var dependFactory types.DependencyFactory
	var builderFactory types.BuilderFactory
	child.Assign(&dependFactory)
	child.Assign(&builderFactory)

	depend, _ := dependFactory.Instance(info)
	builder, _ := builderFactory.Instance(factory.NewTypeFactory(info), depend)

	child.AsRegister().RegisterFactory(builder.AsFactory(), &info, "admin")

	child.Assign(&info, "admin")
	fmt.Println(info)
}
