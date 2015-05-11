package main

import (
	"github.com/vancouverwill/accountingApp/models"
	"log"
	"testing"
)

func TestCreateOrder(t *testing.T) {
	accountHolder := models.PrepareAccountForTesting("Lucas Brasi", "Developer")
	log.Println(accountHolder)
	order := accountHolder.NewOrder()
	order.AddItem("Sony Playstation 3", 800)
	order.PrepareRevenue()
	order.PreparePayment()
	order.FinalizeOrder()
}
