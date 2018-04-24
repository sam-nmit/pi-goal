package pidns

import (
	"net"

	"github.com/miekg/dns"
)

func (ds *DNSServer) ServeDNS(w dns.ResponseWriter, r *dns.Msg) {
	msg := dns.Msg{}
	msg.SetReply(r)
	switch r.Question[0].Qtype {
	case dns.TypeA:
		msg.Authoritative = true

		domain := msg.Question[0].Name

		address := ds.Resolver.Resolve(domain)
		// log.Println(domain, address)

		if address != nil {
			for _, a := range *address {
				msg.Answer = append(msg.Answer, &dns.A{
					Hdr: dns.RR_Header{Name: domain, Rrtype: dns.TypeA, Class: dns.ClassINET, Ttl: 60},
					A:   net.ParseIP(a),
				})
			}
		}

	}
	w.WriteMsg(&msg)
}
