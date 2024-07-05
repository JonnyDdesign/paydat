package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/create-payment-intent", handleCreatePaymentIntent)

	http.ListenAndServe(addr:"localhost:4242", handler:nil)
}

func handleCreatePaymentIntent(w http.ResponseWriter, r *http.Request) {
	fmt.Println(a...: "ENDPOINT CALLED!")
}