package main

import (
	"github.com/vancouverwill/accountingApp/models"
	"log"
	"testing"
)

/*func TestCreateTransaction(t *testing.T) {
	t.Log("TestCreateTransaction")

	// manually add in dual accounting just to test the save transaction is working. the actual caculations of how it should be split up are handled by the order model
	models.SaveTransactionByType(5, "payment", -650, "test product sale")
	models.SaveTransactionByType(5, "tax", 50, "test tax")
	models.SaveTransactionByType(5, "revenue", 600, "test revenue")

	//	transaction := Transaction{}
	//	var accountHolder models.AccountHolder = models.GetAccountHolderByName("Darrel Mathes")

	//	if accountHolder.JobTitle != "Sales South East" {
	//		t.Error("testGetAccountByAccountName did not work as expected. the address was not as expected, it was ", accountHolder.JobTitle)
	//	}
	//	t.Log("testGetAccountByAccountName successful")
}*/

func prepareAccountForTesting(name string, job string) models.AccountHolder {
	var accountHolder models.AccountHolder = models.CreateAccountHolder("Issac Newton", "Engineer")
	accountHolder.SetAccountHolderCurrency("US DOLLAR")
	accountHolder.SetAccountHolderTaxRate("US Tax")
	accountHolder.Save()

	return accountHolder
}

func TestCreateOrder(t *testing.T) {
	accountHolder := prepareAccountForTesting("Lucas Brasi", "Developer")
	log.Println(accountHolder)
	order := accountHolder.NewOrder()
	order.AddItem("Sony Playstation 3", 799)
	order.PrepareRevenue()
	order.PreparePayment()
	order.FinalizeOrder()
}
