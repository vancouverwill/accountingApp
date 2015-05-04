package main

import (
	"github.com/vancouverwill/accountingApp/controllers"
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		controllers.Index,
	},
	Route{
		"TodoIndex",
		"GET",
		"/todos",
		controllers.TodoIndex,
	},
	Route{
		"TodoShow",
		"GET",
		"/todos/{todoId}",
		controllers.TodoShow,
	},
	Route{
		"TodoCreate",
		"POST",
		"/todos",
		controllers.TodoCreate,
	},
	Route{
		"TransactionsIndex",
		"GET",
		"/transactions",
		controllers.TransactionsIndex,
	},
	Route{
		"TransactionsTodoCreate",
		"POST",
		"/transactions",
		controllers.TransactionsTodoCreate,
	},
}
