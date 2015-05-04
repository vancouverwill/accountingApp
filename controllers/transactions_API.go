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

// getTransactions(start_date,end_date, AccountAccountHolderOrCompany string, accountholder = null, accountId = null)

func TransactionsIndex(w http.ResponseWriter, r *http.Request) {
	log.Println("TransactionsIndex")

	transactions := models.GetTransactions()

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(transactions); err != nil {
		panic(err)
	}
}

func TransactionsTodoCreate(w http.ResponseWriter, r *http.Request) {
	log.Println("TodoCreate")
	var transaction models.Transaction
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &transaction); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}
	log.Println(transaction)
	t := models.MyTransaction.RepoCreateTransaction(transaction)
	//	t := todo
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(t); err != nil {
		panic(err)
	}
}
