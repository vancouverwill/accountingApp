package main

import (
	"github.com/vancouverwill/accountingApp/models"
	"testing"
)

func TestCreateTransaction(t *testing.T) {
	t.Log("TestCreateTransaction")

	models.SaveTransactionByType(5, "product", -650, "test product sale")
	models.SaveTransactionByType(5, "tax", 50, "test tax")
	models.SaveTransactionByType(5, "revenue", 600, "test revenue")

	//	transaction := Transaction{}
	//	var accountHolder models.AccountHolder = models.GetAccountHolderByName("Darrel Mathes")

	//	if accountHolder.JobTitle != "Sales South East" {
	//		t.Error("testGetAccountByAccountName did not work as expected. the address was not as expected, it was ", accountHolder.JobTitle)
	//	}
	//	t.Log("testGetAccountByAccountName successful")
}
