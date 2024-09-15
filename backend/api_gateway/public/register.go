package public

import (
	"errors"
	"net/http"
	msauth "snipetz/api_gateway/microservices/ms_auth"
	"snipetz/api_gateway/models"
	snipetzerror "snipetz/commons/errors"

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
	if errors.Is(err, snipetzerror.ErrorRefusedByMicroservice) {
		c.JSON(400, map[string]string{
			"reason": "bad fields",
		})
		return
	} else if err != nil {
		c.Status(500)
		return
	}
	if rsp.Status != "valid" {
		lg.Warn.Println("Registering refused by auth :", rsp.InvalidReason)
		c.JSON(400, map[string]string{
			"reason": rsp.InvalidReason,
		})
		return
	}

	// TODO close Transaction
	err = auth.Microservice.TransactionClose(rsp.TransactionId)
	if err != nil {
		lg.Error.Println("error closing auth transaction :", err.Error())
	}

	c.Status(200)
}
