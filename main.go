package main

import (
	"net/http"

	"github.com/TafhimFaisal/golang/controller"
)

func main() {
	controller.RegistarController()
	http.ListenAndServe(":3000", nil)
}
