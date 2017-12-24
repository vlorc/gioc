package main

import (
	"fmt"
	. "github.com/vlorc/gioc"
	. "github.com/vlorc/gioc/module/operation"
)

func main() {
	NewRootModule(
		Import(
			Module1,
			Module2,
		),
		Bootstrap(func(param struct {
			id     int64
			age    int
			gender int
			email  string
			name   string
		}) {
			fmt.Println(param)
		}, func(param struct{ id int64 }) {
			fmt.Println("next id: ", param.id)
		}),
	)
}
