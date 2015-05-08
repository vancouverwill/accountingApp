package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/vancouverwill/accountingApp/models"
	"log"
	"net/http"
)

func AccountsIndex(response http.ResponseWriter, request *http.Request) {
	log.Println("getBalances START")
	response.Header().Set("Content-Type", "application/json; charset=UTF-8")
	err := request.ParseForm()
	if err != nil {
		http.Error(response, fmt.Sprintf("error parsing url %v", err), 500)
	}

	var AccountName = request.FormValue("AccountName")

	log.Println("AccountName", AccountName)

	account := models.GetAccountHolderByName(AccountName)

	response.Header().Set("Content-Type", "application/json; charset=UTF-8")
	response.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(response).Encode(account); err != nil {
		panic(err)
	}
}
