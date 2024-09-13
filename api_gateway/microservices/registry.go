package microservices

import "slices"

type Microservice struct {
	id  int
	uri string
}

type MicroServiceRegistry map[string][]Microservice

var microserviceCount int = 0

var registery MicroServiceRegistry = MicroServiceRegistry{
	"auth": []Microservice{},
}

func (m MicroServiceRegistry) AddMicroservice(category string, uri string) {
	registery[category] = append(registery[category], Microservice{uri: uri, id: microserviceCount})
	microserviceCount++
}

func (m MicroServiceRegistry) RemoveMicroservice(id int) {
	for k := range registery {
		for i := range registery[k] {
			if registery[k][i].id == id {
				// removing index i item
				slices.Delete(registery[k], i, i)
			}
		}
	}
}
