package pinger

import (
	"fmt"
	"net/http"
	"time"
)

type Pinger struct {
	url      string
	interval int
}

func NewPinger(url string, interval int) *Pinger {
	return &Pinger{
		url:      url,
		interval: interval,
	}
}

func (p *Pinger) Ping() error {
	timeout := time.Duration(time.Duration(5) * time.Second)
	req, err := http.NewRequest("GET", p.url, nil)

	client := &http.Client{
		Timeout: timeout,
	}

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("Failed to ping url: %v, %v", p.url, err)
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Request failed for: %v with status: %v", p.url, resp.Status)
	}

	return nil
}
