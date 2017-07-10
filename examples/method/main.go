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
	Software struct {
		Version string
	} `inject:"extends"`
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
		"Version":  "1.0.1",
	} {
		container.AsRegister().RegisterInstance(v, k)
	}

	child := container.Child()
	register(child, (*Identity)(nil))
	register(child, (*Personal)(nil))
	register(child, getUser)

	fmt.Println(child.Resolve((**User)(nil)))
}

func getUser(param struct {
	Name string `inject:"lower"`
}, identity *Identity, personal *Personal) (*User, error, int64) {
	fmt.Println("getUser by name:", param.Name, "version:", identity.Software.Version)
	return &User{1, identity, personal}, nil, 1
}

func register(container types.Container, impType interface{}, name ...string) types.Dependency {
	var dependFactory types.DependencyFactory
	var builderFactory types.BuilderFactory
	var builder types.Builder
	container.Assign(&dependFactory)
	container.Assign(&builderFactory)

	typ := utils.TypeOf(impType)
	depend, err := dependFactory.Instance(typ)
	if nil != err {
		panic(err)
	}

	if reflect.Func == depend.Type().Kind() {
		builder, err = builderFactory.Instance(factory.NewParamFactory(depend.Length()), depend)
		container.AsRegister().RegisterMethod(builder.AsFactory(), impType, nil)
	} else {
		builder, err = builderFactory.Instance(factory.NewTypeFactory(typ), depend)
		container.AsRegister().RegisterFactory(builder.AsFactory(), reflect.PtrTo(depend.Type()), name...)
	}
	if nil != err {
		panic(err)
	}

	return depend
}
