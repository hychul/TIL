package app

import (
	"net/http"
	"strconv"

	"example.com/todo/model"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

var rd *render.Render

func indexHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/todo.html", http.StatusTemporaryRedirect)
}

func getTodoListHandler(w http.ResponseWriter, r *http.Request) {
	list := model.GetTodos()

	rd.JSON(w, http.StatusOK, list)
}

func createTodoHandler(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	todo := model.AddTodo(name)
	rd.JSON(w, http.StatusCreated, todo)
}

func updateTodoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	complete := r.FormValue("complete") == "true"

	ok := model.CompleteTodo(id, complete)
	if ok {
		rd.JSON(w, http.StatusOK, model.Success{true})
	} else {
		rd.JSON(w, http.StatusOK, model.Success{false})
	}
}

func removeTodoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	ok := model.RemoveTodo(id)
	if ok {
		rd.JSON(w, http.StatusOK, model.Success{true})
	} else {
		rd.JSON(w, http.StatusOK, model.Success{false})
	}
}

func MakeHandler() http.Handler {
	rd = render.New()
	mux := mux.NewRouter()

	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/todos", getTodoListHandler).Methods("GET")
	mux.HandleFunc("/todos", createTodoHandler).Methods("POST")
	mux.HandleFunc("/todos/{id:[0-9]+}", updateTodoHandler).Methods("PUT")
	mux.HandleFunc("/todos/{id:[0-9]+}", removeTodoHandler).Methods("DELETE")

	return mux
}
