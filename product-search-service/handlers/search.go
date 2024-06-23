package handlers

import (
	"net/http"
	"product-search-service/services"
	"product-search-service/utils"
	"strconv"
)

func SearchProducts(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	keyword := query.Get("keyword")
	priceMin, _ := strconv.ParseFloat(query.Get("price_min"), 64)
	priceMax, _ := strconv.ParseFloat(query.Get("price_max"), 64)
	category := query.Get("category")
	sortBy := query.Get("sort_by")
	sortOrder := query.Get("sort_order")

	cacheKey := r.URL.String()
	if cachedProducts, found := services.GetFromCache(cacheKey); found {
		utils.RespondWithJSON(w, http.StatusOK, cachedProducts)
		return
	}

	products := services.SearchProducts(keyword, priceMin, priceMax, category, sortBy, sortOrder)
	services.AddToCache(cacheKey, products, services.CacheDuration)

	utils.RespondWithJSON(w, http.StatusOK, products)
}

func SuggestProducts(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	userID := query.Get("user_id")

	suggestions := services.GetSuggestions(userID)
	utils.RespondWithJSON(w, http.StatusOK, suggestions)
}
