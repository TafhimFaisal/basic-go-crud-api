package controller

import "net/http"

// RegistarController Registar Controller
func RegistarController() {
	uc := newUserController()

	http.Handle("/user", *uc)
	http.Handle("/user/", *uc)
}
