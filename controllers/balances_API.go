package controllers

import (
	"log"
	"net/http"
)

/**
*
* sales = total sales til now
* payments = total payments til now
* balance = payments - sales
*
* return @param json exp {“sales” :  2200”, payments : 2000, "balance" : 200}
*
**/
//get_balances( AccountAccountHolderOrCompany string)

func BalancesIndex(w http.ResponseWriter, r *http.Request) {
	log.Println("getBalances")
}
