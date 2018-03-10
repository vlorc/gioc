package main

import (
	. "github.com/vlorc/gioc/module"
	. "github.com/vlorc/gioc/module/operation"
	"net"
	"strconv"
)

var ListenModule = NewModuleFactory(
	Export(
		Method(func(param struct {
			port int    `inject:"default(8080)"`
			host string `inject:"default('0.0.0.0')"`
		}) (net.Listener, error) {
			return net.Listen("tcp", net.JoinHostPort(param.host, strconv.Itoa(param.port)))
		}), Singleton(),
	),
)
