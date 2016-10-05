package gallium

import (
	"net/http"
	"net/http/cookiejar"
	"strings"
	"testing"
)

func TestNewServer(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	})
	s := newServer(mux)
	if !strings.HasPrefix(s.BaseURL(), "http://127.0.0.1:") {
		t.Fatalf("Invalid URL: %s", s.BaseURL())
	}

	jar, err := cookiejar.New(nil)
	if err != nil {
		t.Fatalf("Failed creating cookie jar: %s", err)
	}
	client := &http.Client{
		Jar: jar,
	}

	// Fresh unauthorized request
	res, err := client.Get(s.BaseURL())
	if err != nil {
		t.Fatalf("Failed requesting URL")
	}
	if res.StatusCode != http.StatusUnauthorized {
		t.Fatalf("Expected Status: %d Received Status: %d", http.StatusUnauthorized, res.StatusCode)
	}

	// Initial request with authorization token
	res, err = client.Get(s.URL())
	if err != nil {
		t.Fatalf("Failed requesting URL")
	}
	if res.StatusCode != http.StatusOK {
		t.Fatalf("Expected Status: %d Received Status: %d", http.StatusOK, res.StatusCode)
	}

	// Subsequent request with authorized cookie
	res, err = client.Get(s.BaseURL())
	if err != nil {
		t.Fatalf("Failed requesting URL")
	}
	if res.StatusCode != http.StatusOK {
		t.Fatalf("Expected Status: %d Received Status: %d", http.StatusOK, res.StatusCode)
	}

	if err := s.Close(); err != nil {
		t.Fatalf("Failed closing: %s", err)
	}
}
