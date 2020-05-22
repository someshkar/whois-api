package api

import (
	"fmt"
	"log"
	"net/http"

	jsoniter "github.com/json-iterator/go"
	"github.com/someshkar/whois-api/lib"
	"github.com/someshkar/whois-api/structs"
)

// MultiHandler handles Whois requests for multiple domains
func MultiHandler(w http.ResponseWriter, r *http.Request) {
	// Make sure it's a POST request
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Please only POST data here")
		return
	}

	// Decode JSON body
	decoder := jsoniter.NewDecoder(r.Body)
	var body structs.MultiBody

	err := decoder.Decode(&body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	allWhois := lib.GetMultiWhois(body.Domains)

	jsonAllWhois, err := jsoniter.Marshal(allWhois)
	if err != nil {
		log.Fatalln(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonAllWhois)
}
