package models

import (
	"fmt"
	"log"
)

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

/**
*
* todo
*
**/
func GetAccountByName(accountName string) Account {
	db, e := myDb.setup()
	defer db.Close()
	if e != nil {
		fmt.Print(e)
	}
	var (
		id              int
		accountHolderId int
		name            string
		address         string
		currencyId      int
		taxRateId       int
	)
	err := db.QueryRow("SELECT id, accountHolderId, name, address, currencyId, taxRateId FROM accounts WHERE name = ?", accountName).Scan(&id, &accountHolderId, &name, &address, &currencyId, &taxRateId)
	if err != nil {
		fmt.Print(err)
	}

	log.Println("getAccountByName", id)

	account := Account{Id: id, AccountHolderId: accountHolderId, Name: name, Address: address, CurrencyId: currencyId, TaxRateId: taxRateId}

	log.Println(account)
	return account
}
