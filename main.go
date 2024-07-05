package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/create-payment-intent", handleCreatePaymentIntent)

	var err error = http.ListenAndServe(addr:"localhost:4242", handler:nil)
	if err != nil {
		log.Fatal(err)
	}
}

func handleCreatePaymentIntent(w http.ResponseWriter, r *http.Request) {
	fmt.Println(a...: "ENDPOINT CALLED!")
}