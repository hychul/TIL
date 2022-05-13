package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/antage/eventsource"
	"github.com/gorilla/pat"
	"github.com/urfave/negroni"
)

type Message struct {
	Sender string `json:"sender"`
	Msg    string `json:"msg"`
}

var msgCh chan Message

func joinUserHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("name")
	broadcastMessage("", fmt.Sprintf("user joined : %s", username))
}

func leftUserHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("name")
	broadcastMessage("", fmt.Sprintf("user left : %s", username))
}

func postMessageHandler(w http.ResponseWriter, r *http.Request) {
	sender := r.FormValue("sender")
	msg := r.FormValue("msg")
	broadcastMessage(sender, msg)
}

func broadcastMessage(sender, msg string) {
	msgCh <- Message{sender, msg}
}

func processMsgCh(es eventsource.EventSource) {
	for msg := range msgCh {
		data, _ := json.Marshal(msg)
		es.SendEventMessage(string(data), "", strconv.Itoa(time.Now().Nanosecond()))
	}
}

func main() {
	msgCh = make(chan Message)

	es := eventsource.New(nil, nil)
	defer es.Close()

	go processMsgCh(es)

	mux := pat.New()
	mux.Post("/users", joinUserHandler)
	mux.Delete("/users", leftUserHandler)
	mux.Post("/messages", postMessageHandler)
	mux.Handle("/stream", es)

	n := negroni.Classic()
	n.UseHandler(mux)

	http.ListenAndServe(":3000", n)
}
