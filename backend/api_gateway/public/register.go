package public

import (
	"net/http"
	msauth "snipetz/api_gateway/microservices/ms_auth"
	"snipetz/api_gateway/models"

	"github.com/f7ed0/golog/lg"
	"github.com/gin-gonic/gin"
)

func RegisterHandler(c *gin.Context) {
	var registerInformation models.RegisterForm

	err := c.ShouldBindBodyWithJSON(&registerInformation)

	rsp, err := msauth.RegisterUser(registerInformation.AuthRegisterForm)

	if err != nil {
		c.Status(http.StatusInternalServerError)
		lg.Error.Println(err.Error())
		return
	}
	// TODO
	_ = rsp
	c.Status(200)
}
