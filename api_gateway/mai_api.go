package main

import (
	"snipetz/api_gateway/microservices"

	"github.com/f7ed0/golog/lg"
	"github.com/gin-gonic/gin"
)

func main() {
	lg.Init(lg.ALL, true)
	lg.Info.Println("Starting api gateway...")

	r := gin.Default()
	r.GET("/microservices/connect", microservices.Connect)
	r.Run()
}
