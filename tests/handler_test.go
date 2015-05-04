package main

import (
	"github.com/vancouverwill/accountingApp/controllers"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestIndex(t *testing.T) {
	t.Log("TestIndex")
	req, err := http.NewRequest("GET", "http://localhost:8080/", nil)
	if err != nil {
		t.Error("index() did not work as expected.")
		t.Log(err)
	}

	w := httptest.NewRecorder()
	controllers.Index(w, req)

	if w.Code != 200 {
		t.Error("index() did not work as expected. the status was not 200")
	}

	t.Log("status:", w.Code, "body:", w.Body.String())
}

func TestTodoIndex(t *testing.T) {
	t.Log("TestTodoIndex")
	handler := controllers.TodoIndex

	req, err := http.NewRequest("GET", "http://localhost:8080/todos", nil)
	if err != nil {
		t.Error("index() did not work as expected.")
		t.Log(err)
	}

	w := httptest.NewRecorder()
	handler(w, req)

	if w.Code != 200 {
		t.Error("TodoIndex() did not work as expected. the status was not 200")
	}

	//	t.Log("status:", w.Code, "body:", w.Body.String())
}
