package main

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

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func TodoIndex(w http.ResponseWriter, r *http.Request) {
	log.Println("TodoIndex")
	//	todos := Todos{
	//		Todo{Name: "Write presentation"},
	//		Todo{Name: "Host meetup"},
	//	}

	todos := models.RepoGetTodos()

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(todos); err != nil {
		panic(err)
	}
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
	log.Println("TodoShow")
	vars := mux.Vars(r)
	todoId := vars["todoId"]

	todoIdString, err := strconv.Atoi(todoId)
	if err != nil {
		fmt.Print(err)
	}

	todo := models.RepoFindTodo(todoIdString)
	//	fmt.Fprintln(w, "Todo show:", todoId)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(todo); err != nil {
		panic(err)
	}
}

func TodoCreate(w http.ResponseWriter, r *http.Request) {
	log.Println("TodoCreate")
	var todo models.Todo
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &todo); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}
	log.Println(todo)
	t := models.RepoCreateTodo(todo)
	//	t := todo
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(t); err != nil {
		panic(err)
	}
}
