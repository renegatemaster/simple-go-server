package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Order struct {
	Id   int
	Name string
}

var orders []Order

func main() {
	http.HandleFunc("/orders", ordersHandler)
	http.HandleFunc("/health", healthCheckHandler)

	log.Println("\nServer started listening on port 8080")
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func ordersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getOrders(w, r)
	case http.MethodPost:
		postOrder(w, r)
	default:
		http.Error(w, "Provided method is not allowed.", http.StatusMethodNotAllowed)
	}
}

func getOrders(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(orders)
	fmt.Fprintf(w, "get orders %v", orders)
}

func postOrder(w http.ResponseWriter, r *http.Request) {
	var order Order
	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	orders = append(orders, order)
	fmt.Fprintf(w, "post new order '%v'", order)
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "http server works correctly")
}
