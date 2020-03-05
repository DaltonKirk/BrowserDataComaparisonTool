package main

type BrowserData struct {
	ID                 int
	Date               string
	ClientID           string
	DeviceCategory     string
	Browser            string
	BrowserVersion     string
	Sessions           string
	Transactions       string
	TransactionRevenue string
}

type ClientModel struct {
	ClientID string
	Browser  string
}

type CSVRow struct {
	Title string
	Count string
}
