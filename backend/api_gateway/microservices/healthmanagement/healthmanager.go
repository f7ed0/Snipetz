package healthmanagement

import (
	"net/http"

	"snipetz/api_gateway/microservices/registery"
	"time"

	"github.com/f7ed0/golog/lg"
)

func HealthManager() {
	for {
		time.Sleep(10 * time.Second)
		reg := registery.GetMicroservicesRegistery()
		for _, v := range reg {
			for _, i := range v {
				if time.Now().Unix()-i.LastSign > 20 {
					lg.Info.Println("sending heartbeat to", i)

					resp, err := http.Get(i.Uri + "/heartbeat")
					if err != nil {
						lg.Error.Fatalln("Error in healthmanager : ", err.Error())
					}
					if resp.StatusCode == http.StatusNoContent {
						registery.UpdateHeartBeat(i.Id)
					} else {
						registery.RemoveMicroservice(i.Id)
					}
				}
			}
		}
	}
}
