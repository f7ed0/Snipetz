package main

import (
	"net/http"
	"snipetz/api_gateway/microservices/healthmanagement"
	"snipetz/api_gateway/public"
	"snipetz/commons/utils"
	"time"

	"github.com/f7ed0/golog/lg"
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
)

func main() {
	lg.Init(lg.ALL, true)
	lg.Info.Println("Starting api gateway...")

	ips, err := utils.GetIPs()

	if err != nil {
		lg.Error.Fatalln(err.Error())
	}

	lg.Info.Printf("\n---------------------\npublic interface ip : %v\nmicroservice interface ip : %v\n---------------------", ips["eth0"][0], ips["eth1"][0])

	microservice_side := &http.Server{
		Addr:         ips["eth1"][0].String() + ":80",
		Handler:      microserviceHandler(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	api_side := &http.Server{
		Addr:         ips["eth0"][0].String() + ":80",
		Handler:      apiHandler(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	var g errgroup.Group

	lg.Info.Println("Starting microservice http server")
	g.Go(func() error {
		return microservice_side.ListenAndServe()
	})

	lg.Info.Println("Starting public http server")
	g.Go(func() error {
		return api_side.ListenAndServe()
	})

	go healthmanagement.HealthManager()

	if err := g.Wait(); err != nil {
		lg.Error.Fatalln(err.Error())
	}
}

func microserviceHandler() http.Handler {
	e := gin.Default()
	e.POST("/microservices/connect", healthmanagement.Connect)
	return e
}

func apiHandler() http.Handler {
	e := gin.Default()
	e.GET("/ping", func(c *gin.Context) {
		c.JSON(200, "pong")
	})
	e.GET("/microservices/list", public.ShowMicroservices)
	e.POST("/register", public.RegisterHandler)
	return e
}
