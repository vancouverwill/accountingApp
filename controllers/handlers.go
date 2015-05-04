package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/vancouverwill/accountingApp/models"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func Index(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(response, "Welcome!")
}

func TodoIndex(response http.ResponseWriter, request *http.Request) {
	log.Println("TodoIndex")

	todos := models.RepoGetTodos()

	response.Header().Set("Content-Type", "application/json; charset=UTF-8")
	response.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(response).Encode(todos); err != nil {
		panic(err)
	}
}

func TodoShow(response http.ResponseWriter, request *http.Request) {
	log.Println("TodoShow")
	vars := mux.Vars(request)
	todoId := vars["todoId"]

	todoIdString, err := strconv.Atoi(todoId)
	if err != nil {
		fmt.Print(err)
	}

	todo := models.RepoFindTodo(todoIdString)
	//	fmt.Fprintln(w, "Todo show:", todoId)

	response.Header().Set("Content-Type", "application/json; charset=UTF-8")
	response.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(response).Encode(todo); err != nil {
		panic(err)
	}
}

func TodoCreate(response http.ResponseWriter, request *http.Request) {
	log.Println("TodoCreate")
	var todo models.Todo
	body, err := ioutil.ReadAll(io.LimitReader(request.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := request.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &todo); err != nil {
		response.Header().Set("Content-Type", "application/json; charset=UTF-8")
		response.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(response).Encode(err); err != nil {
			panic(err)
		}
	}
	log.Println(todo)
	t := models.RepoCreateTodo(todo)
	//	t := todo
	response.Header().Set("Content-Type", "application/json; charset=UTF-8")
	response.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(response).Encode(t); err != nil {
		panic(err)
	}
}
