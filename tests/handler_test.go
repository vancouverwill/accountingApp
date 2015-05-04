package main

import (
	//	"github.com/gorilla/mux"
	"github.com/vancouverwill/accountingApp/controllers"
	//	"io/ioutil"
	"net/http"
	"net/http/httptest"
	//	"strings"
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

//func TestGet400(t *testing.T) {
//	//    setup()
//	respRec = httptest.NewRecorder()
//	//Testing get of non existent question type
//	req, err = http.NewRequest("GET", "http://localhost:8080/todos", nil)
//	if err != nil {
//		t.Fatal("Creating 'GET /questions/1/SC' request failed!")
//	}

//	m.ServeHTTP(respRec, req)

//	if respRec.Code != http.StatusBadRequest {
//		t.Fatal("Server error: Returned ", respRec.Code, " instead of ", http.StatusBadRequest)
//	}

//	t.Log(respRec.Body)
//}

//func testServer(t *testing.T) {
//	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		t.Log(w, "Hello, client")
//	}))
//	defer ts.Close()

//	res, err := http.Get(ts.URL)
//	if err != nil {
//		t.Error(err)
//	}

//	greeting, err := ioutil.ReadAll(res.Body)
//	res.Body.Close()
//	if err != nil {
//		t.Error(err)
//	}

//	t.Logf("%s", greeting)
//}

// pretty good
/*func TestHeader3D(t *testing.T) {
	resp := httptest.NewRecorder()

	req, err := http.NewRequest("GET", "http://localhost:8080/", nil)
	if err != nil {
		t.Fatal(err)
	}

	http.DefaultServeMux.ServeHTTP(resp, req)
	if p, err := ioutil.ReadAll(resp.Body); err != nil {
		t.Fail()
	} else {
		if strings.Contains(string(p), "Error") {
			t.Errorf("header response shouldn't return error: %s", p)
		} else if !strings.Contains(string(p), `Welcome!`) {
			t.Errorf("header response doen't match:\n%s", p)
		}
	}

	//	t.Log(resp.Body)
	//	t.Log(resp.Code)
}*/

//func TestEchosContent(t *testing.T) {
//	expectedBody := "Welcome!"
//	handler := new(EchoHandler)
//	recorder := httptest.NewRecorder()
//	url := fmt.Sprintf("http://example.com/echo?say=%s", expectedBody)
//	req, err := http.NewRequest("GET", url, nil)
//	assert.Nil(t, err)

//	handler.ServeHTTP(recorder, req)

//	assert.Equal(t, expectedBody, recorder.Body.String())
//}
