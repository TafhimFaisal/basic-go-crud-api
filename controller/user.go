package controller

import (
	"net/http"
	"regexp"
)

// UserController to controll user
type UserController struct {
	userIDPattern *regexp.Regexp
}

func (uc UserController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello from user controller !!!"))
}

func newUserController() *UserController {
	return &UserController{
		userIDPattern: regexp.MustCompile(`^/users/(\d+)/?`),
	}
}
