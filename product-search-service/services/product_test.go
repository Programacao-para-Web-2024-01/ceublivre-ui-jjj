// PEDI PRO CHAT GEPETAS FAZER UM TESTE PRA VER SE TA TUDO CERTINHO

package services

import (
	"testing"
)

func TestSearchProducts(t *testing.T) {
	tests := []struct {
		name      string
		keyword   string
		priceMin  float64
		priceMax  float64
		category  string
		sortBy    string
		expected  int
		sortOrder string
	}{
		{"Keyword: laptop", "laptop", 0, 0, "", "", 1, ""},
		{"Category: Accessories", "", 0, 0, "Accessories", "", 3, ""},
		{"Price range and category", "", 50, 150, "Accessories", "", 1, ""},
		{"Sort by rating", "", 0, 0, "", "rating", 10, ""},
		{"Keyword: smart", "smart", 0, 0, "", "", 2, ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SearchProducts(tt.keyword, tt.priceMin, tt.priceMax, tt.category, tt.sortBy, tt.sortOrder)
			if len(result) < tt.expected {
				t.Errorf("expected at least %d products, got %d", tt.expected, len(result))
			}
		})
	}
}
