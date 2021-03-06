package data

import (
	"time"
	"io"
	"encoding/json"
)

type Product struct {
	ID					int `json:"id"`
	Name				string `json:"name"`
	Description	string `json:"description"`
	Price				float32 `json:"price"`
	SKU					string `json:"sku"`
	CreatedOn		string `json:"-"`
	UpdatedOn		string `json:"-"`
	DeletedOn		string `json:"-"`
}

type Products []*Product

func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func GetProducts() Products {
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