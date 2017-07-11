package main

import (
	"fmt"
	"github.com/vlorc/gioc"
	"github.com/vlorc/gioc/factory"
	"github.com/vlorc/gioc/types"
)

type Personal struct {
	Name   string `inject:"'name'"`
	Age    *int    `inject:"'age' default"`
	Gender int    `inject:"'gender' optional"`
	Email  string `inject:"'email' optional"`
}

type User struct {
	Id       int64 `inject:"lower"`
	Personal func()****Personal `inject:"lazy extends"`
}

func main() {
	container := gioc.NewRootContainer()
	for k, v := range map[string]interface{}{
		"id":       int64(123),
		"name":     "admin",
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

	depend, err := dependFactory.Instance(info)
	if nil != err {
		panic(err)
	}
	builder, err := builderFactory.Instance(factory.NewTypeFactory(info), depend)
	if nil != err {
		panic(err)
	}
	child.AsRegister().RegisterFactory(builder.AsFactory(), &info, "")
	if err = child.Assign(&info, ""); nil != err {
		panic(err)
	}

	fmt.Println(****info.Personal())
}
