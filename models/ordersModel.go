package models

type Order struct {
	Id              int
	Name            string
	TaxRate         float32
	AccountHolderId int
//	RevenueAccountId int
//	TaxAccountId int
//	ProductSalesAccountId int
	CurrencyId      int
	TaxRateId       int
	Cost            float32
}

func (o *Order) addItem(name String, cost float32) {
	o.Name = name
	o.Cost = cost
}

func (o *Order) processProduct()
{
	productTransaction = Transaction{}
	
	saveTransactionByType(o.AccountHolderId, int, "product", -o.Cost)
}
