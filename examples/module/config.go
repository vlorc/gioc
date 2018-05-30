package main

import (
	. "github.com/vlorc/gioc/module"
	. "github.com/vlorc/gioc/module/operation"
)

var ConfigModule = NewModuleFactory(
	Export(
		Mapping(map[string]interface{}{
			"port": 80,
			"host": "127.0.0.1",
		}),
	),
)
