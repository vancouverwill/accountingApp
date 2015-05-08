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

func GetCurrencyByAccountId(accountId int) Currency {
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
	selectStatement := "SELECT c.id, c.name, c.exchangeRate "
	selectStatement += "FROM currencies AS c "
	selectStatement += "JOIN accountHolders AS ah ON ah.currencyId = c.id "
	selectStatement += "JOIN accounts AS a ON a.accountHolderId = ah.id "
	selectStatement += "WHERE a.id = ?"

	err := db.QueryRow(selectStatement, accountId).Scan(&id, &name, &exchangeRate)
	if err != nil {
		fmt.Print(err)
	}

	currency := Currency{Id: id, Name: name, ExchangeRate: exchangeRate}

	log.Println(currency)
	return currency
}

func GetCurrencyByCurrencyName(currencyName string) Currency {
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
