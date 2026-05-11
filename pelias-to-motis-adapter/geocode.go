package main

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

type MotisPlace struct {
	Type     string  `json:"type"`
	Name     string  `json:"name"`
	ID       string  `json:"id"`
	Lat      float64 `json:"lat"`
	Lon      float64 `json:"lon"`
	Category string  `json:"category"`
}

type SearchParams struct {
	Text    string
	Lang    string
	Sources []string
	Layers  []string
}

type Feature struct {
	Type       string                 `json:"type"`
	Geometry   Geometry               `json:"geometry"`
	Properties map[string]interface{} `json:"properties"`
}

type Geometry struct {
	Type        string    `json:"type"`
	Coordinates []float64 `json:"coordinates"`
}

type Query struct {
	Text       string            `json:"text"`
	Size       int               `json:"size"`
	Lang       string            `json:"lang"`
	Layers     []string          `json:"layers"`
	Sources    []string          `json:"sources"`
	Private    bool              `json:"private"`
	QuerySize  int               `json:"querySize"`
	ParsedText map[string]string `json:"parsed_text"`
}

type Geocoding struct {
	Version     string   `json:"version"`
	Attribution string   `json:"attribution"`
	Query       Query    `json:"query"`
	Warnings    []string `json:"warnings"`
	Engine      Engine   `json:"engine"`
	Timestamp   int64    `json:"timestamp"`
}

type Engine struct {
	Name    string `json:"name"`
	Author  string `json:"author"`
	Version string `json:"version"`
}

type PeliasResponse struct {
	Geocoding Geocoding `json:"geocoding"`
	Type      string    `json:"type"`
	Features  []Feature `json:"features"`
}

var geocodeEndpoint = func() string {
	if host := os.Getenv("MOTIS_HOST"); host != "" {
		return host + "/api/v1/geocode"
	}

	return "http://localhost:8083/api/v1/geocode"
}()

func geocode(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "*")

	if r.Method == http.MethodOptions {
		return
	}

	params := parseSearchParams(r.URL.Query())

	motisResp, err := fetchMotis(params)
	if err != nil {
		http.Error(w, fmt.Sprintf("MOTIS request failed: %v", err), http.StatusInternalServerError)
		return
	}

	peliasResp := buildPeliasResponse(params, motisResp)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(peliasResp)
}

func parseSearchParams(query url.Values) SearchParams {
	text := query.Get("text")
	lang := query.Get("lang")
	if lang == "" {
		lang = "en"
	}

	sources := parseCommaSeparated(query.Get("sources"))
	layers := parseCommaSeparated(query.Get("layers"))

	return SearchParams{
		Text:    text,
		Lang:    lang,
		Sources: sources,
		Layers:  layers,
	}
}

func parseCommaSeparated(s string) []string {
	if s == "" {
		return []string{}
	}
	decoded, _ := url.QueryUnescape(s)
	parts := strings.Split(decoded, ",")
	result := make([]string, 0, len(parts))
	for _, p := range parts {
		trimmed := strings.TrimSpace(p)
		if trimmed != "" {
			result = append(result, trimmed)
		}
	}
	return result
}

const transportMode = "AIRPLANE,HIGHSPEED_RAIL,LONG_DISTANCE,NIGHT_RAIL,COACH,RIDE_SHARING,REGIONAL_RAIL,SUBURBAN,SUBWAY,TRAM,BUS,FERRY,ODM,FUNICULAR,AERIAL_LIFT,OTHER"

func fetchMotis(params SearchParams) ([]MotisPlace, error) {
	targetURL := fmt.Sprintf("%s?text=%s&language=%s&mode=%s",
		geocodeEndpoint,
		url.QueryEscape(params.Text),
		params.Lang,
		transportMode,
	)

	req, err := http.NewRequest("GET", targetURL, nil)
	if err != nil {
		slog.Error("failed to create request", "error", err)
		return nil, err
	}

	req.Header.Set("Accept", "*/*")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("lang", params.Lang)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		slog.Error("failed to send request", "error", err)
		return nil, err
	}

	defer resp.Body.Close()

	var places []MotisPlace
	if err := json.NewDecoder(resp.Body).Decode(&places); err != nil {
		slog.Error("failed to decode request response", "error", err)
		return nil, err
	}

	return places, nil
}

func buildPeliasResponse(params SearchParams, motisPlaces []MotisPlace) PeliasResponse {
	features := make([]Feature, len(motisPlaces))
	for i, place := range motisPlaces {
		features[i] = Feature{
			Type: "Feature",
			Geometry: Geometry{
				Type:        "Point",
				Coordinates: []float64{place.Lon, place.Lat},
			},
			Properties: map[string]interface{}{
				"id":     place.ID,
				"layer":  "venue",
				"source": "openstreetmap",
				"name":   place.Name,
				"label":  place.Name,
			},
		}
	}

	return PeliasResponse{
		Geocoding: Geocoding{
			Version:     "0.1",
			Attribution: "http://some-attribution",
			Query: Query{
				Text:       params.Text,
				Size:       len(features),
				Lang:       params.Lang,
				Layers:     params.Layers,
				Sources:    params.Sources,
				Private:    false,
				QuerySize:  100,
				ParsedText: map[string]string{"neighbourhood": params.Text, "name": params.Text},
			},
			Warnings: []string{"Invalid Parameter: digitransit-subscription-key"},
			Engine: Engine{
				Name:    "Pelias",
				Author:  "Mapzen",
				Version: "1.0",
			},
			Timestamp: time.Now().UnixMilli(),
		},
		Type:     "FeatureCollection",
		Features: features,
	}
}
