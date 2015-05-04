package models

import (
	//	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func GetBalanceForAccountId(relatedToIdInt int) (float32, float32) {
	log.Println("getBalanceForAccountId", relatedToIdInt)
	db, e := myDb.setup()
	defer db.Close()
	if e != nil {
		fmt.Print(e)
	}
	var paymentBalance float32
	var productBalance float32
	selectStatement := "select (select IFNULL(sum(amount), 0.00) AS \"account_amount\" from `transactions` where accountId =  ? AND  paymentOrProduct = \"payment\") AS \"paymentBalance\", (select IFNULL(sum(amount), 0.00) AS \"account_amount\" from `transactions` where accountId =  ? AND  paymentOrProduct = \"product\") AS \"productBalance\""
	err := db.QueryRow(selectStatement, relatedToIdInt, relatedToIdInt).Scan(&paymentBalance, &productBalance)
	if err != nil {
		fmt.Print(err)
	}

	log.Println("paymentBalance", paymentBalance)
	log.Println("productBalance", productBalance)

	//	select (select sum(amount) AS "account_amount" from `transactions` where accountId =  9
	//AND  paymentOrProduct = "payment") AS "paymentBalance",

	//(select sum(amount) AS "account_amount" from `transactions` where accountId =  9
	//AND  paymentOrProduct = "product") AS "productBalance";

	return paymentBalance, productBalance
}
