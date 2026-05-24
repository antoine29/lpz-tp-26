package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"pelias-to-motis-adapter/helpers"
)

func Geocode(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "*")

	if r.Method == http.MethodOptions {
		return
	}

	motisParams := helpers.BuildMotisParams(r.URL.Query())

	motisResp, err := helpers.MotisGeocode(motisParams)
	if err != nil {
		http.Error(w, fmt.Sprintf("MOTIS request failed: %v", err), http.StatusInternalServerError)
		return
	}

	peliasResp := helpers.BuildPeliasResponse(motisParams, motisResp)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(peliasResp)
}
