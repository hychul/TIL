package main

import (
	"net/http"

	"example.com/module/myapp"
)

func main() {
	http.ListenAndServe(":3000", myapp.NewHttpHandler())
}
