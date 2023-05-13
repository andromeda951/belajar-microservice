package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func merchantMiddle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		author := r.Header.Get("Authorization")

		if author != "merchant" {
			w.Write([]byte("Anda tidak punya akses"))
			return
		}

		next.ServeHTTP(w, r)
	}
}

func superMiddle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		author := r.Header.Get("Authorization")

		if author != "su-admin" {
			w.Write([]byte("Anda tidak punya akses"))
			return
		}

		next.ServeHTTP(w, r)
	}
}

type Merchant struct {
	Id          string `json:"id"`
	Owner           string `json:"owner"`
	Address         string `json:"address"`
	ShopName        string `json:"shop_name"`
	NumberOfProduct string `json:"number_of_product"`
}

func getMerchant(w http.ResponseWriter, r *http.Request) {
	response, _ := http.Get("http://localhost:8000/api/v1/merchants/01")

	w.Header().Add("Content-Type", "application/json")
	data, _ := io.ReadAll(response.Body)
	merch := &Merchant{}
	json.Unmarshal(data, merch)
	json.NewEncoder(w).Encode(merch)
}

type Toko struct {
	Id              string `json:"id"`
	ShopName        string `json:"name"`
	NumberOfProduct string `json:"number_of_product"`
}

func getAllShop(w http.ResponseWriter, r *http.Request) {
	response, _ := http.Get("http://localhost:9000/v1/shops")

	w.Header().Add("Content-Type", "application/json")
	data, _ := io.ReadAll(response.Body)
	toko := &[]Toko{}
	json.Unmarshal(data, toko)
	json.NewEncoder(w).Encode(toko)
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/merchants/01", merchantMiddle(getMerchant))
	mux.HandleFunc("/shops", superMiddle(getAllShop))

	fmt.Println("Server running")

	http.ListenAndServe(":6000", mux)
}
