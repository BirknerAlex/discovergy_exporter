package main

import (
	"flag"
	"github.com/BirknerAlex/discovergy_exporter/pkg/discovergy"
	"github.com/BirknerAlex/discovergy_exporter/pkg/prometheus"
	"log"
)

func main() {
	log.Printf("Starting Discovergy Exporter")

	email := flag.String("email", "", "Discovergy account email")
	password := flag.String("password", "", "Discovergy account password")
	listen := flag.String("listen", ":2112", "Listen address")

	flag.Parse()

	if *email == "" {
		log.Fatalf("Missing email address for Discovergy account")
	}

	if *password == "" {
		log.Fatalf("Missing password for Discovergy account")
	}

	go func() {
		err := discovergy.RunUpdater(*email, *password)
		if err != nil {
			log.Fatalf("Failed to run updater: %v", err)
		}
	}()

	if err := prometheus.Serve(*listen); err != nil {
		log.Fatalf("Serving failed: %v", err)
	}
}
