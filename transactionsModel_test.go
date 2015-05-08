package main

import (
	"github.com/vancouverwill/accountingApp/models"
	"testing"
)

func TestCreateTransaction(t *testing.T) {
	t.Log("TestCreateTransaction")

	models.SaveTransactionByType(5, "revenue", 499, "my first sale")

	//	transaction := Transaction{}
	//	var accountHolder models.AccountHolder = models.GetAccountHolderByName("Darrel Mathes")

	//	if accountHolder.JobTitle != "Sales South East" {
	//		t.Error("testGetAccountByAccountName did not work as expected. the address was not as expected, it was ", accountHolder.JobTitle)
	//	}
	//	t.Log("testGetAccountByAccountName successful")
}
