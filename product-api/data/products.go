package data

import "time"

type Product struct {
	ID					int
	Name				string
	Description	string
	Price				float32
	SKU					string
	CreatedOn		string
	UpdatedOn		string
	DeletedOn		string
}

func GetProducts() []*Product {
	return productList
}

var productList = []*Product{
	&Product{
		ID: 1,
		Name: "latte",
		Description: "frothy milky coffee",
		Price: 2.45,
		SKU: "abc323",
		CreatedOn: time.Now().UTC().String(),
		UpdatedOn: time.Now().UTC().String(),
	},
	&Product{
		ID: 2,
		Name: "Espresso",
		Description: "short and strong coffee with no milk",
		Price: 1.99,
		SKU: "fjd34",
		CreatedOn: time.Now().UTC().String(),
		UpdatedOn: time.Now().UTC().String(),
	},
}