package models

type Account struct {
	Id              int    `json:"id"`
	AccountHolderId int    `json:"accountHolderId"`
	Name            string `json:"name"`
	Address         string `json:"address"`
	CurrentCredit   int    `json:"currentCredit"`
	CurrentDebit    int    `json:"currentDebit"`
	CurrencyId      int    `json:"currencyId"`
	Updated         int    `json:"updated"`
	Created         int    `json:"created"`
}

// createAccount(name String, address String)
// deleteAccount(accountId int)
// updateAccount()
