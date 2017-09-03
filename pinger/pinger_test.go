package pinger

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPing(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.RequestURI {
		case "/valid-url":
			w.WriteHeader(http.StatusOK)
		case "/bad-request":
			w.WriteHeader(http.StatusBadRequest)
		default:
			http.Error(w, "not found", http.StatusNotFound)
			return
		}
	})
	httpStub := httptest.NewServer(handler)

	defer httpStub.Close()

	tt := []struct {
		name string
		url  string
		err  string
	}{
		{name: "valid HTTP url", url: "/valid-url"},
		{name: "invalid HTTP url", url: "/bad-request", err: fmt.Sprintf("Request failed for: %s/bad-request with status: 400 Bad Request", httpStub.URL)},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			pinger := NewPinger(httpStub.URL+tc.url, 5, 2)

			err := pinger.Ping()

			if err != nil {
				if tc.err != "" {
					// check errors
					if tc.err != err.Error() {
						t.Errorf("expected error %v; got %v", tc.err, err.Error())
					}
				}
			}
		})
	}
}
