package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/user/repo/checkout-service/internal"
)

func main() {
	router := mux.NewRouter()

	// Checkout Service
	router.HandleFunc("/checkout/placeorder", internal.CheckoutPlaceOrderHandler).Methods("POST")

	port := 8080
	log.Printf("The server is running on port %d...\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))
}
