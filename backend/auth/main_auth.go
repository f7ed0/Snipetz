package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"
	"snipetz/auth/dbconnection"
	authhandlers "snipetz/auth/handlers"
	"snipetz/commons/handlers"
	"snipetz/commons/schema"
	"snipetz/commons/utils"

	"github.com/f7ed0/golog/lg"
	"github.com/gin-gonic/gin"
)

func main() {
	lg.Init(lg.ALL, true)
	lg.Info.Println("Starting authentication microservice")
	ips, err := utils.GetIPs()
	if err != nil {
		lg.Error.Fatalln(err.Error())
	}

	lg.Info.Println("Connecting to mongoDB...")

	err = dbconnection.Cntr.Init(os.Getenv("db_url"))

	if err != nil {
		lg.Error.Fatalf(err.Error())
	}

	data, err := json.Marshal(&schema.ConnectionRequest{MicroserviceType: "auth", URI: "http://" + ips["eth0"][0].String()})
	if err != nil {
		lg.Error.Fatalln(err.Error())
	}
	resp, err := http.Post("http://api/microservices/connect", "json", bytes.NewBuffer(data))
	if err != nil {
		lg.Error.Fatalln(err.Error())
	}
	if resp.StatusCode == http.StatusNoContent {
		lg.Info.Println("Accepted by the gateway.")
	} else {
		lg.Error.Fatalln("Rejected by the gateway. Status", resp.StatusCode)
	}

	g := gin.Default()
	g.Handle("GET", "/heartbeat", handlers.Heartbeathandler)
	g.Handle("POST", "/transaction/undo", handlers.TransactionRevert)
	g.Handle("POST", "/transaction/close", handlers.TransactionClose)
	g.Handle("POST", "/register", authhandlers.Register)
	g.Run(":80")
}
