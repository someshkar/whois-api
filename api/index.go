package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/likexian/whois-go"
	whoisparser "github.com/likexian/whois-parser-go"
)

// Handler handles POSTed JSON data
func Handler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	domain := r.Form.Get("domain")
	whois := getWhois(domain)

	jsonWhois, err := json.Marshal(whois)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	// w.Write(jsonWhois)
	fmt.Fprintf("%s", string(jsonWhois))
}

// getWhois gets structured WhoIs data
func getWhois(domain string) whoisparser.WhoisInfo {
	raw, err := whois.Whois(domain)
	if err != nil {
		log.Fatalln(err)
	}

	result, err := whoisparser.Parse(raw)
	if err != nil {
		log.Fatalln(err)
	}

	return result
}
