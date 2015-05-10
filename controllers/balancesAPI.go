package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/vancouverwill/accountingApp/models"
	"log"
	"net/http"
	"strconv"
)

//
// query paramters : AccountHolderOrCompany & relatedToId
//
// exp get balance of an account with id relatedToId : `curl -H "Content-Type: application/json" -dG "http://localhost:8080/balances/?AccountHolderOrCompany=Account&relatedToId=18"`
//
// sample return json
//
// {"revenueBalance" : total sales til now,
//   "paymentBalance" : total payments til now,
//    "paymentBalanceAfterTax" : "total payments minus tax",
//     "balance" : "payments - sales"}
//
// return @param json exp {“sales” :  2200”, payments : 2000, "balance" : 200}
//
func BalancesIndex(response http.ResponseWriter, request *http.Request) {
	log.Println("getBalances START")

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
		//		if AccountHolderOrCompany == "Account" || AccountHolderOrCompany == "AccountHolder" {
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
			return
		}

		revenueBalance, Tax, paymentBalance := models.GetBalanceForAccountholderId(relatedToIdInt)
		object := make(map[string]float32)
		object["revenueBalance"] = revenueBalance
		object["Tax"] = Tax
		object["paymentBalance"] = paymentBalance
		object["balance"] = revenueBalance - paymentBalance

		response.WriteHeader(http.StatusAccepted)
		if err := json.NewEncoder(response).Encode(object); err != nil {
			panic(err)

		}
		return
	}

	revenueBalance, Tax, paymentBalance := models.GetBalanceAcrossCompany()
	object := make(map[string]float32)
	object["revenueBalance"] = revenueBalance
	object["Tax"] = Tax
	object["paymentBalance"] = paymentBalance
	object["balance"] = revenueBalance - paymentBalance

	response.WriteHeader(http.StatusAccepted)
	if err := json.NewEncoder(response).Encode(object); err != nil {
		panic(err)
	}

	log.Println("made it to the end")
	response.WriteHeader(http.StatusOK)
	return
}
