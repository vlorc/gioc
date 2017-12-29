package main

import (
	. "github.com/vlorc/gioc"
	. "github.com/vlorc/gioc/module/operation"
	"net/http"
	"net"
)

func main() {
	NewRootModule(
		Import(
			ConfigModule,
			ListenModule,
			ServerModule,
		),
		Bootstrap(func(param struct{ server *http.Server `inject:"''"`; listen net.Listener `inject:"''"`}) {
			param.server.Serve(param.listen)
		}),
	)
}
