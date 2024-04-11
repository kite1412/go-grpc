package model

type Item struct {
	Name      string
	Additions []string
}

type Order struct {
	Username string
	Quantity int64
	Item     Item
}