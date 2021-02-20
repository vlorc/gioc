package main

import (
	"crypto/tls"
	. "github.com/vlorc/gioc/module"
	. "github.com/vlorc/gioc/module/operation"
	"github.com/vlorc/gioc/types"
	"net"
)

var TLSModule = NewModuleFactory(
	Condition(
		And(
			HavingFile(types.StringType, "certFile"),
			HavingFile(types.StringType, "keyFile"),
		),
		Declare(
			Method(func(param struct {
				config   *tls.Config `inject:"'default' optional"`
				certFile string
				keyFile  string
			}) (cfg *tls.Config, err error) {
				if "" == param.certFile || "" == param.keyFile {
					return
				}
				if nil != param.config {
					cfg = param.config.Clone()
				} else {
					cfg = &tls.Config{}
				}
				cert, err := tls.LoadX509KeyPair(param.certFile, param.keyFile)
				if nil == err {
					cfg.Certificates = append(cfg.Certificates, cert)
				}
				return
			}), Singleton(),
		),
		Primary(
			Method2(func(param struct {
				listen net.Listener `inject:"''"`
				config *tls.Config  `inject:"''"`
			}) net.Listener {
				if nil != param.config {
					return tls.NewListener(param.listen, param.config)
				}
				return param.listen
			}),
		),
	),
)
