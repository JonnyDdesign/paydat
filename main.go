package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/create-payment-intent", handleCreatePaymentIntent)
}

func handleCreatePaymentIntent(w http.ResponseWriter, r *http.Request) {
	fmt.Println(a...: "ENDPOINT CALLED!")
}