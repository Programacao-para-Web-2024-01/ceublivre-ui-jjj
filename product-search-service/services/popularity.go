package services

import (
	"product-search-service/models"
	"sort"
	"sync"
)

var (
	searchCounts = make(map[string]int)
	countsLock   = sync.RWMutex{}
)

func IncrementSearchCount(productID string) {
	countsLock.Lock()
	defer countsLock.Unlock()
	searchCounts[productID]++
}

func GetMostSearchedProducts() []models.Product {
	countsLock.RLock()
	defer countsLock.RUnlock()

	type kv struct {
		Key   string
		Value int
	}

	var ss []kv
	for k, v := range searchCounts {
		ss = append(ss, kv{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})

	var mostSearched []models.Product
	for _, kv := range ss {
		for _, product := range products {
			if product.ID == kv.Key {
				mostSearched = append(mostSearched, product)
				break
			}
		}
		if len(mostSearched) >= 5 { //5 produtos mais pesquisados
			break
		}
	}

	return mostSearched
}
