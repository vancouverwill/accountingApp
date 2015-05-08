package models

import (
	"fmt"
	"log"
)

type Currency struct {
	Id           int     `json:"id"`
	Name         string  `json:"name"`
	ExchangeRate float32 `json:"exchangeRate"` // base currency is US dollar
	Updated      int     `json:"updated"`
	Created      int     `json:"created"`
}

/**
*
* return json exp [{"name" : "xxxx", "exchangeRate" : 2.00},....]
*
**/
// listCurrencies()

// createCurrency(name String, exchangeRate float)

// updateCurrencyExchangeRate(currencyID int, exchangeRate float)

// deleteCurrency()

func getCurrencyByAccountId(accountId int) Currency {
	db, e := myDb.setup()
	defer db.Close()
	if e != nil {
		fmt.Print(e)
	}
	var (
		id           int
		name         string
		exchangeRate float32
	)
	err := db.QueryRow("SELECT c.id, c.name, c.exchangeRate FROM "+"currencies"+" AS c JOIN accounts AS a ON a.currencyId = c.id WHERE a.id = ?", accountId).Scan(&id, &name, &exchangeRate)
	if err != nil {
		fmt.Print(err)
	}

	currency := Currency{Id: id, Name: name, ExchangeRate: exchangeRate}

	log.Println(currency)
	return currency
}

func getCurrencyByCurrencyName(currencyName string) Currency {
	log.Println("getCurrencyByCurrencyName")
	db, e := myDb.setup()
	defer db.Close()
	if e != nil {
		fmt.Print(e)
	}
	var (
		id           int
		name         string
		exchangeRate float32
	)
	err := db.QueryRow("SELECT c.id, c.name, c.exchangeRate FROM currencies AS c WHERE c.name = ?", currencyName).Scan(&id, &name, &exchangeRate)
	if err != nil {
		fmt.Print(err)
	}

	currency := Currency{Id: id, Name: name, ExchangeRate: exchangeRate}

	return currency
}
