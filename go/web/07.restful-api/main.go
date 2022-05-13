package main

import (
	"net/http"

	"example.com/restful/myapp"
)

func main() {
	http.ListenAndServe(":3000", myapp.NewHandler())
}
