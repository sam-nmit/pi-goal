package rest

import (
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
)

func (server *RESTServer) getHosts(c *gin.Context) {

	rgx, err := regexp.Compile(c.Param("rx"))
	if err != nil {
		c.Error(err)
		return
	}
	rMap := make(map[string]string)

	for k, v := range server.dns.Resolver.Hosts() {
		if rgx.MatchString(k) {
			rMap[k] = v
		}
	}

	c.JSON(http.StatusOK, struct {
		Hosts map[string]string
	}{
		rMap,
	})
}
