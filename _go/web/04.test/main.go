package main

import (
	"net/http"

	"./myapp"
)

func main() {
	http.ListenAndServe(":3000", myapp.NewHttpHandler())
}
