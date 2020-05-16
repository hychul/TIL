package app

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

var rd *render.Render

type Todo struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	CreateAt  time.Time `json:"create_at"`
}

var todoMap map[int]*Todo

func indexHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/todo.html", http.StatusTemporaryRedirect)
}

func getTodoListHandler(w http.ResponseWriter, r *http.Request) {
	list := []*Todo{}
	for _, v := range todoMap {
		list = append(list, v)
	}

	rd.JSON(w, http.StatusOK, list)
}

func addTodoHandler(w http.ResponseWriter, r *http.Request) {

}

func MakeHandler() http.Handler {
	todoMap = make(map[int]*Todo)

	todoMap[1] = &Todo{1, "First", false, time.Now()}
	todoMap[2] = &Todo{2, "Second", true, time.Now()}

	rd = render.New()
	mux := mux.NewRouter()

	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/todos", getTodoListHandler).Methods("GET")
	mux.HandleFunc("/todos", addTodoHandler).Methods("POST")

	return mux
}
