package main

import (
	"net/http"

	"github.com/dabfleming/go-boilerplate/router"
)

func main() {
	r := router.Router()
	http.ListenAndServe(":3333", r)
}
