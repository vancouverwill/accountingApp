package main

import (
	"github.com/vancouverwill/accountingApp/models"
	"log"
	"testing"
)

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
	order.AddItem("Sony Playstation 3", 800)
	order.PrepareRevenue()
	order.PreparePayment()
	order.FinalizeOrder()

}
