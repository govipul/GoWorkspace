package model

// Customer data type
type Customer struct {
	//ID of the customer
	ID int
	//Name of the customer
	Name        string
	addressData address
	married     bool
}
