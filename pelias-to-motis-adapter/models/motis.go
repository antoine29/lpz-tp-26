package models

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
