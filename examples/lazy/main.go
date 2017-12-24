package main

import (
	"fmt"
	"github.com/vlorc/gioc"
	"github.com/vlorc/gioc/factory"
	"github.com/vlorc/gioc/types"
)

type Personal struct {
	name   string `inject:"'name'"`
	age    *int   `inject:"'age' default"`
	gender int    `inject:"'gender' optional"`
	email  string `inject:"'email' optional"`
}

type User struct {
	id       int64               `inject:"lower"`
	personal func() ****Personal `inject:"lazy extends"`
}

func main() {
	container := gioc.NewRootContainer()
	for k, v := range map[string]interface{}{
		"id":     int64(123),
		"name":   "admin",
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
	child.AsRegister().RegisterFactory(builder.AsFactory(), &info, "")
	if err = child.AsProvider().Assign(&info, ""); nil != err {
		panic(err)
	}

	fmt.Println(****info.personal())
}
