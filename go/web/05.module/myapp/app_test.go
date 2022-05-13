package myapp

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestIndexPathHandler(t *testing.T) {
	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/", nil)

	mux := NewHttpHandler()
	mux.ServeHTTP(res, req)

	if res.Code != http.StatusOK {
		t.Fatal("Failed!: response code is not equal", res.Code)
	}

	data, _ := ioutil.ReadAll(res.Body)
	if "Hello Web" != string(data) {
		t.Fatal("Failed!: response body is not equal", string(data))
	}
}

func TestEchoPathHandler(t *testing.T) {
	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/echo?text=test", nil)

	mux := NewHttpHandler()
	mux.ServeHTTP(res, req)

	if res.Code != http.StatusOK {
		t.Fatal("Failed!: response code is not equal", res.Code)
	}

	data, _ := ioutil.ReadAll(res.Body)
	if "test" != string(data) {
		t.Fatal("Failed!: response body is not equal", string(data))
	}
}

func TestUserPathHandler_WithoutJson(t *testing.T) {
	res := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/user", nil)

	mux := NewHttpHandler()
	mux.ServeHTTP(res, req)

	if res.Code != http.StatusBadRequest {
		t.Fatal("Failed!: response code is not equal", res.Code)
	}
}

func TestUserPathHandler_WithJson(t *testing.T) {
	res := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/user", strings.NewReader(`{"name":"hychul", "email":"hychome@gmail.com"}`))

	mux := NewHttpHandler()
	mux.ServeHTTP(res, req)

	if res.Code != http.StatusOK {
		t.Fatal("Failed!: response code is not equal", res.Code)
	}

	user := new(User)
	err := json.NewDecoder(res.Body).Decode(user)
	if err != nil {
		t.Fatal("Failed!: body json decoding error occured", err)
	}

	if user.Name != "hychul" {
		t.Fatalf("Failed!: name is not equal\nexpected : %s\nactual : %s", "hychul", user.Name)
	}

	if user.Email != "hychome@gmail.com" {
		t.Fatalf("Failed!: name is not equal\nexpected : %s\nactual : %s", "hychome@gmail.com", user.Email)
	}
}
