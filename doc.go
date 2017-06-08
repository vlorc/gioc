// package gioc is a lightweight Ioc framework,it provides register and factory and depend solution

/*
Package gioc provides methods for generate a root container. Typically your application
will have a single, minimal file in its main package similar to:

	package main

	import "github.com/vlorc/gioc"

	func main() {
		container := gioc.NewRootContainer()
	}


*/
package gioc
