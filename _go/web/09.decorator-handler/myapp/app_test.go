package myapp

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndex(t *testing.T) {
	assert := assert.New(t)

	ts := httptest.NewServer(NewHandler())
	defer ts.Close()

	res, err := http.Get(ts.URL)
	assert.NoError(err)
	assert.Equal(http.StatusOK, res.StatusCode)

	data, _ := ioutil.ReadAll(res.Body)
	assert.Equal("Hello Web", string(data))
}

func TestGetUsers(t *testing.T) {
	assert := assert.New(t)

	ts := httptest.NewServer(NewHandler())
	defer ts.Close()

	res, err := http.Get(ts.URL + "/users")
	assert.NoError(err)
	assert.Equal(http.StatusOK, res.StatusCode)

	data, _ := ioutil.ReadAll(res.Body)
	assert.Contains("No Users", string(data))

	// create
	res, err = http.Post(ts.URL+"/users", "application/json", strings.NewReader(`{"name":"hychul", "email":"hychome@gmail.com"}`))
	assert.NoError(err)
	assert.Equal(http.StatusCreated, res.StatusCode)
	res, err = http.Post(ts.URL+"/users", "application/json", strings.NewReader(`{"name":"hychul", "email":"hychome@gmail.com"}`))
	assert.NoError(err)
	assert.Equal(http.StatusCreated, res.StatusCode)

	res, err = http.Get(ts.URL + "/users")
	assert.NoError(err)
	assert.Equal(http.StatusOK, res.StatusCode)

	users := []*User{}
	err = json.NewDecoder(res.Body).Decode(&users)
	assert.NoError(err)
	assert.Equal(2, len(users))
}

func TestGetUser(t *testing.T) {
	assert := assert.New(t)

	ts := httptest.NewServer(NewHandler())
	defer ts.Close()

	res, err := http.Get(ts.URL + "/users/1")
	assert.NoError(err)
	assert.Equal(http.StatusOK, res.StatusCode)

	data, _ := ioutil.ReadAll(res.Body)
	assert.Contains("No User ID:1", string(data))
}

func TestCreateUser(t *testing.T) {
	assert := assert.New(t)

	ts := httptest.NewServer(NewHandler())
	defer ts.Close()

	res, err := http.Post(ts.URL+"/users", "application/json", strings.NewReader(`{"name":"hychul", "email":"hychome@gmail.com"}`))
	assert.NoError(err)
	assert.Equal(http.StatusCreated, res.StatusCode)

	user := new(User)
	err = json.NewDecoder(res.Body).Decode(user)
	assert.NoError(err)
	assert.NotEqual(0, user.ID)

	id := user.ID

	// save check
	res, err = http.Get(ts.URL + "/users/" + strconv.Itoa(id))
	assert.NoError(err)
	assert.Equal(http.StatusOK, res.StatusCode)

	savedUser := new(User)
	err = json.NewDecoder(res.Body).Decode(savedUser)
	assert.NoError(err)
	assert.Equal(user.ID, savedUser.ID)
	assert.Equal(user.Name, savedUser.Name)
}

func TestUpdateUser(t *testing.T) {
	assert := assert.New(t)

	ts := httptest.NewServer(NewHandler())
	defer ts.Close()

	req, _ := http.NewRequest("PUT", ts.URL+"/users", strings.NewReader(`{"id":1, "update_name":true, "name":"hyoungchul", "update_email":true, "email":""}`))
	res, err := http.DefaultClient.Do(req)
	assert.NoError(err)
	assert.Equal(http.StatusOK, res.StatusCode)
	data, _ := ioutil.ReadAll(res.Body)
	assert.Contains("No User ID:1", string(data))

	// create
	res, _ = http.Post(ts.URL+"/users", "application/json", strings.NewReader(`{"name":"hychul", "email":"hychome@gmail.com"}`))
	user := new(User)
	json.NewDecoder(res.Body).Decode(user)

	updateStr := fmt.Sprintf(`{"id":%d, "update_name":true, "name":"hyoungchul", "update_email":true, "email":""}`, user.ID)

	req, _ = http.NewRequest("PUT", ts.URL+"/users", strings.NewReader(updateStr))
	res, err = http.DefaultClient.Do(req)
	assert.NoError(err)
	assert.Equal(http.StatusOK, res.StatusCode)

	updateUser := new(User)
	err = json.NewDecoder(res.Body).Decode(updateUser)
	assert.NoError(err)
	assert.Equal(user.ID, updateUser.ID)
	assert.Equal("hyoungchul", updateUser.Name)
	assert.Equal("", updateUser.Email)
}

func TestDeleteUser(t *testing.T) {
	assert := assert.New(t)

	ts := httptest.NewServer(NewHandler())
	defer ts.Close()

	req, _ := http.NewRequest("DELETE", ts.URL+"/users/1", nil)
	res, err := http.DefaultClient.Do(req)
	assert.NoError(err)
	assert.Equal(http.StatusOK, res.StatusCode)

	data, _ := ioutil.ReadAll(res.Body)
	assert.Contains("No User ID:1", string(data))

	// create
	http.Post(ts.URL+"/users", "application/json", strings.NewReader(`{"name":"hychul", "email":"hychome@gmail.com"}`))

	req, _ = http.NewRequest("DELETE", ts.URL+"/users/1", nil)
	res, err = http.DefaultClient.Do(req)
	assert.NoError(err)
	assert.Equal(http.StatusOK, res.StatusCode)

	data, _ = ioutil.ReadAll(res.Body)
	assert.Contains("Deleted User ID:1", string(data))
}
