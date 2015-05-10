package models

import ()

type Order struct {
	Id              int
	Name            string
	TaxRate         float32
	AccountHolderId int
	CurrencyId      int
	TaxRateId       int
	Amount          float32
	RevenueMadeInUs float32
	TaxPayable      float32
	PaymentPaid     float32
}

func (o *Order) AddItem(name string, amount float32) {
	o.Name = name
	o.Amount = amount
}

func (o *Order) PrepareRevenue() {

	currency := GetCurrencyByAccountHolderId(o.AccountHolderId)
	o.RevenueMadeInUs = o.Amount * currency.ExchangeRate

	taxRate := GetTaxRateByAccountHolderId(o.AccountHolderId)
	o.TaxPayable = o.RevenueMadeInUs * (taxRate.TaxRate)
}

/**
*
* When we take money out of the payment account
*
**/
func (o *Order) PreparePayment() {
	o.PaymentPaid = -o.Amount - o.TaxPayable
}

func (o *Order) FinalizeOrder() {
	SaveTransactionByType(o.AccountHolderId, "revenue", o.RevenueMadeInUs, o.Name+" revenue")
	SaveTransactionByType(o.AccountHolderId, "tax", o.TaxPayable, o.Name+" tax")
	SaveTransactionByType(o.AccountHolderId, "payment", o.PaymentPaid, o.Name)
}
