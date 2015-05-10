package main

import (
	"github.com/vancouverwill/accountingApp/models"
	"testing"
)

func TestGetBalancesByCompany(t *testing.T) {
	t.Log("TestGetBalancesByCompany")
	revenue, tax, payment := models.GetBalanceAcrossCompany()

	t.Log("revenue", revenue, "tax", tax, "payment", payment)
	if revenue+tax+payment != 0 {
		t.Error("TestGetBalancesByCompany() did not work as expected. The totals did not come to zero")
	}
	t.Log("testGetAccountByAccountName successful")
}

func TestCreateUserAddAnOrderThenVerifyBalanceIsZero(t *testing.T) {
	t.Log("TestCreateUserAddAnOrderThenVerifyBalanceIsZero")
	var accountHolder models.AccountHolder = models.CreateAccountHolder("Issac Newton", "Engineer")
	accountHolder.SetAccountHolderCurrency("US DOLLAR")
	accountHolder.SetAccountHolderTaxRate("US Tax")
	accountHolder.Save()
	t.Log("accountHolder", accountHolder)
	order := accountHolder.NewOrder()
	t.Log("order", order)

	revenue, tax, payment := models.GetBalanceForAccountholderId(accountHolder.Id)

	if revenue != 0 || tax != 0 || payment != 0 {
		t.Error("TestCreateUserAddAnOrderThenVerifyBalanceIsZero() did not work as expected.")
		t.Error("The overall did not add to zero balances", revenue, tax, payment)
	}
}
