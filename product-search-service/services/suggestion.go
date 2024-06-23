package services

import (
	"product-search-service/models"
	"strings"
)

var userSearchHistory = map[string][]string{
	"1": {"Laptop", "Smartphone"},
	"2": {"Headphones", "Tablet"},
	"3": {"Camera", "Smartwatch"},
}

func GetSuggestions(userID string) []models.Product {
	var suggestions []models.Product
	history, found := userSearchHistory[userID]
	if !found {
		return suggestions
	}

	for _, keyword := range history {
		for _, product := range products {
			if contains(product.Name, keyword) || contains(product.Brand, keyword) {
				suggestions = append(suggestions, product)
			}
		}
	}

	suggestions = removeDuplicates(suggestions)
	return suggestions
}

func removeDuplicates(products []models.Product) []models.Product {
	encountered := map[string]bool{}
	var result []models.Product

	for _, product := range products {
		if !encountered[product.ID] {
			encountered[product.ID] = true
			result = append(result, product)
		}
	}

	return result
}

func contains(source, keyword string) bool {
	return strings.Contains(strings.ToLower(source), strings.ToLower(keyword))
}
