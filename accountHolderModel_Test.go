package main

import (
	"github.com/vancouverwill/accountingApp/models"
	"testing"
)

func TestCreateAccountHolder(t *testing.T) {
	t.Log("TestCreateAccountHolder")

	var accountHolder models.AccountHolder = models.CreateAccountHolder("Oliver Cronwell", "Developer")
	accountHolder.SetAccountHolderCurrency("US DOLLAR")
	accountHolder.SetAccountHolderTaxRate("US Tax")
	accountHolder.Save()

	t.Log("after setting account", accountHolder)

	if accountHolder.Name != "Oliver Cronwell" {
		t.Error("The name is not what we started with!")
	}
}
