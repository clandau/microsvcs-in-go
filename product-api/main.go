package main

import (
	"log"
	"os"
	"net/http"
	"github.com/clandau/microsvcs-in-go/product-api/handlers"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	hh := handlers.NewHello(l)

	
}