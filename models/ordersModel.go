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
	AmountInUs      float32
	RevenueMadeInUs float32
	TaxPayable      float32
	ComissionPaid   float32
	PaymentPaid     float32
}

func (o *Order) AddItem(name string, amount float32) {
	o.Name = name
	o.Amount = amount
}

func (o *Order) PrepareRevenue() {
	var comissionRate float32 = 0.1
	currency := GetCurrencyByAccountHolderId(o.AccountHolderId)

	o.AmountInUs = o.Amount * currency.ExchangeRate

	o.ComissionPaid = o.AmountInUs * comissionRate
	o.RevenueMadeInUs = o.AmountInUs - o.ComissionPaid

	taxRate := GetTaxRateByAccountHolderId(o.AccountHolderId)
	o.TaxPayable = o.AmountInUs * (taxRate.TaxRate)
}

/**
*
* When we take money out of the payment account
*
**/
func (o *Order) PreparePayment() {
	o.PaymentPaid = -o.AmountInUs - o.TaxPayable
}

func (o *Order) FinalizeOrder() {
	SaveTransactionByType(o.AccountHolderId, "revenue", o.RevenueMadeInUs, o.Name+" revenue")
	SaveTransactionByType(o.AccountHolderId, "tax", o.TaxPayable, o.Name+" tax")
	SaveTransactionByType(o.AccountHolderId, "commission", o.ComissionPaid, o.Name+" commission")
	SaveTransactionByType(o.AccountHolderId, "payment", o.PaymentPaid, o.Name+" payment")
}
