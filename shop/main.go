package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func getDetailOfShop(w http.ResponseWriter, r *http.Request) {

	// asumsi kita sudah mendapatkan detail shop dengan id 01 dari db
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"id":                "01",
		"name":              "Andromeda Shop",
		"number_of_product": "10",
	})
}

func getAllShop(w http.ResponseWriter, r *http.Request) {

	// asumsi kita sudah mendapatkan all shop dari db
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode([]map[string]string{
		{
			"id":                "01",
			"name":              "Andromeda Shop",
			"number_of_product": "10",
		},
		{
			"id":                "02",
			"name":              "Berkah Shop",
			"number_of_product": "55",
		},
		{
			"id":                "03",
			"name":              "Baby Shop",
			"number_of_product": "100",
		},
	})
}

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/v1/shops/01", getDetailOfShop)
	mux.HandleFunc("/v1/shops", getAllShop)

	fmt.Println("Server Running")

	http.ListenAndServe(":9000", mux)
}
