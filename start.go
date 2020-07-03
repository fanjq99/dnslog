package dnslog

import (
	"github.com/fanjq99/dnslog/api"
	"github.com/fanjq99/dnslog/config"
	"github.com/fanjq99/dnslog/dns"
)

func Run(setting config.YmlConfig, daemon bool) {
	go api.NewHttpServer(setting).Run()

	if daemon {
		go dns.NewUDPServer(dns.NewHandler(setting)).Run()
	} else {
		dns.NewUDPServer(dns.NewHandler(setting)).Run()
	}
}
