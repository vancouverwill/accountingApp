package controllers

import (
	"encoding/json"
	"github.com/vancouverwill/accountingApp/models"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

// createTransaction(account_id, payment_or_product)

// getTransaction(transactionId)

// getTransactions(start_date,end_date, AccountAccountHolderOrCompany string, relatedId = null)

func TransactionsIndex(response http.ResponseWriter, request *http.Request) {
	log.Println("TransactionsIndex")

	transactions := models.GetTransactions()

	response.Header().Set("Content-Type", "application/json; charset=UTF-8")
	response.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(response).Encode(transactions); err != nil {
		panic(err)
	}
}

func TransactionsTodoCreate(response http.ResponseWriter, request *http.Request) {
	log.Println("TransactionCreate")
	var transaction models.Transaction

	transaction = jsonToObject(response, request, transaction)
	log.Println(transaction)
	t := models.MyTransaction.SaveTransaction(transaction)
	//	t := todo
	response.Header().Set("Content-Type", "application/json; charset=UTF-8")
	response.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(response).Encode(t); err != nil {
		panic(err)
	}
}

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
