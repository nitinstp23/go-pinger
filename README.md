# Go-Pinger

A CLI application which periodically pings an HTTP endpoint.

## Project Setup

* `curl -O /usr/local/go1.8.3.darwin-amd64.tar.gz https://storage.googleapis.com/golang/go1.8.3.darwin-amd64.tar.gz`
* `tar -C /usr/local -xzf go1.8.3.darwin-amd64.tar.gz`
* `export PATH=$PATH:/usr/local/go/bin`
* `make build`
* `./go_ping -u http://example.com -i 5`
