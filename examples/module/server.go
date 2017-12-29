package main

import (
	. "github.com/vlorc/gioc/module"
	. "github.com/vlorc/gioc/module/operation"
	"net/http"
)

var ServerModule = NewModuleFactory(
	Export(
		Method(func(param struct{ Handle http.Handler `inject:"'' optional"`}) *http.Server{
			return &http.Server{Handler: param.Handle}
		}),Singleton(),
	),
)
