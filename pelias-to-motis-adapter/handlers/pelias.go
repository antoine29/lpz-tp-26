package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"pelias-to-motis-adapter/helpers"
)

func PeliasSearch(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "*")

	if r.Method == http.MethodOptions {
		return
	}

	peliasParams := helpers.BuildPeliasSearchParams(r.URL.Query())

	motisPlaces, err := helpers.MotisGeocode(peliasParams)
	if err != nil {
		http.Error(w, fmt.Sprintf("MOTIS request failed: %v", err), http.StatusInternalServerError)
		return
	}

	peliasResp := helpers.BuildPeliasSearchResponse(peliasParams, motisPlaces)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(peliasResp)
}

func PeliasReverse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "*")

	if r.Method == http.MethodOptions {
		return
	}

	motisParams := helpers.BuildPeliasReverseParams(r.URL.Query())

	motisResp, err := helpers.MotisReverseGeocode(motisParams)
	if err != nil {
		http.Error(w, fmt.Sprintf("MOTIS request failed: %v", err), http.StatusInternalServerError)
		return
	}

	peliasResp := helpers.BuildPeliasReverseResponse(motisParams, motisResp)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(peliasResp)
}
