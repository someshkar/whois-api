package lib

import (
	"log"

	"github.com/likexian/whois-go"
	whoisparser "github.com/likexian/whois-parser-go"
)

// GetMultiWhois returns an array of Whois lookups
// for an array of domains
func GetMultiWhois(domains []string) []whoisparser.WhoisInfo {
	allWhois := []whoisparser.WhoisInfo{}

	whoisCh := make(chan whoisparser.WhoisInfo)
	errorCh := make(chan error)

	for _, domain := range domains {
		go getChanWhois(domain, whoisCh, errorCh)
	}

	allWhois = append(allWhois, <-whoisCh)

	log.Fatalln(<-errorCh)

	return allWhois
}

// getChanWhois sends Whois data to a channel
func getChanWhois(domain string, whoisCh chan<- whoisparser.WhoisInfo, errorCh chan<- error) {
	raw, err := whois.Whois(domain)
	if err != nil {
		// return whoisparser.WhoisInfo{}, err
		whoisCh <- whoisparser.WhoisInfo{}
		errorCh <- err
	}

	result, err := whoisparser.Parse(raw)
	if err != nil {
		// return whoisparser.WhoisInfo{}, err
		whoisCh <- whoisparser.WhoisInfo{}
		errorCh <- err
	}

	whoisCh <- result
}
