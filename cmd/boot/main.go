package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/mehdi-fakhraei60/Dcare/gateway"
)

func main() {
	if err := gateway.Prove(context.Background(), os.Getenv("GATEWAY_IDENTITY_URL"), "dcare"); err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "{\"app\":\"dcare\",\"proved\":true}")
	})
	addr := os.Getenv("APP_ADDR")
	if addr == "" {
		addr = ":8080"
	}
	log.Fatal(http.ListenAndServe(addr, nil))
}
