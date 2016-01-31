package main

import (
	"github.com/dabfleming/go-boilerplate/router"
	"net/http"
)

func main() {
	r := router.Router()
	http.ListenAndServe(":3333", r)
}
