package controller

import (
	"encoding/json"
	"io"
	"net/http"
)

// RegistarController Registar Controller
func RegistarController() {
	uc := newUserController()

	http.Handle("/user", *uc)
	http.Handle("/user/", *uc)
}

func encodeResponseAsJSON(data interface{}, w io.Writer) {
	enc := json.NewEncoder(w)
	enc.Encode(data)
}
