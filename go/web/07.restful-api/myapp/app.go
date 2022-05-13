package myapp

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

var userMap map[int]*User
var lastID int

type User struct {
	ID       int       `json:"id"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	CreateAt time.Time `json:"create_at"`
}

type UpdateUser struct {
	ID          int       `json:"id"`
	UpdateName  bool      `json:"update_name"`
	Name        string    `json:"name"`
	UpdateEmail bool      `json:"update_email"`
	Email       string    `json:"email"`
	CreateAt    time.Time `json:"create_at"`
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello Web")
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
	user.CreateAt = time.Now()

	userMap[user.ID] = user
	// <<<

	w.WriteHeader(http.StatusCreated)
	w.Header().Add("content-type", "application/json")
	data, _ := json.Marshal(user)
	fmt.Fprint(w, string(data))
}

func getUsersHandler(w http.ResponseWriter, r *http.Request) {
	if len(userMap) == 0 {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "No Users")
		return
	}

	users := []*User{}
	for _, user := range userMap {
		users = append(users, user)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Add("content-type", "application/json")
	data, _ := json.Marshal(users)
	fmt.Fprint(w, string(data))
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

func updateUserHandler(w http.ResponseWriter, r *http.Request) {
	updateUser := new(UpdateUser)
	err := json.NewDecoder(r.Body).Decode(updateUser)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}

	// >>> select user
	user, ok := userMap[updateUser.ID]
	if !ok {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "No User ID:", updateUser.ID)
		return
	}
	// <<<

	// >>> patch
	if updateUser.UpdateName {
		user.Name = updateUser.Name
	}

	if updateUser.UpdateEmail {
		user.Email = updateUser.Email
	}
	// <<<

	// >>> update user
	userMap[user.ID] = user
	// <<<

	w.WriteHeader(http.StatusOK)
	w.Header().Add("content-type", "application/json")
	data, _ := json.Marshal(user)
	fmt.Fprint(w, string(data))
}

func deleteUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}

	_, ok := userMap[id]
	if !ok {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "No User ID:", id)
		return
	}

	delete(userMap, id)
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Deleted User ID:", id)
}

func NewHandler() http.Handler {
	userMap = make(map[int]*User)
	lastID = 0

	mux := mux.NewRouter()

	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/users", createUserHandler).Methods("POST")
	mux.HandleFunc("/users", getUsersHandler).Methods("GET")
	mux.HandleFunc("/users/{id:[0-9]+}", getUserHandler).Methods("GET")
	mux.HandleFunc("/users", updateUserHandler).Methods("PUT")
	mux.HandleFunc("/users/{id:[0-9]+}", deleteUserHandler).Methods("DELETE")

	return mux
}
