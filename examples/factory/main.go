package main

import (
	"fmt"
	"github.com/vlorc/gioc"
	"github.com/vlorc/gioc/factory"
	"github.com/vlorc/gioc/types"
)

func main() {
	container := gioc.NewRootContainer()
	age := 17

	// register an int type value factory,this is similar to RegisterInstance
	container.AsRegister().RegisterFactory(factory.NewValueFactory(age), (*int)(nil), "age")

	// create a custom func factory
	inc := factory.NewFuncFactory(func(types.Provider) (interface{}, error) {
		age++
		return age, nil
	})

	// register an int type
	container.AsRegister().RegisterFactory(inc, &age, "inc")

	// convert custom factory into singleton mode factory
	container.AsRegister().RegisterFactory(factory.NewSingleFactory(inc), &age, "once")

	// get an instance type int and name age
	fmt.Println(container.AsProvider().Resolve((*int)(nil), "age"))

	// same as above,this value add 1 every times
	fmt.Println(container.AsProvider().Resolve((*int)(nil), "inc"))

	// same as above,but only once
	fmt.Println(container.AsProvider().Resolve((*int)(nil), "once"))
}
