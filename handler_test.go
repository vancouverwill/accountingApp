package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

//func Test_Index(t *testing.T) {

//	if Index() < 0 {
//		t.Error("Add2Ints did not work as expected.")
//	} else {
//		t.log("content was returned")
//	}
//}

func Test_Index(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "something failed", http.StatusInternalServerError)
	}

	req, err := http.NewRequest("GET", "http://localhost:8080/todos", nil)
	if err != nil {
		t.Error("index() did not work as expected.")
		t.Log(err)

	}

	w := httptest.NewRecorder()
	handler(w, req)

	t.Log("%d - %s", w.Code, w.Body.String())

	t.Log(req)
}
