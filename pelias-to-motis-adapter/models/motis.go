package models

type MotisPlace struct {
	Type     string  `json:"type"`
	Name     string  `json:"name"`
	ID       string  `json:"id"`
	Lat      float64 `json:"lat"`
	Lon      float64 `json:"lon"`
	Category string  `json:"category"`
}
