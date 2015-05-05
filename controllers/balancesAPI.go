package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/vancouverwill/accountingApp/models"
	"log"
	"net/http"
	"strconv"
)

type ErrorDetails struct {
	FieldName    string `json:"fieldName"`
	ExpectedType string `json:"expectedType"`
	SuppliedType string `json:"suppliedType"`
}

type ResponseError struct {
	ErrorCode    string `json:"errorCode"`
	ErrorDetails `json:"errorDetails"`
}

/**
*
*
* exp get balance of an account with id relatedToId : `curl -H "Content-Type: application/json" -dG "http://localhost:8080/balances/?AccountAccountHolderOrCompany=Account&relatedToId=18"`
*
* sample return json
*
* {"productBalance" : total sales til now,
*   "paymentBalance" : total payments til now,
*    "paymentBalanceAfterTax" : "total payments minus tax",
*     "balance" : "payments - sales"}
*
* return @param json exp {“sales” :  2200”, payments : 2000, "balance" : 200}
*
**/
func BalancesIndex(response http.ResponseWriter, request *http.Request) {
	log.Println("getBalances START")

	response.Header().Set("Content-Type", "application/json; charset=UTF-8")
	err := request.ParseForm()
	if err != nil {
		http.Error(response, fmt.Sprintf("error parsing url %v", err), 500)
	}

	var AccountAccountHolderOrCompany = request.FormValue("AccountAccountHolderOrCompany")
	var relatedToId = request.FormValue("relatedToId")

	if AccountAccountHolderOrCompany != "Account" && AccountAccountHolderOrCompany != "AccountHolder" && AccountAccountHolderOrCompany != "Company" {
		error := ResponseError{"invalid_field_value_type", ErrorDetails{"AccountAccountHolderOrCompany", "Account|AccountHolder|Company", AccountAccountHolderOrCompany}}
		response.WriteHeader(http.StatusBadRequest)
		if err := json.NewEncoder(response).Encode(error); err != nil {
			panic(err)
		}
	}

	if AccountAccountHolderOrCompany == "Account" || AccountAccountHolderOrCompany == "AccountHolder" {
		relatedToIdInt, err := strconv.Atoi(relatedToId)
		log.Println("relatedToIdInt", relatedToIdInt)
		if err != nil {
			http.Error(response, "{\"error_code\" : \"server_error\"}", 500)
		}
		if relatedToIdInt <= 0 {
			error := ResponseError{"invalid_field_value_type", ErrorDetails{"relatedToId", "positive int", relatedToId}}
			response.WriteHeader(http.StatusBadRequest)
			if err := json.NewEncoder(response).Encode(error); err != nil {
				panic(err)
			}
		}

		if AccountAccountHolderOrCompany == "Account" {
			paymentBalance, paymentBalanceAfterTax, productBalance := models.GetBalanceForAccountId(relatedToIdInt)
			object := make(map[string]float32)
			object["paymentBalance"] = paymentBalance
			object["paymentBalanceAfterTax"] = paymentBalanceAfterTax
			object["productBalance"] = productBalance
			object["balance"] = paymentBalance - productBalance

			response.WriteHeader(http.StatusAccepted)
			if err := json.NewEncoder(response).Encode(object); err != nil {
				panic(err)
			}
		} else {
			paymentBalance, paymentBalanceAfterTax, productBalance := models.GetBalanceForAccountholderId(relatedToIdInt)
			object := make(map[string]float32)
			object["paymentBalance"] = paymentBalance
			object["paymentBalanceAfterTax"] = paymentBalanceAfterTax
			object["productBalance"] = productBalance
			object["balance"] = paymentBalance - productBalance

			response.WriteHeader(http.StatusAccepted)
			if err := json.NewEncoder(response).Encode(object); err != nil {
				panic(err)
			}
		}
	}

	paymentBalance, paymentBalanceAfterTax, productBalance := models.GetBalanceAcrossCompany()
	object := make(map[string]float32)
	object["paymentBalance"] = paymentBalance
	object["paymentBalanceAfterTax"] = paymentBalanceAfterTax
	object["productBalance"] = productBalance
	object["balance"] = paymentBalance - productBalance

	response.WriteHeader(http.StatusAccepted)
	if err := json.NewEncoder(response).Encode(object); err != nil {
		panic(err)
	}

	log.Println("made it to the end")
	response.WriteHeader(http.StatusOK)
}
