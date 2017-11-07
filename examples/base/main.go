package main

import (
	"fmt"
	"github.com/vlorc/gioc"
	"github.com/vlorc/gioc/types"
	"math/rand"
	"time"
)

var source = rand.New(rand.NewSource(time.Now().UnixNano()))

func testInstance(register types.Register, provider types.Provider, name string) {
	src := source.Int63()
	dst := source.Int63()

	register.RegisterInstance(src, name)
	provider.Assign(&dst, name)

	fmt.Printf("[%s] src:%v dst:%v equal:%v\n", name, src, dst, src == dst)
}

func main() {
	container := gioc.NewRootContainer()
	key := []string{"age", "gender", "high", "width"}

	for _, v := range key {
		testInstance(container.AsRegister(), container.AsProvider(), v)
	}

	child := container.Child()
	for _, v := range key {
		var value int64
		child.AsProvider().Assign(&value, v)
		fmt.Println(v, ":", value)
	}

}
