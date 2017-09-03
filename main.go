package main

import (
	"flag"
	"log"
	"os"

	p "github.com/nitinstp23/go-pinger/pinger"
)

var (
	ServiceUrl     = flag.String("u", "", "HTTP endpoint URL {string} (Required)")
	PingInterval   = flag.Int("i", 5, "Ping Interval {int} (Optional) (Default 5 seconds)")
	RequestTimeout = flag.Int("t", 5, "Request Timeout {int} (Optional) (Default 2 seconds)")
)

func init() {
	flag.Parse()

	if *ServiceUrl == "" {
		log.Printf("Missing required flags\n")
		flag.PrintDefaults()
		os.Exit(1)
	}

	log.Printf("ServiceUrl: %s, PingInterval: %v\n", *ServiceUrl, *PingInterval)
}

func main() {
	pinger := p.NewPinger(*ServiceUrl, *PingInterval, *RequestTimeout)

	err := pinger.Ping()
	if err != nil {
		log.Println("Error while pinging URL", *ServiceUrl, err)
	} else {
		log.Println("Ping successful", *ServiceUrl)
	}
}
