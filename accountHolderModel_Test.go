package main

import (
	"github.com/vancouverwill/accountingApp/models"
	"testing"
)

func TestCreateAccountHolder(t *testing.T) {
	t.Log("TestCreateAccountHolder")

	var accountHolder models.AccountHolder = models.CreateAccountHolder("James", "Developer")
	accountHolder.SetAccountHolderCurrency("US DOLLAR")
	accountHolder.SetAccountHolderTaxRate("US Tax")

	t.Log("after setting account", accountHolder)

}
