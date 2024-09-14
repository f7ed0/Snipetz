package msauth

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"snipetz/api_gateway/microservices/registery"
	"snipetz/commons/schema"

	"github.com/f7ed0/golog/lg"
)

func RegisterUser(register_schema schema.AuthRegisterForm) (response schema.AuthRegisterResponse, err error) {
	response = schema.AuthRegisterResponse{}
	b, err := json.Marshal(register_schema)
	if err != nil {
		return
	}
	addr, err := registery.GetMicroserviceAddress("auth")
	if err != nil {
		return
	}
	resp, err := http.Post(addr, "json", bytes.NewBuffer(b))
	if err != nil {
		return
	}
	if resp.StatusCode != http.StatusOK {
		err = errors.New("Microservice refused the request")
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		lg.Error.Println(err.Error())
		return
	}
	json.Unmarshal(body, &response)
	return
}
