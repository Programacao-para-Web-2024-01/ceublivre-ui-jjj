package services

import (
	"product-search-service/models"
	"sync"
	"time"
)

var (
	cache     = make(map[string][]models.Product)
	cacheLock = sync.RWMutex{}
)

func GetFromCache(key string) ([]models.Product, bool) {
	cacheLock.RLock()
	defer cacheLock.RUnlock()
	products, found := cache[key]
	return products, found
}

func AddToCache(key string, products []models.Product, duration int) {
	cacheLock.Lock()
	defer cacheLock.Unlock()
	cache[key] = products

	go func() {
		time.Sleep(time.Duration(duration) * time.Second)
		cacheLock.Lock()
		delete(cache, key)
		cacheLock.Unlock()
	}()
}
