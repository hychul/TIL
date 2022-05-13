package myapp // this package MUST be under $GOPATH/github.com/hychul/myapp/

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type User struct {
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	CreateAt time.Time `json:"create_at"`
}

type echoHandler struct{}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello Web")
}

func (f *echoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	echoText := r.URL.Query().Get("text")
	fmt.Fprintf(w, "%s", echoText)
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	user := new(User)

	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}

	user.CreateAt = time.Now()

	json, _ := json.Marshal(user)
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(json))
}

func NewHttpHandler() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", indexHandler)

	mux.Handle("/echo", &echoHandler{})

	mux.HandleFunc("/user", userHandler)

	return mux
}
