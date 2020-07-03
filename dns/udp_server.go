package dns

import (
	"time"

	"github.com/fanjq99/common/log"
	"github.com/miekg/dns"
)

type UDPServer struct {
	dnsServer *dns.Server
}

func NewUDPServer(handler *Handler) *UDPServer {
	udpHandler := dns.NewServeMux()
	udpHandler.HandleFunc(".", handler.DoUDP)
	dnsServer := &dns.Server{
		Net:          "udp",
		Addr:         "0.0.0.0:53",
		Handler:      udpHandler,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}
	return &UDPServer{
		dnsServer: dnsServer,
	}
}

func (u *UDPServer) Run() {
	log.Info("dns server listen in", u.dnsServer.Net, u.dnsServer.Addr)
	err := u.dnsServer.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
	defer u.dnsServer.Shutdown()
}
