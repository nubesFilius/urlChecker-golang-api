package url

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nubesFilius/urlChecker-golang-api/services"
)

const (
	health = "OK"
)

var (
	UrlController urlControllerInterface = &urlController{}
)

type urlControllerInterface interface {
	Check(c *gin.Context)
}

type urlController struct{}

func (uc *urlController) Check(c *gin.Context) {
	//TODO: Need to be able to grab entire url rather than just the first one.
	domain := c.Param("url")
	if !services.UrlService.MalwareExists(domain) {
		c.JSON(http.StatusOK, fmt.Sprintf("Domain '%v' not in the list", domain))
	} else {
		c.JSON(http.StatusForbidden, fmt.Sprintf("Domain '%v' contains malware", domain))
	}
}
