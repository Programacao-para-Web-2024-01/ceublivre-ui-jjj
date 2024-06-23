package models

type Product struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	Brand    string  `json:"brand"`
	Price    float64 `json:"price"`
	Category string  `json:"category"`
	Rating   float64 `json:"rating"`
}
