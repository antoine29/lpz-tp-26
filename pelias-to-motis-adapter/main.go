package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8084"
	}

	http.HandleFunc("/v1/search", geocode)
	fmt.Printf("Starting adapter service on :%s\n", port)
	http.ListenAndServe(":"+port, nil)
}
