package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"pelias-to-motis-adapter/handlers"
	"pelias-to-motis-adapter/helpers"
)

func main() {
	port, isPortSet := os.LookupEnv("PORT")
	if !isPortSet || port == "" {
		port = "8084"
	}

	motisHost, motisHostIsSet := os.LookupEnv("MOTIS_HOST")
	if !motisHostIsSet || motisHost == "" {
		slog.Error("MOTIS_HOST env var mising or empty")
		os.Exit(1)
	}

	helpers.MotisHost = motisHost

	http.HandleFunc("/v1/search", handlers.Geocode)
	fmt.Printf("Starting adapter service on :%s\n", port)
	http.ListenAndServe(":"+port, nil)
}
