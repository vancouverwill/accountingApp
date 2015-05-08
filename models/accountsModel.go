package models

type Account struct {
	Id              int    `json:"id"`
	AccountHolderId int    `json:"accountHolderId"`
	AccountType     string `json:"accountType"`
	Updated         int    `json:"updated"`
	Created         int    `json:"created"`
}

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
