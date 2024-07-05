package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/create-payment-intent", handleCreatePaymentIntent)

	var err error = http.ListenAndServe("localhost:4242", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func handleCreatePaymentIntent(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ENDPOINT CALLED!")
}