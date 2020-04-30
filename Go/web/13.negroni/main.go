package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/gorilla/pat"
	"github.com/unrolled/render"
)

var rd *render.Render

var userMap map[int]*User
var lastID int

type User struct {
	ID    int    `json:id`
	Name  string `json:name`
	Email string `json:email`
}

func getIndexHandler(w http.ResponseWriter, r *http.Request) {
	// rd.HTML(w, http.StatusOK, "index-tmpl", "Hychul")
	rd.HTML(w, http.StatusOK, "index-html", "Hychul")
}

func getUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		rd.Text(w, http.StatusBadRequest, err.Error())
		return
	}

	// >>> select user
	user, ok := userMap[id]
	if !ok {
		rd.Text(w, http.StatusOK, fmt.Sprint("No User ID:", id))
		return
	}
	// <<<

	rd.JSON(w, http.StatusOK, user)
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	user := new(User)
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		rd.Text(w, http.StatusBadRequest, err.Error())
		return
	}

	// >>> insert user
	lastID++

	user.ID = lastID

	userMap[user.ID] = user
	// <<<

	rd.JSON(w, http.StatusCreated, user)
}

func main() {
	rd = render.New(render.Options{
		Directory:  "htmls",
		Extensions: []string{".html", ".tmpl"},
	})
	mux := pat.New()

	mux.Get("/", getIndexHandler)
	mux.Get("/users", getUserHandler)
	mux.Post("/users", createUserHandler)

	http.ListenAndServe(":3000", mux)
}
