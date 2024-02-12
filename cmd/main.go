package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/gorilla/mux"
)

var secretKey = []byte("secret-key")

type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

// Static data for three products
var products = []Product{
	{ID: 1, Name: "Product 1", Price: 19.99},
	{ID: 2, Name: "Product 2", Price: 29.99},
	{ID: 3, Name: "Product 3", Price: 39.99},
}

func main() {
	router := mux.NewRouter()

	// Auth Service
	router.HandleFunc("/auth/login", authLoginHandler).Methods("POST")
	router.HandleFunc("/auth/logout", authLogoutHandler).Methods("POST")

	// Product Service
	router.HandleFunc("/products", productListHandler).Methods("GET")
	router.HandleFunc("/products/{id}", productDetailHandler).Methods("GET")

	// Checkout Service
	router.HandleFunc("/checkout/placeorder", checkoutPlaceOrderHandler).Methods("POST")

	port := 8080
	log.Printf("Server is running on port %d...\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))
}

func authLoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	username := r.FormValue("username")
	password := r.FormValue("password")

	// For simplicity, we'll use a hardcoded username and password
	// This should be replaced with a proper authentication mechanism
	if username == "user" && password == "pass" {
		token, err := createToken(username)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"error": "Error generating the token"}`))
			return
		}

		w.Write([]byte(fmt.Sprintf(`{"token": "%s"}`, token)))
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`{"error": "Invalid credentials"}`))
	}
}

func authLogoutHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// In this simple example, we'll just return a success message
	w.Write([]byte(`{"message": "Logout successful"}`))
}

func productListHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response, err := json.Marshal(products)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "Internal Server Error"}`))
		return
	}

	w.Write(response)
}

func productDetailHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	productID, ok := vars["id"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "Product ID is missing"}`))
		return
	}

	id, err := strconv.Atoi(productID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "Product ID has wrong format"}`))
		return
	}

	product := findProductByID(products, id)
	if product == nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"error": "Product not found"}`))
		return
	}

	response, err := json.Marshal(product)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "Internal Server Error"}`))
		return
	}

	w.Write(response)
}

func checkoutPlaceOrderHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	token := r.Header.Get("Authorization")

	if token == "" {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`{"error": "Missing authorization header"}`))
		return
	}

	if verifyToken(token) {
		// In this simple example, we'll just return a success message
		w.Write([]byte(`{"message": "Order placed successfully"}`))
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`{"error": "Invalid token"}`))
	}
}

func createToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": username,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func verifyToken(tokenString string) bool {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	return err == nil && token.Valid
}

func findProductByID(products []Product, id int) *Product {
	for _, product := range products {
		if product.ID == id {
			return &product
		}
	}
	return nil
}
