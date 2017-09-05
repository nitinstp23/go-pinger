package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	p "github.com/nitinstp23/go-pinger/pinger"
)

var (
	serviceUrl     = flag.String("u", "", "HTTP endpoint URL {string} (Required)")
	pingInterval   = flag.Int("i", 5, "Ping Interval {integer} (Seconds) (Optional) (Default 5 seconds)")
	requestTimeout = flag.Int("t", 2, "Request Timeout {integer} (Seconds) (Optional) (Default 2 seconds)")
	quitChan       = make(chan os.Signal, 1)
)

func init() {
	flag.Parse()

	if *serviceUrl == "" {
		log.Printf("Missing required flags\n")
		flag.PrintDefaults()
		os.Exit(1)
	}

	AddSignalHandler()
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

func AddSignalHandler() {
	signal.Notify(
		quitChan,
		os.Interrupt,
		syscall.SIGABRT,
		syscall.SIGTERM,
		syscall.SIGQUIT)
}
