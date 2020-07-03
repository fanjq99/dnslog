package dns

import (
	"net"
	"strings"
	"time"

	"github.com/fanjq99/common/db"
	"github.com/fanjq99/common/log"
	"github.com/fanjq99/dnslog/config"
	"github.com/go-redis/redis"
	"github.com/miekg/dns"
)

type Handler struct {
	setting     config.YmlConfig
	redisClient *redis.Client
	dnsDomain   string
}

func (h *Handler) DoTCP(w dns.ResponseWriter, req *dns.Msg) {
	h.do("tcp", w, req)
}

func (h *Handler) DoUDP(w dns.ResponseWriter, req *dns.Msg) {
	h.do("udp", w, req)
}

func NewHandler(setting config.YmlConfig) *Handler {
	c, err := db.GetRedisClient(setting.Redis.Addr,setting.Redis.Password,setting.Redis.Database)
	if err != nil {
		log.Fatal(err)
	}
	dnsDomain := setting.DnsDomain
	if !strings.HasSuffix(dnsDomain, ".") {
		dnsDomain = dnsDomain + "."
	}
	return &Handler{
		setting:     setting,
		redisClient: c,
		dnsDomain:   dnsDomain,
	}
}

func (h *Handler) do(netType string, w dns.ResponseWriter, req *dns.Msg) {
	q := req.Question[0]

	var remoteIp net.IP
	if netType == "tcp" {
		remoteIp = w.RemoteAddr().(*net.TCPAddr).IP
	} else {
		remoteIp = w.RemoteAddr().(*net.UDPAddr).IP
	}

	responseIp := "127.0.0.1"

	m := new(dns.Msg)
	m.SetReply(req)
	ttl := 600
	queryHost := q.Name
	var rmd string

	if strings.HasSuffix(queryHost, h.dnsDomain) {
		responseIp = h.setting.ServerIp
		rmd = strings.Split(queryHost, ".")[0]
	}
	log.Info("dns", remoteIp, q.Name, rmd)

	switch q.Qtype {
	case dns.TypeA:
		rrHeader := dns.RR_Header{
			Name:   q.Name,
			Rrtype: dns.TypeA,
			Class:  dns.ClassINET,
			Ttl:    uint32(ttl),
		}

		a := &dns.A{
			Hdr: rrHeader,
			A:   net.ParseIP(responseIp).To4(),
		}
		m.Answer = append(m.Answer, a)
		_ = w.WriteMsg(m)
		if rmd != "" {
			err := h.redisClient.Set(rmd, remoteIp.String(), time.Duration(h.setting.SaveTime)*time.Second).Err()
			if err != nil {
				log.Error(err)
			}
		}
	case dns.TypeAAAA:
		rrHeader := dns.RR_Header{
			Name:   q.Name,
			Rrtype: dns.TypeAAAA,
			Class:  dns.ClassINET,
			Ttl:    uint32(ttl),
		}

		aaa := &dns.AAAA{
			Hdr:  rrHeader,
			AAAA: net.ParseIP(responseIp).To16(),
		}
		m.Answer = append(m.Answer, aaa)
		_ = w.WriteMsg(m)

		if rmd != "" {
			err := h.redisClient.Set(rmd, remoteIp.String(), time.Duration(h.setting.SaveTime)*time.Second)
			if err != nil {
				log.Error(err)
			}
		}
	default:
		log.Info("not support dns type", q.Qtype, remoteIp)
		_ = w.WriteMsg(m)
	}
}
