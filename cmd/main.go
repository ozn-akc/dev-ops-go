package main

import (
	"fmt"
	"log"
	"net/http"
	
	"github.com/gorilla/mux"
	"github.com/user/repo/internal"
)

func main() {
	router := mux.NewRouter()

	// Auth Service
	router.HandleFunc("/auth/login", internal.AuthLoginHandler).Methods("POST")
	router.HandleFunc("/auth/logout", internal.AuthLogoutHandler).Methods("POST")

	// Product Service
	router.HandleFunc("/products", internal.ProductListHandler).Methods("GET")
	router.HandleFunc("/products/{id}", internal.ProductDetailHandler).Methods("GET")

	// Checkout Service
	router.HandleFunc("/checkout/placeorder", internal.CheckoutPlaceOrderHandler).Methods("POST")

	port := 8080
	log.Printf("Server is running on port %d...\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))
}
