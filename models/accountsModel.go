package models

type Account struct {
	Id              int    `json:"id"`
	AccountHolderId int    `json:"accountHolderId"`
	Name            string `json:"name"`
	Address         string `json:"address"`
	CurrencyId      int    `json:"currencyId"`
	TaxRateId       int    `json:"taxRateId"`
	Updated         int    `json:"updated"`
	Created         int    `json:"created"`
}

/**
*
* todo
*
**/
// createAccount(name String, address String)

/**
*
* todo
*
**/
// deleteAccount(accountId int)

/**
*
* todo
*
**/
// updateAccount()
