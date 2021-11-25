package health

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	health = "OK"
)

var (
	HealthController healthControllerInterface = &healthController{}
)

type healthControllerInterface interface {
	Check(c *gin.Context)
}

type healthController struct{}

func (cont *healthController) Check(c *gin.Context) {
	c.String(http.StatusOK, health)
}
