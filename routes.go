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
		"BalancesIndex",
		"GET",
		"/balances",
		controllers.BalancesIndex,
	},
	Route{
		"TransactionsIndex",
		"GET",
		"/transactions",
		controllers.TransactionsIndex,
	},
	Route{
		"Transaction",
		"GET",
		"/transactions/{transactionId}",
		controllers.Transaction,
	},
	Route{
		"TransactionsCreate",
		"POST",
		"/transactions",
		controllers.TransactionsCreate,
	},
	Route{
		"AccountsIndex",
		"GET",
		"/accounts",
		controllers.AccountsIndex,
	},
}
