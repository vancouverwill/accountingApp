package models

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

/**
* get current balances for a single account
**/
func GetBalanceForAccountId(accountId int) (float32, float32, float32) {
	log.Println("getBalanceForAccountId", accountId)
	db, e := myDb.setup()
	defer db.Close()
	if e != nil {
		fmt.Print(e)
	}
	var paymentBalance float32
	var paymentBalanceAfterTax float32
	var productBalance float32
	selectStatement := "SELECT (select IFNULL(sum(amount), 0.00) AS \"paymentBalance\" "
	selectStatement += "FROM `transactions` AS t "
	selectStatement += "JOIN accounts AS a ON a.id = t.accountId "
	selectStatement += "JOIN taxRates AS tr ON a.`taxRateId` = tr.id "
	selectStatement += "WHERE t.accountId =  ? AND  paymentOrProduct = \"payment\"), "
	selectStatement += "(SELECT IFNULL(sum(amount - (amount * tr.taxRate)), 0.00) AS \"paymentBalanceAfterTax\" "
	selectStatement += "FROM `transactions` AS t "
	selectStatement += "JOIN accounts AS a ON a.id = t.accountId "
	selectStatement += "JOIN taxRates AS tr ON a.`taxRateId` = tr.id "
	selectStatement += "WHERE t.accountId =  ? AND  paymentOrProduct = \"payment\"), "
	selectStatement += "(SELECT IFNULL(sum(amount), 0.00) AS \"accountAmount\" "
	selectStatement += "FROM `transactions` AS t "
	selectStatement += "WHERE accountId =  ? AND  paymentOrProduct = \"product\") AS \"productBalance\""

	err := db.QueryRow(selectStatement, accountId, accountId, accountId).Scan(&paymentBalance, &paymentBalanceAfterTax, &productBalance)
	if err != nil {
		fmt.Print(err)
	}

	log.Println("getBalanceForAccountId", "finished")

	return paymentBalance, paymentBalanceAfterTax, productBalance
}
