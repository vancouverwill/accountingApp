package models

import (
	//	"database/sql"
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
	Amount           float32   `json:"amount"`
	Date             time.Time `json:"date"`
	Updated          int       `json:"updated"`
	Created          int       `json:"created"`
}

type TransactionModel struct {
}

type Transactions []Transaction

var dbName string = "transactions"

var MyTransaction = TransactionModel{}

//create_transaction(account_id int, payment_or_product string, amount float)
func (tm TransactionModel) RepoCreateTransaction(t Transaction) Transaction {
	log.Println("RepoCreateTransaction")
	log.Println(t)

	db, e := myDb.setup()

	if e != nil {
		fmt.Print(e)
	}

	stmt, err := db.Prepare("INSERT INTO transactions (accountId, details, paymentOrProduct, amount, date, updated, created) values (?, ?, ?, ?, ?, UNIX_TIMESTAMP(), UNIX_TIMESTAMP())")
	if err != nil {
		fmt.Print(err)
	}
	res, err := stmt.Exec(t.AccountId, t.Details, t.PaymentOrProduct, t.Amount, t.Date)
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
*
**/
//func getTransactions(start_date ???, end_date ???, AccountAccountHolderOrCompany, relatedId Int null) Transactions
func GetTransactions() Transactions {
	db, e := myDb.setup()

	if e != nil {
		fmt.Print(e)
	}

	rows, err := db.Query("select id, accountId, details, paymentOrProduct, amount, date from " + dbName)
	if err != nil {
		fmt.Print(err)
	}

	//	var results = make([]Todo, 1000)

	var results Transactions

	if err != nil {
		fmt.Print(err)
	}
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
		log.Print(rows.Columns())
		log.Print(id)
		log.Print(details)

		layout := "2006-01-02"

		dateString, err := time.Parse(layout, date)
		if err != nil {
			fmt.Println(err)
		}
		//		todo := &Todo{Id: id, Name: name, Completed: completed}
		transaction := Transaction{Id: id, AccountId: accountId, Details: details, PaymentOrProduct: paymentOrProduct, Amount: amount, Date: dateString}
		//		b, err := json.Marshal(todo)
		//		if err != nil {
		//			fmt.Println(err)
		//			return
		//		}
		//		results[i] = fmt.Sprintf("%s", string(b))
		//		results[i] := todo
		results = append(results, transaction)
		i++
	}
	//	result = result[:i]

	log.Println(results)

	return results
}

//deleteTransaction(transaction_id int)
