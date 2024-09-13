package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Heartbeathandler(c *gin.Context) {
	c.Status(http.StatusNoContent)
}
