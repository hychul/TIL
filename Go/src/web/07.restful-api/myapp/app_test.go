package myapp

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestIndex(t *testing.T) {
	ts := httptest.NewServer(NewHandler())
	defer ts.Close()

	res, err := http.Get(ts.URL)
	if err != nil {
		t.Fatal("Failed!: get index has error")
	}
	if res.StatusCode != http.StatusOK {
		t.Fatal("Failed!: status code is not OK", res.StatusCode)
	}

	data, _ := ioutil.ReadAll(res.Body)
	if string(data) != "Hello Web" {
		t.Fatalf("Failed!: response body is not equal\nexpacted : %s\nactual : %s", "Hello Web", string(data))
	}
}

func TestUsers(t *testing.T) {
	ts := httptest.NewServer(NewHandler())
	defer ts.Close()

	res, err := http.Get(ts.URL + "/users")
	if err != nil {
		t.Fatal("Failed!: get index has error")
	}
	if res.StatusCode != http.StatusOK {
		t.Fatal("Failed!: status code is not OK", res.StatusCode)
	}

	data, _ := ioutil.ReadAll(res.Body)
	if !strings.Contains(string(data), "Get UserInfo") {
		t.Fatal("Failed!: response body is not contained \"Get UserInfo\"")
	}
}
