package main

import (
	"crypto/tls"
	. "github.com/vlorc/gioc/module"
	. "github.com/vlorc/gioc/module/operation"
	"net"
)

var TlsModule = NewModuleFactory(
	Declare(
		Method(func(param struct {
			config   *tls.Config `inject:"'default' optional"`
			certFile string      `inject:"optional"`
			keyFile  string      `inject:"optional"`
		}) (cfg *tls.Config, err error) {
			if "" == param.certFile || "" == param.keyFile {
				return
			}
			if nil != param.config {
				cfg = param.config.Clone()
			} else {
				cfg = &tls.Config{}
			}
			cfg.Certificates = make([]tls.Certificate, 1)
			cfg.Certificates[0], err = tls.LoadX509KeyPair(param.certFile, param.keyFile)
			return
		}), Singleton(),
	),
	Export(
		Method(func(param struct {
			listen net.Listener `inject:"''"`
			config *tls.Config  `inject:"''"`
		}) net.Listener {
			if nil != param.config {
				return tls.NewListener(param.listen, param.config)
			}
			return param.listen
		}), Singleton(),
	),
)
