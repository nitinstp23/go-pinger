package main

import (
	"flag"
	"log"
	"net/url"
	"os"
	"os/signal"
	"syscall"
	"time"

	p "github.com/nitinstp23/go-pinger/pinger"
)

var (
	serviceUrl     = flag.String("u", "", "HTTP endpoint URL {string} (Required) (E.g. - http://example.com)")
	pingInterval   = flag.Int("i", 5, "Ping Interval {integer} (Seconds) (Optional) (Default 5 seconds)")
	requestTimeout = flag.Int("t", 2, "Request Timeout {integer} (Seconds) (Optional) (Default 2 seconds)")
	quitChan       = make(chan os.Signal, 1)
)

func init() {
	flag.Parse()

	validateServiceUrl(*serviceUrl)
	validatePingInterval(*pingInterval)
	addSignalHandler()

	log.Printf("Start pinging %s every %v seconds\n", *serviceUrl, *pingInterval)
}

func main() {
	pinger := p.NewPinger(*serviceUrl, *requestTimeout)
	ticker := time.NewTicker(time.Duration(*pingInterval) * time.Second)

	for {
		select {
		case <-ticker.C:
			err := pinger.Ping()
			if err != nil {
				log.Println("Error while pinging URL", *serviceUrl, err)
			} else {
				log.Println("Ping successful", *serviceUrl)
			}
		case s := <-quitChan:
			ticker.Stop()

			switch s {
			case syscall.SIGABRT:
				log.Println("Aborting...")
				os.Exit(0)
			case os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT:
				log.Println("Shutting Down...")
				os.Exit(1)
			}
		}
	}
}

func addSignalHandler() {
	signal.Notify(
		quitChan,
		os.Interrupt,
		syscall.SIGABRT,
		syscall.SIGTERM,
		syscall.SIGQUIT)
}

func validateServiceUrl(serviceUrl string) {
	if serviceUrl == "" {
		log.Println("Missing required flag -u")
		flag.PrintDefaults()
		os.Exit(1)
	}

	if _, err := url.ParseRequestURI(serviceUrl); err != nil {
		log.Printf("Invalid value %s for flag -u", serviceUrl)
		flag.PrintDefaults()
		os.Exit(1)
	}
}

func validatePingInterval(pingInterval int) {
	if pingInterval <= 0 {
		log.Printf("Invalid value %v for flag -i", pingInterval)
		flag.PrintDefaults()
		os.Exit(1)
	}
}
