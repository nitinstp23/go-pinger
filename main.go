package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	serviceUrl := flag.String("url", "", "HTTP endpoint URL {string} (Required)")
	pingInterval := flag.Int("interval", 5, "Ping Interval {int} (Required)")

	flag.Parse()

	if *serviceUrl == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	fmt.Printf("serviceUrl: %s, pingInterval: %v\n", *serviceUrl, *pingInterval)
}
