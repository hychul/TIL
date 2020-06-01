package app

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
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
