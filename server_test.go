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
		if c, err := r.Cookie(tokenKey); err == nil && c.Value != "" {
			t.Fatal("Token cookie should be stripped")
		}
		if r.URL.Query().Get(tokenKey) != "" {
			t.Fatal("Token query parmeter should be stripped")
		}
		if strings.Contains(r.RequestURI, tokenKey) {
			t.Fatal("Token should be removed from RequestURI")
		}
		http.SetCookie(w, &http.Cookie{Name: "custom-cookie", Value: "cookie"})
	})
	mux.HandleFunc("/assertKeepsParams", func(w http.ResponseWriter, r *http.Request) {
		if c, err := r.Cookie("custom-cookie"); err != nil || c.Value == "" {
			t.Fatal("Lost custom-cookie")
		}
		if r.URL.Query().Get("custom-param") == "" {
			t.Fatal("Lost custom param")
		}
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

	// Subsequent request keeps params and cookie
	res, err = client.Get(s.BaseURL() + "/assertKeepsParams?custom-param=ok")
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
