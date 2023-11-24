package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

const apiUrl = "https://api.coingecko.com/api/v3/simple/price?ids=bitcoin,ethereum,tether&vs_currencies=cad"

func main() {
	http.HandleFunc("/price/", GetPrice)

	log.Fatal(http.ListenAndServe(":9090", nil))
}

func GetPrice(w http.ResponseWriter, r *http.Request) {
	cryptoName := extractCryptoName(r.URL.Path)

	price, err := fetchCryptoPrice(cryptoName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]float64{cryptoName: price}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func extractCryptoName(path string) string {
	// Trim the leading "/price/" from the path
	path = strings.TrimPrefix(path, "/price/")
	// Split the path by "/"
	parts := strings.Split(path, "/")
	if len(parts) >= 1 {
		return parts[0]
	}
	return ""
}

func fetchCryptoPrice(cryptoName string) (float64, error) {
	resp, err := http.Get(apiUrl)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	var result map[string]map[string]float64
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return 0, err
	}

	return result[cryptoName]["cad"], nil
}
