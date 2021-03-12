package main

import (
	"net/http"

	"github.com/TafhimFaisal/basic_crud_with_golang/controller"
)

func main() {
	controller.RegistarController()
	http.ListenAndServe(":3000", nil)
}
