package helpers

import (
	"time"

	"pelias-to-motis-adapter/models"
)

func BuildPeliasSearchResponse(params models.PeliasSearchParams, motisPlaces []models.MotisPlace) models.PeliasResponse {
	features := make([]models.Feature, len(motisPlaces))
	for i, place := range motisPlaces {
		features[i] = models.Feature{
			Type: "Feature",
			Geometry: models.Geometry{
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

	return models.PeliasResponse{
		Geocoding: models.Geocoding{
			Version:     "0.1",
			Attribution: "http://some-attribution",
			Query: models.Query{
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
			Engine: models.Engine{
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

func BuildPeliasReverseResponse(params models.PeliasReverseParams, motisPlaces []models.MotisPlace) models.PeliasResponse {
	features := make([]models.Feature, len(motisPlaces))
	for i, place := range motisPlaces {
		features[i] = models.Feature{
			Type: "Feature",
			Geometry: models.Geometry{
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

	return models.PeliasResponse{
		Geocoding: models.Geocoding{
			Version:     "0.1",
			Attribution: "http://some-attribution",
			Query: models.Query{
				// Text:       params.Text,
				Size:      len(features),
				Lang:      params.Lang,
				Layers:    params.Layers,
				Private:   false,
				QuerySize: 100,
			},
			Warnings: []string{"Invalid Parameter: digitransit-subscription-key"},
			Engine: models.Engine{
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
