package main

import (
	"github.com/dabfleming/foo/router"
	"net/http"
)

func main() {
	r := router.Router()
	http.ListenAndServe(":3333", r)
}
