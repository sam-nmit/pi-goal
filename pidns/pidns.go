package pidns

import (
	"regexp"

	"github.com/miekg/dns"
	"github.com/sam-nmit/pi-goal/pidns/resolver"
	"github.com/sam-nmit/pi-goal/utils"

	"bufio"
	"log"
	"os"
	"path/filepath"
)

type DNSServer struct {
	dns      *dns.Server
	Resolver resolver.INameResolver
}

func DefaultServer(address string) *DNSServer {

	server := &DNSServer{
		Resolver: resolver.NewInMemoryResolver(),
	}

	dns := &dns.Server{
		Addr:    address,
		Net:     "udp",
		Handler: server,
	}

	server.dns = dns

	return server
}

func (server *DNSServer) Start() error {
	log.Println("Statring on ", server.dns.Addr)
	return server.dns.ListenAndServe()
}

func (server *DNSServer) AddRule(rule string, rx *regexp.Regexp, nameIndex, hostIndex int) {

	results := rx.FindStringSubmatch(rule)
	log.Println(rule)
	log.Println(results)
	server.Resolver.AddRule(results[nameIndex], results[hostIndex])

}

func (server *DNSServer) LoadRules(dir string, rules []utils.RuleEntry) {

	for _, r := range rules {

		if r.IsRawList() {
			for _, rawrule := range r.Sources {
				name, addr := r.Parse(rawrule)
				server.Resolver.AddRule(name, addr)
			}
		} else {

			for _, rFile := range r.Sources {

				path := filepath.Join(dir, rFile)
				if f, err := os.Open(path); err != nil {
					log.Printf("Failed to open config file %s\n", rFile)
				} else {
					defer f.Close()
					log.Printf("Loading [%s] %s\n", r.Format, rFile)
					lineReader := bufio.NewScanner(f)

					for lineReader.Scan() {
						line := lineReader.Text()
						if len(line) > 0 && line[0] != '#' {
							name, addr := r.Parse(line)
							server.Resolver.AddRule(name, addr)
						}
					}

				}
			}

		}
	}

}
