package authhandlers

import (
	"snipetz/commons/schema"

	"github.com/gin-gonic/gin"
)

func HandleLogin(c *gin.Context) {
	var data schema.AuthLoginForm
	err := c.ShouldBindBodyWithJSON(&data)
	if err != nil || !data.AllFieldValid() {
		c.JSON(400, schema.AuthLoginResponse{
			Status:        "invalid",
			InvalidReason: "Bad fields",
		})
		return
	}

}
