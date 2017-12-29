package main

import (
	"net"
	. "github.com/vlorc/gioc/module"
	. "github.com/vlorc/gioc/module/operation"
	"strconv"
)

var ListenModule = NewModuleFactory(
	Declare(
		Method(func(param struct{port int `inject:"default(8080)"`; host string `inject:"default('0.0.0.0')"`}) string{
			return net.JoinHostPort(param.host,strconv.Itoa(param.port))
		}),Id("addr"),
	),
	Export(
		Method(func(param struct{ addr string }) (net.Listener,error){
			return net.Listen("tcp",param.addr)
		}),Singleton(),
	),
)
