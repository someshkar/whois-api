package lib

import (
	"github.com/likexian/whois-go"
	whoisparser "github.com/likexian/whois-parser-go"
)

// GetWhois does a WHOIS lookup for a supplied domain
func GetWhois(domain string) (whoisparser.WhoisInfo, error) {
	raw, err := whois.Whois(domain)
	if err != nil {
		return whoisparser.WhoisInfo{}, err
	}

	result, err := whoisparser.Parse(raw)
	if err != nil {
		return whoisparser.WhoisInfo{}, err
	}

	return result, nil
}
