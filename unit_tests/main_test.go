package main_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"../app"
)

var a app.App

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d, but got %d\n", expected, actual)
	}
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	a.Router.ServeHTTP(rr, req)

	return rr
}

func checkHeader(t *testing.T, expected, actual string) {
	if expected != actual {
		t.Errorf("Expected response header %s, bout got %s\n", expected, actual)
	}
}

func TestPeople(t *testing.T) {
	a.Init()
	req, _ := http.NewRequest("GET", "/people", nil)
	reponse := executeRequest(req)

	checkResponseCode(t, http.StatusOK, reponse.Code)
	checkHeader(t, reponse.Header().Get("Content-Type"), "application/json; charset=UTF-8")
}

func TestWrongStatus(t *testing.T) {
	a.Init()
	req, _ := http.NewRequest("GET", "/", nil)
	reponse := executeRequest(req)

	checkResponseCode(t, http.StatusNotFound, reponse.Code)
}

func TestCreateNewPerson(t *testing.T) {
	a.Init()
	req, _ := http.NewRequest("POST", "/people/4", strings.NewReader(`{"id":"4","firstname":"Ron","lastname":"Smith","address":{"city":"City YY","state":"State YY"}}`))
	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)
	checkHeader(t, response.Header().Get("Content-Type"), "application/json; charset=UTF-8")

	var m map[interface{}]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)

	t.Errorf("fail!")
}
