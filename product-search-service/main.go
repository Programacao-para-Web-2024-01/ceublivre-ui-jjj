package main

import (
	"log"
	"net/http"

	"product-search-service/handlers"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/search", handlers.SearchProducts).Methods("GET")
	r.HandleFunc("/suggestions", handlers.SuggestProducts).Methods("GET")

	//Configurar o middleware CORS
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // Permitir todas as origens, você pode restringir isso conforme necessário
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	handler := c.Handler(r)

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
