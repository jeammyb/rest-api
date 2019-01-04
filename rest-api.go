// Simple Golang REST API example that provide a JSON with details about contries.
// It consumes data from the API provided by https://restcountries.eu

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

var client = &http.Client{Timeout: 10 * time.Second}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/country/{name}", country).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", r))
}

func country(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]

	data, err := countryData(name)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(404)
		return
	}

	json.NewEncoder(w).Encode(data)

}

func countryData(name string) ([]CountryResponse, error) {
	// getting country data
	url := fmt.Sprintf("https://restcountries.eu/rest/v2/name/%s", name)
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		err = fmt.Errorf("Error getting data: status code %s", resp.Status)
		return nil, err
	}

	var countries []CountryResponse
	err = json.NewDecoder(resp.Body).Decode(&countries)
	if err != nil {
		return nil, err
	}

	return countries, nil
}

// CountryResponse contains data about a country
type CountryResponse struct {
	Name       string `json:"name"`
	Capital    string `json:"capital"`
	Region     string `json:"region"`
	Subregion  string `json:"subregion"`
	Population int32  `json:"population"`
}
