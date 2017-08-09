package pinger

import "fmt"

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
	return fmt.Errorf("error")
}
