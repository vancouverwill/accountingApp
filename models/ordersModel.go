package models

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
	Cost       float32
}

func (o *Order) addItem(name string, cost float32) {
	o.Name = name
	o.Cost = cost
}

func (o *Order) ProcessProduct() {
	amount := -o.Cost
	SaveTransactionByType(o.AccountHolderId, "product", amount, o.Name)
}

func (o *Order) ProcessPayment() {

	currency := models.GetCurrencyByAccountId(o.AccountHolderId)

	log.Println("currency", currency)

	amountInUS := o.Amount * currency.ExchangeRate

	taxRate = GetTaxByAccountId(o.AccountHolderId)

	amountInUS *

		SaveTransactionByType(o.AccountHolderId, "payment", amount, o.Name)
}
