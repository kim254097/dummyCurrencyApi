package db

type Currency struct {
	ID      string `json:"ID"`
	Name    string `json:"Name"`
	Country string `json:"Country"`
}

type AllCurrency []Currency
