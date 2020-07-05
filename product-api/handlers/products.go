package handlers

import (
	"log"
	"net/http"
	"github.com/clandau/microsvcs-in-go/product-api/data"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	// handle get request
	if r.Method == http.MethodGet {
		p.GetProducts(rw, r)
		return
	}

	// handle an update
	if r.Method == http.MethodPut {
		
	}

	// catch all
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to Marshal JSON", http.StatusInternalServerError)
	}
}