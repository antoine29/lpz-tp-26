package helpers

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"net/url"
	"strings"

	"pelias-to-motis-adapter/models"
)

var MotisHost string

func BuildMotisParams(query url.Values) models.SearchParams {
	text := query.Get("text")
	lang := query.Get("lang")
	if lang == "" {
		lang = "en"
	}

	sources := parseCommaSeparated(query.Get("sources"))
	layers := parseCommaSeparated(query.Get("layers"))

	return models.SearchParams{
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

func MotisGeocode(params models.SearchParams) ([]models.MotisPlace, error) {
	targetURL := fmt.Sprintf("%s?text=%s&language=%s&mode=%s",
		MotisHost+"/api/v1/geocode",
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

	var places []models.MotisPlace
	if err := json.NewDecoder(resp.Body).Decode(&places); err != nil {
		slog.Error("failed to decode request response", "error", err)
		return nil, err
	}

	return places, nil
}
