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
	Review       Review  `json:"review"`
	Reviews      int     `json:"reviews"`
}

type Review struct {
	ProductID   int    `json:"product_id"`
	Description string `json:"description"`
	Stars       int    `json:"stars"`
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
	log.Println("fetching products..")
	data := make([]Product, len(products))
	copy(data, products)

	// get reviews
	reviews := getReviews()
	for i, _ := range data {
		for k, _ := range reviews {
			if data[i].ID == reviews[k].ProductID {
				data[i].Review = reviews[k]
				data[i].Reviews += 1
			}
		}
	}
	json.NewEncoder(w).Encode(data)
}

func getReviews() []Review {
	log.Println("fetching reviews..")
	res, err := http.Get(reviewsAPI)
	if err != nil {
		log.Printf("Error fetching reviews: %s\n", err.Error())
		return []Review{}
	}
	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf("Error reading response body for reviews: %s\n", err.Error())
		return []Review{}
	}

	var reviews []Review
	err = json.Unmarshal(data, &reviews)
	if err != nil {
		log.Printf("Error unmarshaling reviews: %s\n", err.Error())
		return []Review{}
	}
	return reviews
}

func health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

var reviewsAPI string

func main() {
	mux := http.NewServeMux()

	reviewsAddr := os.Getenv("REVIEWS_PROXY_ADDR")
	reviewsAPI = fmt.Sprintf("%s/api/reviews", reviewsAddr)
	log.Printf("reviews api: %s\n", reviewsAPI)

	mux.HandleFunc("/health", health)
	mux.HandleFunc("/api/products", getProducts)

	server := &http.Server{
		Addr:    ":" + os.Getenv("PORT"),
		Handler: mux,
	}

	log.Println("Starting server..")
	log.Fatal(server.ListenAndServe(), nil)
}
