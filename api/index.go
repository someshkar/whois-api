package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/someshkar/whois-api/lib"
)

// Body defines the JSON body for this route
type Body struct {
	Domain string
}

// Handler handles POSTed JSON data
func Handler(w http.ResponseWriter, r *http.Request) {

	// Make sure it's a POST request
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Please only POST data here")
		return
	}

	// Decode JSON body
	decoder := json.NewDecoder(r.Body)
	var body Body

	err := decoder.Decode(&body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	whois, err := lib.GetWhois(body.Domain)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonWhois, err := json.Marshal(whois)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonWhois)
}

// // getWhois gets structured Whois data
// func getWhois(domain string) whoisparser.WhoisInfo {
// 	raw, err := whois.Whois(domain)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	result, err := whoisparser.Parse(raw)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	return result
// }
