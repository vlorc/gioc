package main

import (
	"fmt"
	"github.com/vlorc/gioc"
	"github.com/vlorc/gioc/factory"
	"github.com/vlorc/gioc/types"
)

type User struct {
	Id       int64 `inject:"lower"`
	Personal ****struct {
		Name   string `inject:"lower"`
		Age    int    `inject:"id('age') default(99)"`
		Gender int    `inject:"lower optional"`
		Email  string `inject:"lower optional"`
	} `inject:"extends"`
}

func main() {
	container := gioc.NewRootContainer()
	for k, v := range map[string]interface{}{
		"id":     int64(123),
		"name":   "admin_001",
		"gender": 1,
		"email":  "xxx@163.com",
	} {
		container.AsRegister().RegisterInstance(v, k)
	}

	child := container.NewChild()
	var info *User
	var dependFactory types.DependencyFactory
	var builderFactory types.BuilderFactory
	child.AsProvider().Assign(&dependFactory)
	child.AsProvider().Assign(&builderFactory)

	depend, err := dependFactory.Instance(info)
	if nil != err {
		panic(err)
	}
	builder, err := builderFactory.Instance(factory.NewTypeFactory(info), depend)
	if nil != err {
		panic(err)
	}
	child.AsRegister().RegisterFactory(builder.AsFactory(), &info, "admin")
	if err = child.AsProvider().Assign(&info, "admin"); nil != err {
		panic(err)
	}

	fmt.Println(info)
	fmt.Println(****info.Personal)
}
