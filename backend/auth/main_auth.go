package main

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"os"
	"snipetz/auth/schema"
	"snipetz/commons/handlers"
	common_schema "snipetz/commons/schema"
	"snipetz/commons/util"

	"github.com/f7ed0/golog/lg"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	lg.Init(lg.ALL, true)
	lg.Info.Println("Starting authentication microservice")
	ips, err := util.GetIPs()
	if err != nil {
		lg.Error.Fatalln(err.Error())
	}

	lg.Info.Println("Connecting to mongoDB...")
	cli, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(os.Getenv("db_url")))
	if err != nil {
		lg.Info.Fatalln("Cant access db :", err.Error())
	}
	db := cli.Database("auth")
	cur, err := db.Collection("users").Find(context.TODO(), bson.D{})
	if err != nil {
		lg.Info.Fatalln("Can access db :", err.Error())
	}
	var us []schema.User
	cur.Decode(&us)
	lg.Verbose.Println(us)

	data, err := json.Marshal(&common_schema.ConnectionRequest{MicroserviceType: "auth", URI: "http://" + ips["eth0"][0].String()})
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
