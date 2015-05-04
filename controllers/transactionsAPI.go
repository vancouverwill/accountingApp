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

/**
*
* exp exp : `curl -H "Content-Type: application/json" -g http://localhost:8080/transactions/?AccountAccountHolderOrCompany=Account&relatedToId=18`
*
* getTransactions possible query string parameters (start_date,end_date, AccountAccountHolderOrCompany string, relatedId = null)
*
*
*
**/

/**
*
* returns all transactions
*
*
**/
func TransactionsIndex(response http.ResponseWriter, request *http.Request) {
	log.Println("TransactionsIndex")

	transactions := models.GetTransactions()

	response.Header().Set("Content-Type", "application/json; charset=UTF-8")
	response.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(response).Encode(transactions); err != nil {
		panic(err)
	}
}

/**
* sample usage
* exp : `curl -H "Content-Type: application/json" -g http://localhost:8080/transactions/12`
*
*
**/
func Transaction(response http.ResponseWriter, request *http.Request) {
	log.Println("Transaction")
	vars := mux.Vars(request)
	transactionId := vars["transactionId"]

	transactionIdInt, err := strconv.Atoi(transactionId)
	if err != nil {
		fmt.Print(err)
	}

	transaction := models.GetTransaction(transactionIdInt)

	response.Header().Set("Content-Type", "application/json; charset=UTF-8")
	response.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(response).Encode(transaction); err != nil {
		panic(err)
	}
}

/**
*
* exp usage
* `curl -H "Content-Type: application/json" -d '{"accountId":9,"details":"buying lots of products AGAIN","paymentOrProduct":"product","amount":201,"date":"2015-01-19T00:00:00Z","updated":0,"created":0}' http://localhost:8080/transactions`
*
**/
func TransactionsCreate(response http.ResponseWriter, request *http.Request) {
	log.Println("TransactionCreate")
	var transaction models.Transaction

	transaction = jsonToObject(response, request, transaction)
	log.Println(transaction)
	t := models.MyTransaction.SaveTransaction(transaction)
	response.Header().Set("Content-Type", "application/json; charset=UTF-8")
	response.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(response).Encode(t); err != nil {
		panic(err)
	}
}

/**
*
*
*
**/
func jsonToObject(response http.ResponseWriter, request *http.Request, transaction models.Transaction) models.Transaction {
	body, err := ioutil.ReadAll(io.LimitReader(request.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := request.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &transaction); err != nil {
		response.Header().Set("Content-Type", "application/json; charset=UTF-8")
		response.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(response).Encode(err); err != nil {
			panic(err)
		}
	}
	return transaction
}
