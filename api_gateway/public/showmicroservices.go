package public

import (
	"snipetz/api_gateway/microservices"

	"github.com/gin-gonic/gin"
)

func ShowMicroservices(c *gin.Context) {
	c.JSON(200, microservices.GetMicroservicesRegistery())
}
