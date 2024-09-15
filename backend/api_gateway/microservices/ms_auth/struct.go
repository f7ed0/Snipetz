package msauth

import "snipetz/api_gateway/microservices/registery"

type AuthMicroservice struct {
	registery.Microservice
}

func GetAuthMicroservice() (auth AuthMicroservice, err error) {
	auth.Microservice, err = registery.GetMicroservice("auth")
	return
}
