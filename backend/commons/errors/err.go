package snipetzerror

import "errors"

var ErrorNotImplemented error = errors.New("Not implemented")

var ErrorRefusedByMicroservice error = errors.New("Refused by microservice")
