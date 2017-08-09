package main

import (
	"flag"
	"log"
	"os"

	p "github.com/nitinstp23/go-pinger/pinger"
)

func main() {
	serviceUrl := flag.String("u", "", "HTTP endpoint URL {string} (Required)")
	pingInterval := flag.Int("i", 5, "Ping Interval {int} (Required)")

	flag.Parse()

	if *serviceUrl == "" {
		log.Printf("Missing required flags\n")
		flag.PrintDefaults()
		os.Exit(1)
	}

	log.Printf("serviceUrl: %s, pingInterval: %v\n", *serviceUrl, *pingInterval)

	pinger := p.NewPinger(*serviceUrl, *pingInterval)

	if err := pinger.Ping(); err != nil {
		log.Println("Error while pinging URL - ", *serviceUrl, err)
	}
}
