package models

type TaxRate struct {
	Id      int     `json:"id"`
	Name    string  `json:"name"`
	TaxRate float32 `json:"taxRate"` // e.g. 0.13 is equal to 13 percent
	Updated int     `json:"updated"`
	Created int     `json:"created"`
}
