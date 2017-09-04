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
	ServiceUrl     = flag.String("u", "", "HTTP endpoint URL {string} (Required)")
	PingInterval   = flag.Int("i", 5, "Ping Interval {integer} (Seconds) (Optional) (Default 5 seconds)")
	RequestTimeout = flag.Int("t", 5, "Request Timeout {integer} (Seconds) (Optional) (Default 2 seconds)")
)

var quitChan chan bool

func init() {
	flag.Parse()

	if *ServiceUrl == "" {
		log.Printf("Missing required flags\n")
		flag.PrintDefaults()
		os.Exit(1)
	}

	go SignalHandler(make(chan os.Signal, 1))
	log.Printf("Start pinging %s every %v seconds\n", *ServiceUrl, *PingInterval)
}

func main() {
	pinger := p.NewPinger(*ServiceUrl, *RequestTimeout)

	ticker := time.NewTicker(time.Duration(*PingInterval) * time.Second)
	quitChan = make(chan bool)

	for {
		select {
		case <-ticker.C:
			err := pinger.Ping()
			if err != nil {
				log.Println("Error while pinging URL", *ServiceUrl, err)
			} else {
				log.Println("Ping successful", *ServiceUrl)
			}
		case <-quitChan:
			ticker.Stop()
			return
		}
	}
}

func closeTicker() {
	quitChan <- true
	close(quitChan)
}

func SignalHandler(c chan os.Signal) {
	signal.Notify(c, os.Interrupt, syscall.SIGABRT, syscall.SIGTERM, syscall.SIGQUIT)
	for s := <-c; ; s = <-c {
		switch s {
		case syscall.SIGABRT:
			log.Println("Aborting...")
			closeTicker()
			os.Exit(0)
		case os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT:
			log.Println("Shutting Down...")
			closeTicker()
			os.Exit(1)
		}
	}
}
