package gallium

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"net"
	"net/http"
)

type server struct {
	http.Server

	listener net.Listener
	token    string
}

func newServer(h http.Handler) *server {
	s := &server{}
	l, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		if l, err = net.Listen("tcp6", "[::1]:0"); err != nil {
			panic(fmt.Sprintf("gallium: failed to listen on a port: %v", err))
		}
	}
	s.listener = l
	s.token = token()
	s.Handler = newProtectedHandler(h, s.token)
	go s.Serve(l)
	return s
}

func token() string {
	p := make([]byte, 128)
	_, err := rand.Read(p)
	if err != nil {
		panic(fmt.Sprintf("Failed generating random token: %s", err.Error()))
	}
	return base64.RawURLEncoding.EncodeToString(p)
}

func (s *server) Close() error {
	return s.listener.Close()
}

func (s *server) BaseURL() string {
	return "http://" + s.listener.Addr().String()
}

func (s *server) URL() string {
	return s.BaseURL() + "/?gallium-token=" + s.token
}

type protectedHandler struct {
	handler http.Handler
	token   string
}

func newProtectedHandler(h http.Handler, token string) http.Handler {
	return &protectedHandler{
		handler: h,
		token:   token,
	}
}

func (h *protectedHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if c, err := r.Cookie("x-gallium-token"); err == nil && c.Value == h.token {
		h.handler.ServeHTTP(w, r)
		return
	}
	if r.URL.Query().Get("gallium-token") == h.token {
		http.SetCookie(w, &http.Cookie{
			Name:  "x-gallium-token",
			Value: h.token,
		})
		h.handler.ServeHTTP(w, r)
		return
	}
	w.WriteHeader(http.StatusUnauthorized)
}
