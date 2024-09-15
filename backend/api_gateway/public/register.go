package public

import (
	"net/http"
	"snipetz/api_gateway/microservices"
	msauth "snipetz/api_gateway/microservices/ms_auth"
	"snipetz/api_gateway/models"

	"github.com/f7ed0/golog/lg"
	"github.com/gin-gonic/gin"
)

func RegisterHandler(c *gin.Context) {
	var registerInformation models.RegisterForm

	err := c.ShouldBindBodyWithJSON(&registerInformation)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		lg.Error.Println(err.Error())
		return
	}

	auth, err := msauth.GetAuthMicroservice()
	if err != nil {
		c.Status(http.StatusInternalServerError)
		lg.Error.Println(err.Error())
		return
	}

	rsp, err := auth.RegisterUser(registerInformation.AuthRegisterForm)
	if rsp.Status != "valid" {
		lg.Warn.Println("Registering refused by auth :", rsp.InvalidReason)
	}

	// TODO close Transaction
	microservices.TransactionClose(rsp.TransactionId)

	c.Status(200)
}
