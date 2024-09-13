package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"snipetz/commons/handlers"
	common_schema "snipetz/commons/schema"
	"snipetz/commons/util"

	"github.com/f7ed0/golog/lg"
	"github.com/gin-gonic/gin"
)

func main() {
	lg.Init(lg.ALL, true)
	lg.Info.Println("Starting authentication microservice")
	ips, err := util.GetIPs()
	if err != nil {
		lg.Error.Fatalln(err.Error())
	}

	data, err := json.Marshal(&common_schema.ConnectionRequest{MicroserviceType: "auth", URI: "http://" + ips["eth0"][0].String()})
	lg.Debug.Println(string(data))
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
	g.Run(":80")
}
