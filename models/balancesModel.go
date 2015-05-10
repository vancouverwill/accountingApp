package models

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

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
	var payment float32
	selectStatement := "SELECT"
	selectStatement += "(SELECT IFNULL(sum(amount), 0.00) AS \"revenue\" "
	selectStatement += "FROM `transactions` AS t "
	selectStatement += "JOIN accountTypes AS at ON at.id = t.accountTypeId "
	//		selectStatement += "JOIN accountHolders AS ah ON ah.id = a.accountHolderId "
	//	selectStatement += "JOIN taxRates AS tr ON a.`taxRateId` = tr.id "
	selectStatement += "WHERE t.accountHolderId =  ? AND  at.type = \"revenue\"), "

	selectStatement += "(SELECT IFNULL(sum(amount), 0.00) AS \"tax\" "
	selectStatement += "FROM `transactions` AS t "
	selectStatement += "JOIN accountTypes AS at ON at.id = t.accountTypeId "
	//	selectStatement += "JOIN taxRates AS tr ON a.`taxRateId` = tr.id "
	selectStatement += "WHERE t.accountHolderId =  ? AND  at.type = \"tax\"), "

	selectStatement += "(SELECT IFNULL(sum(amount), 0.00) AS \"payment\" "
	selectStatement += "FROM `transactions` AS t "
	selectStatement += "JOIN accountTypes AS at ON at.id = t.accountTypeId "
	selectStatement += "WHERE t.accountHolderId =  ? AND  at.type = \"payment\") AS \"payment\""

	err := db.QueryRow(selectStatement, accountHolderId, accountHolderId, accountHolderId).Scan(&revenue, &tax, &payment)
	if err != nil {
		fmt.Print(err)
	}

	log.Println()

	return revenue, tax, payment
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
	var tax float32
	var revenue float32
	var payment float32
	selectStatement := "SELECT"
	selectStatement += "(SELECT IFNULL(sum(amount), 0.00) AS \"paymentBalance\" "
	selectStatement += "FROM `transactions` AS t "
	selectStatement += "JOIN accountTypes AS at ON at.id = t.accountTypeId "
	//	selectStatement += "JOIN taxRates AS tr ON a.`taxRateId` = tr.id "
	selectStatement += "WHERE at.type = \"revenue\"), "

	selectStatement += "(SELECT IFNULL(sum(amount), 0.00) AS \"tax\" "
	selectStatement += "FROM `transactions` AS t "
	selectStatement += "JOIN accountTypes AS at ON at.id = t.accountTypeId "
	//	selectStatement += "JOIN taxRates AS tr ON a.`taxRateId` = tr.id "
	selectStatement += "WHERE at.type = \"tax\"), "

	selectStatement += "(SELECT IFNULL(sum(amount), 0.00) AS \"accountAmount\" "
	selectStatement += "FROM `transactions` AS t "
	selectStatement += "JOIN accountTypes AS at ON at.id = t.accountTypeId "
	selectStatement += "WHERE at.type = \"payment\") AS \"payment\""

	err := db.QueryRow(selectStatement).Scan(&revenue, &tax, &payment)
	if err != nil {
		fmt.Print(err)
	}

	return revenue, tax, payment
}
