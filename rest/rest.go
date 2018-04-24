package rest

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/sam-nmit/pi-goal/pidns"
)

type RESTServer struct {
	server  *gin.Engine
	Addr    string
	dns     *pidns.DNSServer
	Running bool
}

func NewServer(dns *pidns.DNSServer, listenAddr string) *RESTServer {
	gin.SetMode(gin.ReleaseMode)

	j := gin.Default()

	server := &RESTServer{
		Addr:   listenAddr,
		dns:    dns,
		server: j,
	}

	j.GET("/hosts/:rx", server.getHosts)

	return server
}

func (s *RESTServer) Start() {
	go func() {
		log.Printf("RestServer] Starting rest interface on %s...\n", s.Addr)
		err := s.server.Run(s.Addr)
		log.Println("RestServer] Rest interface has exited")
		log.Println(err)
	}()
}
