package models

import (
	"time"
)

//create_transaction(account_id int, payment_or_product string, amount float)

type Transaction struct {
	Id               int       `json:"id"`
	AccountId        int       `json:"accountId"`
	Details          string    `json:"details"`
	PaymentOrProduct string    `json:"completed"`
	Amount           float32   `json:"amount"`
	Date             time.Time `json:"due"`
}

/**
*
*
**/
//getTransactions(start_date ???, end_date ???, AccountAccountHolderOrCompany, relatedId Int null) json

//deleteTransaction(transaction_id int)
