package models

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

/**
* get current balances for a single account
**/
/*func GetBalanceForAccountId(accountId int) (float32, float32, float32) {
	log.Println("getBalanceForAccountId", accountId)
	db, e := myDb.setup()
	defer db.Close()
	if e != nil {
		fmt.Print(e)
	}
	var paymentBalance float32
	var paymentBalanceAfterTax float32
	var productBalance float32
	selectStatement := "SELECT"
	selectStatement += "(SELECT IFNULL(sum(amount), 0.00) AS \"paymentBalance\" "
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

	return paymentBalance, paymentBalanceAfterTax, productBalance
}*/

/**
* get current balances combined for account holder
**/
func GetBalanceForAccountholderId(accountHolderId int) (float32, float32, float32) {
	log.Println("GetBalanceForAccountholderId", accountHolderId)
	db, e := myDb.setup()
	defer db.Close()
	if e != nil {
		fmt.Print(e)
	}
	var tax float32
	var revenue float32
	var productSales float32
	selectStatement := "SELECT"
	selectStatement += "(SELECT IFNULL(sum(amount), 0.00) AS \"tax\" "
	selectStatement += "FROM `transactions` AS t "
	selectStatement += "JOIN accounts AS a ON a.id = t.accountId "
	//	selectStatement += "JOIN accountHolders AS ah ON ah.id = a.accountHolderId "
	//	selectStatement += "JOIN taxRates AS tr ON a.`taxRateId` = tr.id "
	selectStatement += "WHERE a.accountHolderId =  ? AND  a.type = \"revenue\"), "

	selectStatement += "(SELECT IFNULL(sum(amount), 0.00) AS \"revenue\" "
	selectStatement += "FROM `transactions` AS t "
	selectStatement += "JOIN accounts AS a ON a.id = t.accountId "
	//	selectStatement += "JOIN taxRates AS tr ON a.`taxRateId` = tr.id "
	selectStatement += "WHERE a.accountHolderId =  ? AND  a.type = \"tax\"), "

	selectStatement += "(SELECT IFNULL(sum(amount), 0.00) AS \"productSales\" "
	selectStatement += "FROM `transactions` AS t "
	selectStatement += "JOIN accounts AS a ON a.id = t.accountId "
	selectStatement += "WHERE a.accountHolderId =  ? AND  a.type = \"product\") AS \"productSales\""

	err := db.QueryRow(selectStatement, accountHolderId, accountHolderId, accountHolderId).Scan(&revenue, &tax, &productSales)
	if err != nil {
		fmt.Print(err)
	}

	return revenue, tax, productSales
}

/**
* get current balances combined for account holder
**/
func GetBalanceAcrossCompany() (float32, float32, float32) {
	log.Println("GetBalanceAcrossComapny")
	db, e := myDb.setup()
	defer db.Close()
	if e != nil {
		fmt.Print(e)
	}
	var paymentBalance float32
	var paymentBalanceAfterTax float32
	var productBalance float32
	selectStatement := "SELECT"
	selectStatement += "(SELECT IFNULL(sum(amount), 0.00) AS \"paymentBalance\" "
	selectStatement += "FROM `transactions` AS t "
	selectStatement += "JOIN accounts AS a ON a.id = t.accountId "
	selectStatement += "JOIN taxRates AS tr ON a.`taxRateId` = tr.id "
	selectStatement += "WHERE paymentOrProduct = \"payment\"), "

	selectStatement += "(SELECT IFNULL(sum(amount - (amount * tr.taxRate)), 0.00) AS \"paymentBalanceAfterTax\" "
	selectStatement += "FROM `transactions` AS t "
	selectStatement += "JOIN accounts AS a ON a.id = t.accountId "
	selectStatement += "JOIN taxRates AS tr ON a.`taxRateId` = tr.id "
	selectStatement += "WHERE paymentOrProduct = \"payment\"), "

	selectStatement += "(SELECT IFNULL(sum(amount), 0.00) AS \"accountAmount\" "
	selectStatement += "FROM `transactions` AS t "
	selectStatement += "JOIN accounts AS a ON a.id = t.accountId "
	selectStatement += "WHERE paymentOrProduct = \"product\") AS \"productBalance\""

	err := db.QueryRow(selectStatement).Scan(&paymentBalance, &paymentBalanceAfterTax, &productBalance)
	if err != nil {
		fmt.Print(err)
	}

	return paymentBalance, paymentBalanceAfterTax, productBalance
}
