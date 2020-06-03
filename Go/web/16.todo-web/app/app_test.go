package app

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddTodo(t *testing.T) {
	// given
	assert := assert.New(t)
	ts := httptest.NewServer(MakeHandler())

	defer ts.Close()

	// when
	resp, err := http.PostForm(ts.URL+"/todos", url.Values{"name": {"Test todo"}})

	//then
	assert.NoError(err)
	assert.Equal(http.StatusCreated, resp.StatusCode)
	var todo Todo
	err = json.NewDecoder(resp.Body).Decode(&todo)
	assert.NoError(err)
	assert.Equal(todo.Name, "Test todo")
}

func TestGetTodos(t *testing.T) {
	// given
	assert := assert.New(t)
	ts := httptest.NewServer(MakeHandler())

	defer ts.Close()

	resp, err := http.PostForm(ts.URL+"/todos", url.Values{"name": {"Test todo 1"}})
	assert.NoError(err)
	assert.Equal(http.StatusCreated, resp.StatusCode)
	var todo Todo
	err = json.NewDecoder(resp.Body).Decode(&todo)
	assert.NoError(err)
	assert.Equal(todo.Name, "Test todo 1")
	id1 := todo.ID

	resp, err = http.PostForm(ts.URL+"/todos", url.Values{"name": {"Test todo 2"}})
	assert.NoError(err)
	assert.Equal(http.StatusCreated, resp.StatusCode)
	err = json.NewDecoder(resp.Body).Decode(&todo)
	assert.NoError(err)
	assert.Equal(todo.Name, "Test todo 2")
	id2 := todo.ID

	// when
	resp, err = http.Get(ts.URL + "/todos")

	// then
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)
	todos := []*Todo{}
	err = json.NewDecoder(resp.Body).Decode(&todos)
	assert.NoError(err)
	assert.Equal(len(todos), 2)
	for _, t := range todos {
		if t.ID == id1 {
			assert.Equal(t.Name, "Test todo 1")
		} else if t.ID == id2 {
			assert.Equal(t.Name, "Test todo 2")
		} else {
			assert.Error(fmt.Errorf("ID MUST BE %d or %d", id1, id2))
		}
	}
}

func TestUpdateTodo(t *testing.T) {
func TestUpdateTodo(t *testing.T) {
	// given
	assert := assert.New(t)
	ts := httptest.NewServer(MakeHandler())

	defer ts.Close()

	resp, err := http.PostForm(ts.URL+"/todos", url.Values{"name": {"Test todo 1"}})
	assert.NoError(err)
	assert.Equal(http.StatusCreated, resp.StatusCode)
	var todo Todo
	err = json.NewDecoder(resp.Body).Decode(&todo)
	assert.NoError(err)
	assert.Equal(todo.Name, "Test todo 1")
	id1 := todo.ID

	resp, err = http.PostForm(ts.URL+"/todos", url.Values{"name": {"Test todo 2"}})
	assert.NoError(err)
	assert.Equal(http.StatusCreated, resp.StatusCode)
	err = json.NewDecoder(resp.Body).Decode(&todo)
	assert.NoError(err)
	assert.Equal(todo.Name, "Test todo 2")
	id2 := todo.ID

	// when
	req, _ := http.NewRequest("PUT", ts.URL+"/todos/"+strconv.Itoa(id1)+"?complete=true", nil)
	resp, err = http.DefaultClient.Do(req)

	// then
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)
	resp, err = http.Get(ts.URL + "/todos")
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)
	todos := []*Todo{}
	err = json.NewDecoder(resp.Body).Decode(&todos)
	assert.NoError(err)
	assert.Equal(len(todos), 2)
	for _, t := range todos {
		if t.ID == id1 {
			assert.Equal(t.Name, "Test todo 1")
			assert.Equal(t.Completed, true)
		} else if t.ID == id2 {
			assert.Equal(t.Name, "Test todo 2")
			assert.Equal(t.Completed, false)
		} else {
			assert.Error(fmt.Errorf("ID MUST BE %d or %d", id1, id2))
		}
	}
}
