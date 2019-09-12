package main

import (
	"flag"
	"log"
)

// Getting parameters from CLI
func cliParameters() (string, string, string, string, bool) {
	var (
		query        = ""
		domain       = flag.String("s", "", "Site's domain")
		countryCode  = flag.String("c", "fr", "Country code")
		languageCode = flag.String("l", "fr", "Language code")
		debug        = flag.Bool("d", false, "Debug mode")
	)
	flag.Parse()
	if *domain == "" {
		log.Fatalln("Site's Domain must be defined, please use \"-s\" flag")
	} else {
		query = "site:" + *domain
	}

	return query, *domain, *countryCode, *languageCode, *debug
}
