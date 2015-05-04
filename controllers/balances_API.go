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
* sales = total sales til now
* payments = total payments til now
* balance = payments - sales
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

		//		var errorDetails string = "{\"error_code\" : \"invalid_field_value_type\","
		//		errorDetails += "\"error_details\" :  { \"fieldName\" : \"AccountAccountHolderOrCompany\","
		//		errorDetails += "\"expectedType\" : \"Account|AccountHolder|Company\","
		//		errorDetails += "\"suppliedType\" : \"not allowed\"}}"

		//		http.Error(response,
		//			errorDetails,
		//			http.StatusBadRequest)
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
			paymentBalance, productBalance := models.GetBalanceForAccountId(relatedToIdInt)
			object := make(map[string]float32)
			object["paymentBalance"] = paymentBalance
			object["productBalance"] = productBalance
			object["balance"] = paymentBalance - productBalance

			response.WriteHeader(http.StatusAccepted)
			if err := json.NewEncoder(response).Encode(object); err != nil {
				panic(err)
			}

		}

	}

	response.WriteHeader(http.StatusOK)
}
