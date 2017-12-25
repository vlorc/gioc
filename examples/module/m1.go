package main

import (
	. "github.com/vlorc/gioc/module"
	. "github.com/vlorc/gioc/module/operation"
)

var Module1 = NewModuleFactory(
	Export(
		Mapping(map[string]interface{}{
			"age":11,
			"gender":2,
		}),
		Instance("xxx@163.com"), Id("email"),
		Instance("admin_001"), Id("name"),
	),
)
