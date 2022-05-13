package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/gorilla/pat"
)

var userMap map[int]*User
var lastID int

type User struct {
	ID    int    `json:id`
	Name  string `json:name`
	Email string `json:email`
}

func getIndexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.New("Index").ParseFiles("templates/index.tmpl")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
		return
	}
	tmpl.ExecuteTemplate(w, "index.tmpl", "Hychul")
}

func getUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}

	// >>> select user
	user, ok := userMap[id]
	if !ok {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "No User ID:", id)
		return
	}
	// <<<

	w.WriteHeader(http.StatusOK)
	w.Header().Add("content-type", "application/json")
	data, _ := json.Marshal(user)
	fmt.Fprint(w, string(data))
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	user := new(User)
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}

	// >>> insert user
	lastID++

	user.ID = lastID

	userMap[user.ID] = user
	// <<<

	w.WriteHeader(http.StatusCreated)
	w.Header().Add("content-type", "application/json")
	data, _ := json.Marshal(user)
	fmt.Fprint(w, string(data))
}

func main() {
	mux := pat.New()

	mux.Get("/", getIndexHandler))
	mux.Get("/users", getUserHandler)
	mux.Post("/users", createUserHandler)

	http.ListenAndServe(":3000", mux)
}
