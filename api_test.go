package main

import (
	"github.com/vancouverwill/accountingApp/controllers"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestIndex(t *testing.T) {
	t.Log("TestIndex")
	req, err := http.NewRequest("GET", "http://localhost:8080/", nil)
	if err != nil {
		t.Error("index() did not work as expected.")
		t.Log(err)
	}

	w := httptest.NewRecorder()
	controllers.Index(w, req)

	if w.Code != 200 {
		t.Error("index() did not work as expected. the status was not 200")
	}

	t.Log("status:", w.Code, "body:", w.Body.String())
}

/**
*
* test simple 404 error page to make sure it returns the correct status if a user is lost
*
**/
func Test404Page(t *testing.T) {
	resp := httptest.NewRecorder()

	req, err := http.NewRequest("GET", "http://localhost:8080/sdfsfd", nil)
	if err != nil {
		t.Fatal(err)
	}

	http.DefaultServeMux.ServeHTTP(resp, req)

	if resp.Code != 404 {
		t.Error("404 error is not being returned!")
	}
}

/**
*
* test the balance with a valid AccountHolderOrCompany and relatedToId
*
**/
func TestBalancesIndex(t *testing.T) {
	req, err := http.NewRequest("GET", "http://localhost:8080/balances/?AccountHolderOrCompany=AccountHolder&relatedToId=9", nil)
	if err != nil {
		t.Error("index() did not work as expected.")
		t.Log(err)
	}

	w := httptest.NewRecorder()
	controllers.BalancesIndex(w, req)

	if w.Code != 200 && w.Code != 202 {
		t.Error("BalancesIndex() did not work as expected. the status was not ", http.StatusOK, ", it was ", w.Code)
	}

	t.Log("status:", w.Code, "body:", w.Body.String())
}

/**
*
* test non valid entry of "country" for AccountHolderOrCompany paramater of GET
*
**/
func TestBalancesValidatesAccountHolderOrCompany(t *testing.T) {
	req, err := http.NewRequest("GET", "http://localhost:8080/balances/?AccountHolderOrCompany=Country&relatedToId=234", nil)
	if err != nil {
		t.Error("index() did not work as expected.")
		t.Log(err)
	}

	w := httptest.NewRecorder()
	controllers.BalancesIndex(w, req)

	if w.Code != 400 {
		t.Error("TransactionsIndex() did not work as expected. the status was not 400, it was ", w.Code)
	}
}

/**
*
* test GET all transactions from the company
*
**/
func TestTransactionsIndex(t *testing.T) {
	req, err := http.NewRequest("GET", "http://localhost:8080/transactions/?AccountHolderOrCompany=Company&relatedToId=0", nil)
	if err != nil {
		t.Error("index() did not work as expected.")
		t.Log(err)
	}

	w := httptest.NewRecorder()
	controllers.TransactionsIndex(w, req)

	if w.Code != 200 && w.Code != 202 {
		t.Error("TransactionsIndex() did not work as expected. the status was not ", http.StatusOK, ", it was ", w.Code)
	}

	t.Log("status:", w.Code, "body:", w.Body.String())
}

/**
*
* test POST transactions from the company
*
* todo working on accepting JSON body as part http.NewRequest
*
**/
/*func TestTransactionsCreate(t *testing.T) {

	var transaction models.Transaction
	transaction = models.Transaction{AccountId: 9, Details: "buying lots of products AGAIN", Amount: 201, Date: "2015-01-19T00:00:00Z", Updated: 0, Created: 0}

	b, err := json.Marshal(transaction)
	if err != nil {
		t.Error("error:", err)
	}
	req, err := http.NewRequest("POST", "http://localhost:8080/transactions/", bytes.NewBufferString(b))
	if err != nil {
		t.Error("index() did not work as expected.")
		t.Log(err)
	}

	w := httptest.NewRecorder()
	controllers.TransactionsCreate(w, req)

	if w.Code != 200 && w.Code != 202 {
		t.Error("TransactionsIndex() did not work as expected. the status was not ", http.StatusOK, ", it was ", w.Code)
	}

	t.Log("status:", w.Code, "body:", w.Body.String())
}*/
