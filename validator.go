package main

import (
	"net/http"
)

type Validator interface {
	Validate(r *http.Request) ApplicationError
}
