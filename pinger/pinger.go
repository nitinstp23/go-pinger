package pinger

import (
	"fmt"
	"net/http"
	"time"
)

type Pinger struct {
	url        string
	interval   int
	reqTimeout time.Duration
}

func NewPinger(url string, interval int, timeoutInt int) *Pinger {
	return &Pinger{
		url:        url,
		interval:   interval,
		reqTimeout: time.Duration(time.Duration(timeoutInt) * time.Second),
	}
}

func (p *Pinger) Ping() error {
	req, err := http.NewRequest("GET", p.url, nil)

	client := &http.Client{
		Timeout: p.reqTimeout,
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
