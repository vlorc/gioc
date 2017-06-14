package main

import (
	"fmt"
	"github.com/vlorc/gioc"
	"github.com/vlorc/gioc/factory"
	"github.com/vlorc/gioc/types"
	"github.com/vlorc/gioc/utils"
	"reflect"
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
	Id        int64 `inject:"'id'"`
	*Identity `inject:"extends"`
	*Personal `inject:"extends"`
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

	register(child, (*Identity)(nil))
	register(child, (*Personal)(nil))
	register(child, getUser)

	fmt.Println(child.Resolve((**User)(nil)))
}

func getUser(identity *Identity, personal *Personal) (*User, error, int64) {
	return &User{1, identity, personal}, nil, 1
}

func register(container types.Container, impType interface{}, name ...string) types.Dependency {
	var dependFactory types.DependencyFactory
	var builderFactory types.BuilderFactory
	container.Assign(&dependFactory)
	container.Assign(&builderFactory)

	typ := utils.TypeOf(impType)
	depend, _ := dependFactory.Instance(typ)

	if reflect.Func == depend.Type().Kind() {
		builder, _ := builderFactory.Instance(factory.NewParamFactory(depend.Length()), depend)
		container.AsRegister().RegisterMethod(builder.AsFactory(), impType, nil)
	} else {
		builder, _ := builderFactory.Instance(factory.NewTypeFactory(typ), depend)
		container.AsRegister().RegisterFactory(builder.AsFactory(), reflect.PtrTo(depend.Type()), name...)
	}

	return depend
}
