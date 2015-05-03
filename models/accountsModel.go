package models

type Account struct {
	Id            int    `json:"id"`
	AccountId     int    `json:"accountId"`
	Name          string `json:"name"`
	Address       string `json:"address"`
	CurrentCredit int    `json:"currentCredit"`
	CurrentDebit  int    `json:"currentDebit"`
	Currency      int    `json:"currency"`
	Updated       int    `json:"updated"`
	Created       int    `json:"created"`
}

// createAccount(name String, address String)
// deleteAccount(accountId int)
// updateAccount()
