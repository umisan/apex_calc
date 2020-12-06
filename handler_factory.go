package main

import "net/http"

type HandlerFactory struct {
}

func (f HandlerFactory) Build(v View) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := v.Validate(r)
		if err != nil {
			handleError(w, err)
			return
		}

		v.WriteResponse(w, r)
	}
}
