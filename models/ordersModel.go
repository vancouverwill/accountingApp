package models

import (
	"log"
)

type Order struct {
	Id              int
	Name            string
	TaxRate         float32
	AccountHolderId int
	CurrencyId      int
	TaxRateId       int
	Amount          float32
}

func (o *Order) AddItem(name string, amount float32) {
	o.Name = name
	o.Amount = amount
}

/**
*
* When we take money out of the payment account
*
**/
func (o *Order) ProcessPayment() {
	amount := -o.Amount
	SaveTransactionByType(o.AccountHolderId, "payment", amount, o.Name)
}

func (o *Order) ProcessRevenue() {

	currency := GetCurrencyByAccountHolderId(o.AccountHolderId)

	log.Println("ProcessRevene currency", currency)
	log.Println("ProcessPayment", o.AccountHolderId)

	amountInUS := o.Amount * currency.ExchangeRate

	taxRate := GetTaxRateByAccountHolderId(o.AccountHolderId)

	revenueMade := amountInUS
	taxPayable := amountInUS * (taxRate.TaxRate)

	SaveTransactionByType(o.AccountHolderId, "revenue", revenueMade, o.Name+" revenue")
	SaveTransactionByType(o.AccountHolderId, "tax", taxPayable, o.Name+" tax")
}
