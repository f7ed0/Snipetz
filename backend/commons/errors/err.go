package snipetzerror

import "errors"

var ErrorNotImplemented error = errors.New("NOT IMPLEMENTED")

var ErrorRefusedByMicroservice error = errors.New("REQUEST REFUSED BY MICROSERVICE")

var ErrorNoJwtSecret error = errors.New("NO JWT TOKEN FOUND IN ENV")
