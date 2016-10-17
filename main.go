package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Product struct {
	ID           int     `json:"id"`
	Name         string  `json:"name"`
	Price        float32 `json:"price"`
	Category     string  `json:"category"`
	ThumbnailURL string  `json:"thumbnail_url"`
	Description  string  `json:"description"`
	Stars        int     `json:"stars"`
	Reviews      int     `json:"reviews"`
}

var products []Product

func init() {
	// load fake product data
	raw, err := ioutil.ReadFile("./products.json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	json.Unmarshal(raw, &products)
}

func getProducts(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(products)
}

func health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/health", health)
	mux.HandleFunc("/api/products", getProducts)

	server := &http.Server{
		Addr:    ":" + os.Getenv("PORT"),
		Handler: mux,
	}

	log.Println("Starting server..")
	log.Fatal(server.ListenAndServe(), nil)
}
