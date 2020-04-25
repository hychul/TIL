package main

import (
	"log"
	"net/http"
	"time"

	"example.com/decorator/myapp"
	"example.com/decorator/webdecorator"
)

func logger(w http.ResponseWriter, r *http.Request, h http.Handler) {
	startTime := time.Now()
	log.Println("[LOGGER] Started")
	h.ServeHTTP(w, r)
	log.Println("[LOGGER] Finished. Elapsed time:", time.Since(startTime).Milliseconds())
}

func NewHandler() http.Handler {
	mux := myapp.NewHandler()
	mux = webdecorator.NewHandler(mux, logger)
	return mux
}

func main() {
	http.ListenAndServe(":3000", NewHandler())
}
