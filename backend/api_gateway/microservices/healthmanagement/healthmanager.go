package healthmanagement

import (
	"snipetz/api_gateway/microservices/registery"
	"time"
)

func HealthManager() {
	for {
		time.Sleep(10 * time.Second)
		reg := registery.GetMicroservicesRegistery()
		for _, v := range reg {
			for _, i := range v {
				if time.Now().Unix()-i.LastSign > 20 {
					if i.IsAlive() {
						registery.UpdateHeartBeat(i.Id)
					} else {
						registery.RemoveMicroservice(i.Id)
					}
				}
			}
		}
	}
}
