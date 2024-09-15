package registery

import (
	"errors"
	"net/http"

	"github.com/f7ed0/golog/lg"
)

type Microservice struct {
	Id       int    `json:"id"`
	Uri      string `json:"uri"`
	LastSign int64  `json:"last_sign"`
}

func (m Microservice) TransactionClose(id string) (err error) {
	resp, err := http.Post(m.Uri+"/transaction/close?transaction_id="+id, "json", nil)
	if err != nil {
		return
	}
	if resp.StatusCode != http.StatusNoContent {
		err = errors.New("Can't close transaction")
	}
	return
}

func (m Microservice) TransactionUndo(id string) {
	// TODO
}

func (m Microservice) IsAlive() bool {
	lg.Info.Println("sending heartbeat to", m.Uri)

	resp, err := http.Get(m.Uri + "/heartbeat")
	if err != nil {
		lg.Error.Println("Error in healthmanager : ", err.Error())
		return false
	}
	lg.Debug.Println(resp.StatusCode == http.StatusNoContent)
	return resp.StatusCode == http.StatusNoContent
}
