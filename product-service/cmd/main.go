package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/user/repo/product-service/internal"
)

func main() {
	router := mux.NewRouter()

	// Product Service
	router.HandleFunc("/products", internal.ProductListHandler).Methods("GET")
	router.HandleFunc("/products/{id}", internal.ProductDetailHandler).Methods("GET")

	port := 8080
	log.Printf("The server is running on port %d...\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))
}
