package models

import (
	"fmt"
	"log"
)

type AccountHolder struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	JobTitle   string `json:"jobTitle"`
	CurrencyId int    `json:"currencyId"`
	TaxRateId  int    `json:"taxRateId"`
	Updated    int    `json:"updated"`
	Created    int    `json:"created"`
}

/**
*
* todo
*
**/
func CreateAccountHolder(name string, jobTitle string) AccountHolder {
	log.Println("createAccount")

	accountHolder := AccountHolder{Name: name, JobTitle: jobTitle}

	return accountHolder

}

func (ah *AccountHolder) SetAccountHolderCurrency(currencyName string) {
	log.Println("SetAccountHolderCurrency")
	currency := getCurrencyByCurrencyName(currencyName)
	ah.CurrencyId = currency.Id
	log.Println(ah)
}

func (ah *AccountHolder) SetAccountHolderTaxRate(name string) {
	log.Println("SetAccountHolderTaxRate")
	taxRate := getTaxRateByName(name)
	ah.TaxRateId = taxRate.Id
	log.Println(ah)
}

func SaveAccount(accountHolder AccountHolder) AccountHolder {
	db, e := myDb.setup()
	defer db.Close()

	if e != nil {
		fmt.Print(e)
	}

	stmt, err := db.Prepare("INSERT INTO accounHolders (name, jobTitle, taxRateId, currencyId) values (?, ?, ?, ?, ?, UNIX_TIMESTAMP(), UNIX_TIMESTAMP())")
	if err != nil {
		fmt.Print(err)
	}
	res, err := stmt.Exec(accountHolder.Name, accountHolder.JobTitle, accountHolder.TaxRateId, accountHolder.CurrencyId)
	if err != nil {
		log.Fatal(err)
	}

	lastId, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	RowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("RowsAffected", RowsAffected)
	accountHolder.Id = int(lastId)
	log.Println("AccountHolder entered")
	return accountHolder
}

/**
*
* todo
*
**/
//removeAccount

/**
*
* todo
*
**/
//listAccounts

/**
*
* todo
*
**/
//getAccountHolderDetails

/**
*
* todo
*
**/
func GetAccountHolderByName(accountName string) AccountHolder {
	db, e := myDb.setup()
	defer db.Close()
	if e != nil {
		fmt.Print(e)
	}
	var (
		id         int
		name       string
		jobTitle   string
		currencyId int
		taxRateId  int
	)
	err := db.QueryRow("SELECT id, name, jobTitle, currencyId, taxRateId FROM accountHolders WHERE name = ?", accountName).Scan(&id, &name, &jobTitle, &currencyId, &taxRateId)
	if err != nil {
		fmt.Print(err)
	}

	log.Println("getAccountByName", id)

	accountHolder := AccountHolder{Id: id, Name: name, JobTitle: jobTitle, CurrencyId: currencyId, TaxRateId: taxRateId}

	log.Println(accountHolder)
	return accountHolder
}
