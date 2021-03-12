package controller

import (
	"encoding/json"
	"net/http"
	"regexp"

	"github.com/TafhimFaisal/basic_crud_with_golang/moduls"
)

// UserController to controll user
type UserController struct {
	userIDPattern *regexp.Regexp
}

func (uc UserController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/users" {
		switch r.Method {
		case http.MethodGet:
			uc.getAll(w, r)
		case http.MethodPost:
			uc.post(w, r)
		default:
			w.WriteHeader(http.StatusNotImplemented)
		}
	} else {

	}
	// w.Write([]byte("hello from user controller !!!"))
}

func (uc *UserController) getAll(w http.ResponseWriter, r *http.Request) {
	encodeResponseAsJSON(moduls.GetUsers(), w)
}

func (uc *UserController) get(id int, w http.ResponseWriter) {
	u, err := moduls.GetUserByID(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	encodeResponseAsJSON(u, w)
}

func (uc *UserController) post(w http.ResponseWriter, r *http.Request) {
	u, err := uc.parseRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	u, err = moduls.AddUser(u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	encodeResponseAsJSON(u, w)
}

func (uc *UserController) put(id int, w http.ResponseWriter, r *http.Request) {
	u, err := uc.parseRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("could not pars user object"))
		return
	}

	if id != u.ID {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("User not found"))
		return
	}

	u, err = moduls.UpdateUser(u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	encodeResponseAsJSON(u, w)
}

func (uc *UserController) delete(id int, w http.ResponseWriter) {
	err := moduls.RemoveUserbyId(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)

}

func (uc *UserController) parseRequest(r *http.Request) (moduls.User, error) {
	dec := json.NewDecoder(r.Body)
	var u moduls.User
	err := dec.Decode(&u)

	if err != nil {
		return moduls.User{}, err
	}

	return u, nil
}

func newUserController() *UserController {
	return &UserController{
		userIDPattern: regexp.MustCompile(`^/users/(\d+)/?`),
	}
}
