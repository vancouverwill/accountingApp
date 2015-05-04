package controllers

import (
	"encoding/json"
	"fmt"
	//	"github.com/gorilla/mux"
	//	"github.com/vancouverwill/accountingApp/models"
	"log"
	"net/http"
	//	"strconv"
)

func Index(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(response, "Welcome!")
}

func FourZeroFourPage(response http.ResponseWriter, request *http.Request) {
	log.Println("FourZeroFourPage")
	response.Header().Set("Content-Type", "application/json; charset=UTF-8")
	response.WriteHeader(http.StatusNotFound)

	error := make(map[string]string)
	error["error_code"] = "page_not_found"

	log.Println(error)

	if err := json.NewEncoder(response).Encode(error); err != nil {
		panic(err)
	}
}
