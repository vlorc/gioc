package main

import (
	. "github.com/vlorc/gioc"
	. "github.com/vlorc/gioc/module/operation"
	"net"
	"net/http"
)

func main() {
	NewRootModule(
		Import(
			ConfigModule,
			ListenModule,
			TLSModule,
			ServerModule,
		),
		Bootstrap(func(server *http.Server, listen net.Listener) {
			server.Serve(listen)
		}),
	)
}
