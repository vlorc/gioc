package main

import (
	. "github.com/vlorc/gioc/module"
	. "github.com/vlorc/gioc/module/operation"
)

var Module1 = NewModuleFactory(
	Declare(
		Instance(11), Name("age"), Export(),
	),
	Declare(
		Instance(2), Name("gender"), Export(),
	),
	Declare(
		Instance("xxx@163.com"), Name("email"), Export(),
	),
	Declare(
		Instance("admin_001"), Name("name"), Export(),
	),
)
