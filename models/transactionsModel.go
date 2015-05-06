package models

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

type Transaction struct {
	Id               int       `json:"id"`
	AccountId        int       `json:"accountId"`
	Details          string    `json:"details"`
	PaymentOrProduct string    `json:"paymentOrProduct"`
	Amount           float32   `json:"amount"` // saved as US dollars
	Date             time.Time `json:"date"`
	Updated          int       `json:"updated"`
	Created          int       `json:"created"`
}

type TransactionModel struct {
}

type Transactions []Transaction

var dbName string = "transactions"

var MyTransaction = TransactionModel{}

/**
*
* amount is recorded as US dollars
*
**/
func (tm TransactionModel) SaveTransaction(t Transaction) Transaction {
	log.Println("RepoCreateTransaction")
	log.Println(t)

	db, e := myDb.setup()
	defer db.Close()

	if e != nil {
		fmt.Print(e)
	}

	currency := getCurrencyByAccountId(t.AccountId)

	log.Println("currency", currency)

	amountInUS := t.Amount * currency.ExchangeRate

	stmt, err := db.Prepare("INSERT INTO transactions (accountId, details, paymentOrProduct, amount, date, updated, created) values (?, ?, ?, ?, ?, UNIX_TIMESTAMP(), UNIX_TIMESTAMP())")
	if err != nil {
		fmt.Print(err)
	}
	res, err := stmt.Exec(t.AccountId, t.Details, t.PaymentOrProduct, amountInUS, t.Date)
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
	t.Id = int(lastId)
	log.Println("transaction entered")
	log.Println(t)
	return t
}

/**
*
* get all transactions
*
**/
func GetTransactions() Transactions {
	log.Println("GetTransactions")
	db, e := myDb.setup()
	defer db.Close()

	if e != nil {
		fmt.Print(e)
	}

	rows, err := db.Query("select id, accountId, details, paymentOrProduct, amount, date from " + dbName)
	if err != nil {
		fmt.Print(err)
	}

	var results = make([]Transaction, 0)

	i := 0
	for rows.Next() {

		var (
			id               int
			accountId        int
			details          string
			paymentOrProduct string
			amount           float32
			date             string
		)
		var err = rows.Scan(&id, &accountId, &details, &paymentOrProduct, &amount, &date)

		layout := "2006-01-02"

		dateString, err := time.Parse(layout, date)
		if err != nil {
			fmt.Println(err)
		}
		transaction := Transaction{Id: id, AccountId: accountId, Details: details, PaymentOrProduct: paymentOrProduct, Amount: amount, Date: dateString}
		results = append(results, transaction)
		i++
	}
	log.Println(results)

	return results
}

/**
*
* get transaction by transaction id
*
**/
func GetTransaction(transactionId int) Transaction {
	db, e := myDb.setup()
	defer db.Close()
	if e != nil {
		fmt.Print(e)
	}
	var (
		id               int
		accountId        int
		details          string
		paymentOrProduct string
		amount           float32
		date             time.Time
	)
	err := db.QueryRow("SELECT id, accountId, details, paymentOrProduct, amount, date FROM "+dbName+" WHERE id = ?", transactionId).Scan(&id, &accountId, &details, &paymentOrProduct, &amount, &date)
	if err != nil {
		fmt.Print(err)
	}

	transaction := Transaction{Id: id, AccountId: accountId, Details: details, PaymentOrProduct: paymentOrProduct, Amount: amount, Date: date}

	log.Println(transaction)
	return transaction
}

/**
*
* get transaction by account id
*
**/
func GetTransactionsForAccountId(accountId int) Transactions {
	log.Println("GetTransactionsForAccountId")
	db, e := myDb.setup()
	defer db.Close()

	if e != nil {
		fmt.Print(e)
	}

	rows, err := db.Query("SELECT id, accountId, details, paymentOrProduct, amount, date FROM "+dbName+" WHERE accountId = ?", accountId)
	if err != nil {
		fmt.Print(err)
	}

	var results = make([]Transaction, 0)

	i := 0
	for rows.Next() {

		var (
			id               int
			accountId        int
			details          string
			paymentOrProduct string
			amount           float32
			date             string
		)
		var err = rows.Scan(&id, &accountId, &details, &paymentOrProduct, &amount, &date)

		layout := "2006-01-02"

		dateString, err := time.Parse(layout, date)
		if err != nil {
			fmt.Println(err)
		}
		transaction := Transaction{Id: id, AccountId: accountId, Details: details, PaymentOrProduct: paymentOrProduct, Amount: amount, Date: dateString}
		results = append(results, transaction)
		i++
	}
	log.Println(results)

	return results
}

/**
*
* get transaction by account holde id
*
**/
func GetTransactionsForAccountHolderId(accountHolderId int) Transactions {
	log.Println("GetTransactionsForAccountHolderId")

	db, e := myDb.setup()
	defer db.Close()

	if e != nil {
		fmt.Print(e)
	}

	rows, err := db.Query("SELECT t.id, t.accountId, t.details, t.paymentOrProduct, t.amount, t.date from transactions AS t JOIN accounts AS a ON a.id = t.accountId WHERE a.accountHolderId = ?", accountHolderId)
	if err != nil {
		fmt.Print(err)
	}

	var results = make([]Transaction, 0)

	i := 0
	for rows.Next() {

		var (
			id               int
			accountId        int
			details          string
			paymentOrProduct string
			amount           float32
			date             string
		)
		var err = rows.Scan(&id, &accountId, &details, &paymentOrProduct, &amount, &date)

		layout := "2006-01-02"

		dateString, err := time.Parse(layout, date)
		if err != nil {
			fmt.Println(err)
		}
		transaction := Transaction{Id: id, AccountId: accountId, Details: details, PaymentOrProduct: paymentOrProduct, Amount: amount, Date: dateString}
		results = append(results, transaction)
		i++
	}
	log.Println(results)

	return results
}

//deleteTransaction(transaction_id int)
