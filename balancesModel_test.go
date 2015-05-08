package main

import (
	"github.com/vancouverwill/accountingApp/models"
	"testing"
)

func TestGetBalancesByCompany(t *testing.T) {
	t.Log("TestGetBalancesByCompany")
	revenue, tax, productSales := models.GetBalanceAcrossCompany()

	//	var ok bool
	//	var temp float32
	if revenue+tax+productSales != 0 {
		t.Error("TestGetBalancesByCompany() did not work as expected. The totals did not come to zero")
	}

	t.Log("testGetAccountByAccountName successful")
}
