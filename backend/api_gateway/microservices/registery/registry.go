package registery

import (
	"errors"
	"math/rand"
	"slices"
	"sync"
	"time"

	"github.com/f7ed0/golog/lg"
)

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

func GetMicroservice(mtype string) (Microservice, error) {
	sync_registery.RLock()
	rt, ok := registery[mtype]
	if !ok || len(rt) <= 0 {
		return Microservice{}, errors.New("mtype not known")
	}
	for len(rt) > 0 {
		lg.Debug.Println(len(rt))
		selected_ms := rt[rand.Intn(len(rt))]
		if selected_ms.IsAlive() {
			sync_registery.RUnlock()
			UpdateHeartBeat(selected_ms.Id)
			return rt[rand.Intn(len(rt))], nil
		} else {
			RemoveMicroservice(selected_ms.Id)
		}
	}
	sync_registery.RUnlock()
	lg.Debug.Println("Loop out")
	return Microservice{}, errors.New("No ms of mtype available")
}
