package registery

import (
	"errors"
	"math/rand"
	"slices"
	"sync"
	"time"
)

type Microservice struct {
	Id       int    `json:"id"`
	Uri      string `json:"uri"`
	LastSign int64  `json:"last_sign"`
}

type MicroServiceRegistry map[string][]Microservice

var microserviceCount int = 0

var registery MicroServiceRegistry = MicroServiceRegistry{
	"auth": []Microservice{},
}

var sync_registery sync.RWMutex

func (m *MicroServiceRegistry) AddMicroservice(category string, uri string) {
	sync_registery.Lock()
	defer sync_registery.Unlock()
	registery[category] = append(registery[category], Microservice{Uri: uri, Id: microserviceCount, LastSign: time.Now().Unix()})
	microserviceCount++
}

func (m *MicroServiceRegistry) RemoveMicroservice(id int) {
	sync_registery.Lock()
	defer sync_registery.Unlock()
	for k := range registery {
		for i := range registery[k] {
			if registery[k][i].Id == id {
				// removing index i item
				slices.Delete(registery[k], i, i)
			}
		}
	}
}

func (m MicroServiceRegistry) copy() (ret MicroServiceRegistry) {
	ret = make(MicroServiceRegistry)
	for k := range m {
		ret[k] = make([]Microservice, 0)
		copy(ret[k], m[k])
	}
	return
}

func GetMicroservicesRegistery() MicroServiceRegistry {
	sync_registery.RLock()
	defer sync_registery.RUnlock()
	return registery
}

func AddMicroservice(category string, uri string) {
	registery.AddMicroservice(category, uri)
}

func RemoveMicroservice(id int) {
	registery.RemoveMicroservice(id)
}

func UpdateHeartBeat(id int) {
	sync_registery.Lock()
	defer sync_registery.Unlock()
	for k := range registery {
		for i := range registery[k] {
			if registery[k][i].Id == id {
				registery[k][i].LastSign = time.Now().Unix()
			}
		}
	}
}

func GetMicroserviceAddress(mtype string) (string, error) {
	sync_registery.RLock()
	defer sync_registery.RUnlock()
	rt, ok := registery[mtype]
	if !ok || len(rt) <= 0 {
		return "", errors.New("mtype not known")
	}
	// TODO check if it si alive
	return rt[rand.Intn(len(rt))].Uri, nil
}
