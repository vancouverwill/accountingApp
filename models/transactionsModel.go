package models

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

type Transaction struct {
	Id        int       `json:"id"`
	AccountId int       `json:"accountId"`
	Details   string    `json:"details"`
	Amount    float32   `json:"amount"` // saved as US dollars
	Date      time.Time `json:"date"`
	Updated   int       `json:"updated"`
	Created   int       `json:"created"`
}

type TransactionViewable struct {
	Transaction
	AccountType string `json:"paymentOrProduct"`
}

type TransactionModel struct {
}

type TransactionViewables []TransactionViewable
type Transactions []Transaction

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
	res, err := stmt.Exec(t.AccountId, t.Details, amountInUS, t.Date)
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
func GetTransactions() TransactionViewables {
	log.Println("GetTransactions")
	db, e := myDb.setup()
	defer db.Close()

	if e != nil {
		fmt.Print(e)
	}

	selectStatement := "SELECT t.id, t.accountId, t.details, t.amount, t.date, a.type AS accountType FROM transactions AS t "
	selectStatement += "JOIN accounts AS a ON a.id = t.accountId"

	rows, err := db.Query(selectStatement)
	if err != nil {
		fmt.Print(err)
	}

	var results = make([]TransactionViewable, 0)

	i := 0
	for rows.Next() {

		var (
			id          int
			accountId   int
			details     string
			amount      float32
			accountType string
			date        string
		)
		var err = rows.Scan(&id, &accountId, &details, &amount, &date, &accountType)

		layout := "2006-01-02"

		dateString, err := time.Parse(layout, date)
		if err != nil {
			fmt.Println(err)
		}
		transaction := TransactionViewable{Transaction{Id: id, AccountId: accountId, Details: details, Amount: amount, Date: dateString}, accountType}
		results = append(results, transaction)
		i++
	}
	//	log.Println(results)

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
		id        int
		accountId int
		details   string
		amount    float32
		date      time.Time
	)
	err := db.QueryRow("SELECT id, accountId, details, amount, date FROM transactions WHERE id = ?", transactionId).Scan(&id, &accountId, &details, &amount, &date)
	if err != nil {
		fmt.Print(err)
	}

	transaction := Transaction{Id: id, AccountId: accountId, Details: details, Amount: amount, Date: date}

	//	log.Println(transaction)
	return transaction
}

/**
*
* get transaction by account id
*
**/
/*func GetTransactionsForAccountId(accountId int) Transactions {
	log.Println("GetTransactionsForAccountId")
	db, e := myDb.setup()
	defer db.Close()

	if e != nil {
		fmt.Print(e)
	}

	rows, err := db.Query("SELECT id, accountId, details, paymentOrProduct, amount, date FROM transactions WHERE accountId = ?", accountId)
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
}*/

/**
*
* get transaction by account holde id
*
**/
func GetTransactionsForAccountHolderId(accountHolderId int) TransactionViewables {
	log.Println("GetTransactionsForAccountHolderId", accountHolderId)

	db, e := myDb.setup()
	defer db.Close()

	if e != nil {
		fmt.Print(e)
	}

	selectStatement := "SELECT t.id, t.accountId, t.details, a.type AS accountType, t.amount, t.date "
	selectStatement += "FROM transactions AS t "
	selectStatement += "JOIN accounts AS a ON a.id = t.accountId "
	selectStatement += "WHERE a.accountHolderId = ? "

	rows, err := db.Query(selectStatement, accountHolderId)
	if err != nil {
		fmt.Print(err)
	}

	var results = make([]TransactionViewable, 0)

	i := 0
	for rows.Next() {

		var (
			id          int
			accountId   int
			accountType string
			details     string
			amount      float32
			date        string
		)
		var err = rows.Scan(&id, &accountId, &details, &accountType, &amount, &date)

		layout := "2006-01-02"

		dateString, err := time.Parse(layout, date)
		if err != nil {
			fmt.Println(err)
		}
		transaction := Transaction{Id: id, AccountId: accountId, Details: details, Amount: amount, Date: dateString}
		transactionViewable := TransactionViewable{transaction, accountType}
		results = append(results, transactionViewable)
		i++
	}
	log.Println(results)

	return results
}

//deleteTransaction(transaction_id int)
