package healthmanagement

import (
	"net/http"
	"snipetz/api_gateway/microservices/registery"
	"snipetz/commons/schema"

	"github.com/f7ed0/golog/lg"
	"github.com/gin-gonic/gin"
)

func Connect(c *gin.Context) {
	var crq schema.ConnectionRequest
	// Cheking body for a connection request
	err := c.ShouldBindBodyWithJSON(&crq)
	if err != nil {
		lg.Error.Println(err.Error())
		c.Status(400)
		return
	}

	// Checking the content for a connection request
	if crq.MicroserviceType == "" || crq.URI == "" {
		c.Status(400)
		return
	}
	if crq.URI == "" {
		c.Status(400)
		return
	}

	registery.AddMicroservice(crq.MicroserviceType, crq.URI)

	c.Status(http.StatusNoContent)

}
