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
* exp exp : `curl -H "Content-Type: application/json" -g http://localhost:8080/transactions/?AccountHolderOrCompany=Account&relatedToId=18`
*
* getTransactions possible query string parameters (start_date,end_date, AccountHolderOrCompany string, relatedId = null)
*
*
*
**/

/**
* GET
* returns all transactions
*
*
**/
func TransactionsIndex(response http.ResponseWriter, request *http.Request) {
	log.Println("TransactionsIndex")

	response.Header().Set("Content-Type", "application/json; charset=UTF-8")
	err := request.ParseForm()
	if err != nil {
		http.Error(response, fmt.Sprintf("error parsing url %v", err), 500)
	}

	var AccountHolderOrCompany = request.FormValue("AccountHolderOrCompany")
	var relatedToId = request.FormValue("relatedToId")

	if AccountHolderOrCompany != "AccountHolder" && AccountHolderOrCompany != "Company" {
		error := ResponseError{"invalid_field_value_type", ErrorDetails{"AccountHolderOrCompany", "Account|AccountHolder|Company", AccountHolderOrCompany}}
		response.WriteHeader(http.StatusBadRequest)
		if err := json.NewEncoder(response).Encode(error); err != nil {
			panic(err)
		}
		return
	}

	if AccountHolderOrCompany == "AccountHolder" {
		relatedToIdInt, err := strconv.Atoi(relatedToId)
		log.Println("relatedToIdInt", relatedToIdInt)
		if err != nil {
			http.Error(response, "{\"error_code\" : \"server_error\"}", 500)
			return
		}
		if relatedToIdInt <= 0 {
			error := ResponseError{"invalid_field_value_type", ErrorDetails{"relatedToId", "positive int", relatedToId}}
			response.WriteHeader(http.StatusBadRequest)
			if err := json.NewEncoder(response).Encode(error); err != nil {
				panic(err)
			}
			return
		}

		//		if AccountHolderOrCompany == "Account" {
		//			transactions := models.GetTransactionsForAccountId(relatedToIdInt)

		//			log.Println("GetTransactionsForAccountId", transactions)

		//			response.WriteHeader(http.StatusAccepted)
		//			if err := json.NewEncoder(response).Encode(transactions); err != nil {
		//				panic(err)
		//			}
		//			return
		//		} else {
		transactions := models.GetTransactionsForAccountHolderId(relatedToIdInt)

		response.WriteHeader(http.StatusAccepted)
		if err := json.NewEncoder(response).Encode(transactions); err != nil {
			panic(err)
		}
		return
		//		}
	}

	transactions := models.GetTransactions()

	response.Header().Set("Content-Type", "application/json; charset=UTF-8")
	response.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(response).Encode(transactions); err != nil {
		panic(err)
	}
	return
}

/**
* GET a single transaction by id
*
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

//
// POST - save transaction
//
// example usage
// `curl -H "Content-Type: application/json" -d '{"accountId":9,"details":"buying lots of products AGAIN","amount":201,"date":"2015-01-19T00:00:00Z","updated":0,"created":0}' http://localhost:8080/transactions`
func TransactionsCreate(response http.ResponseWriter, request *http.Request) {
	log.Println("TransactionCreate")
	var transaction models.Transaction

	transaction = jsonToObject(response, request, transaction)
	log.Println(transaction)
	transaction.SaveTransaction()
	response.Header().Set("Content-Type", "application/json; charset=UTF-8")
	response.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(response).Encode(transaction); err != nil {
		panic(err)
	}
}

//
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
