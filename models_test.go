package main

import (
	"github.com/vancouverwill/accountingApp/models"
	"testing"
)

func TestGetAccountByAccountName(t *testing.T) {
	t.Log("testGetAccountByAccountName")
	var accountHolder models.AccountHolder = models.GetAccountHolderByName("Darrel Mathes")

	if accountHolder.JobTitle != "Sales South East" {
		t.Error("testGetAccountByAccountName did not work as expected. the address was not as expected, it was ", accountHolder.JobTitle)
	}
	t.Log("testGetAccountByAccountName successful")
}

/*func testGetAccountHolderByAccountName(t *testing.T) {
	var account models.Account = models.GetAccountByName("Atlanta WholeSale")

	if account.Address != "200 Centre Avenue" {
		t.Error("testGetAccountByAccountName did not work as expected. the address was not as expected, it was ", account.Address)
	}
	t.Log("testGetAccountByAccountName successful")
}*/

/** test transactions return expected no of transactions for an accountID
*
* 1 - get transactions for accountId
* 2 - record no of transactions at start
* 3 - add new 5 transactions for accountID
* 4 - verify accounts after - accounts before is equal to five
*
**/

/**
*
* 1 - delete transactions for user
* 2- add in some payment transactions and some product transactions
* 3- verify balance has changed as expected
*
*
**/
