package models

import (
	"log"
)

type Order struct {
	Id              int
	Name            string
	TaxRate         float32
	AccountHolderId int
	//	RevenueAccountId int
	//	TaxAccountId int
	//	ProductSalesAccountId int
	CurrencyId int
	TaxRateId  int
	Amount     float32
}

func (o *Order) addItem(name string, amount float32) {
	o.Name = name
	o.Amount = amount
}

func (o *Order) ProcessProduct() {
	amount := -o.Amount
	SaveTransactionByType(o.AccountHolderId, "product", amount, o.Name)
}

func (o *Order) ProcessPayment() {

	currency := GetCurrencyByAccountId(o.AccountHolderId)

	log.Println("currency", currency)

	amountInUS := o.Amount * currency.ExchangeRate

	taxRate := GetTaxRateByAccountId(o.AccountHolderId)

	revenueMade := amountInUS * (1 - taxRate.TaxRate)
	taxPayable := amountInUS * (taxRate.TaxRate)

	SaveTransactionByType(o.AccountHolderId, "payment", revenueMade, o.Name+" Payment")
	SaveTransactionByType(o.AccountHolderId, "tax", taxPayable, o.Name+" tax")
}
