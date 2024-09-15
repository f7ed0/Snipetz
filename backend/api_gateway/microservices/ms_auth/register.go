package msauth

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	snipetzerror "snipetz/commons/errors"
	"snipetz/commons/schema"

	"github.com/f7ed0/golog/lg"
)

func (m AuthMicroservice) RegisterUser(register_schema schema.AuthRegisterForm) (response schema.AuthRegisterResponse, err error) {
	response = schema.AuthRegisterResponse{}
	b, err := json.Marshal(register_schema)
	if err != nil {
		return
	}
	resp, err := http.Post(m.Uri+"/register", "json", bytes.NewBuffer(b))
	if err != nil {
		return
	}
	if resp.StatusCode != http.StatusOK {
		err = snipetzerror.ErrorRefusedByMicroservice
		return
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		lg.Error.Println(err.Error())
		return
	}
	json.Unmarshal(body, &response)
	return
}
