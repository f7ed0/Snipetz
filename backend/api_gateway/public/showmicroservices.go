package public

import (
	"snipetz/api_gateway/microservices/registery"

	"github.com/gin-gonic/gin"
)

func ShowMicroservices(c *gin.Context) {
	c.JSON(200, registery.GetMicroservicesRegistery())
}
