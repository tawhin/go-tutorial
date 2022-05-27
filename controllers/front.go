package controllers

import (
	"encoding/json"
	"io"
	"net/http"
)

func RegisterControllers() {
	uc := newUserController()

	// Go will bind the http.Handle interface to the userController because userController implements a matching set of
	// interface function definitions.
	http.Handle("/users", *uc)
	http.Handle("/users/", *uc)
}

func encodeResponseAsJson(data interface{}, w io.Writer) {
	enc := json.NewEncoder(w)
	enc.Encode(data)
	// As we are writing to the ResponseWriter then the header will be implicitly set
	// to StatusOK
}
