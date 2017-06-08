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

	//the same as RegisterInstance
	//regiter int of value factory
	container.AsRegister().RegisterFactory(
		factory.NewValueFactory(age),
		&age, /*(*int)(nil)*/
		"age",
	)

	//regiter int of func factory
	inc := factory.NewFuncFactory(func(types.Provider) (interface{}, error) {
		age++
		return age, nil
	})

	container.AsRegister().RegisterFactory(
		inc,
		&age,
		"inc",
	)

	// singleton mode,convert the singleton factory
	once := factory.NewSingleFactory(inc)
	container.AsRegister().RegisterFactory(
		once,
		&age,
		"once",
	)

	// get a int type of 'age'
	fmt.Println(container.Resolve((*int)(nil), "age"))

	// get a age + 1
	fmt.Println(container.Resolve((*int)(nil), "inc"))

	// get a age,once + 1
	fmt.Println(container.Resolve((*int)(nil), "once"))
}
