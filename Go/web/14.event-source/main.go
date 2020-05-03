package main

import (
	"log"
	"net/http"

	"github.com/gorilla/pat"
	"github.com/urfave/negroni"
)

func postMessageHandler(w http.ResponseWriter, r *http.Request) {
	sender := r.FormValue("sender")
	msg := r.FormValue("msg")
	log.Println(sender, msg)
}

func main() {
	mux := pat.New()

	mux.Post("/messages", postMessageHandler)

	n := negroni.Classic()
	n.UseHandler(mux)

	http.ListenAndServe(":3000", n)
}
