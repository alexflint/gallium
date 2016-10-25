package gallium

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"
	"net"
	"net/http"
)

const tokenKey = "x-gallium-token"

type emitter interface {
	Emit(event string, data interface{}) error
}

type server struct {
	http.Server
	emitter

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
	sse := newSSESource(h)
	s.emitter = sse
	s.Handler = newProtectedHandler(sse, s.token)
	log.Printf("Starting server %s", s.BaseURL())
	go s.Serve(l)
	return s
}

func token() string {
	var p [128]byte
	_, err := rand.Read(p[:])
	if err != nil {
		panic(fmt.Sprintf("Failed generating random token: %s", err.Error()))
	}
	return base64.RawURLEncoding.EncodeToString(p[:])
}

func (s *server) Close() error {
	return s.listener.Close()
}

func (s *server) BaseURL() string {
	return "http://" + s.listener.Addr().String()
}

func (s *server) URL() string {
	return s.BaseURL() + "/?" + tokenKey + "=" + s.token
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
	if c, err := r.Cookie(tokenKey); err == nil && c.Value == h.token {
		h.serve(w, r)
		return
	}
	if r.URL.Query().Get(tokenKey) == h.token {
		http.SetCookie(w, &http.Cookie{
			Name:  tokenKey,
			Value: h.token,
		})
		h.serve(w, r)
		return
	}
	log.Printf("gallium=%q method=%s path=%q", "unauthorized", r.Method, r.RequestURI)
	w.WriteHeader(http.StatusUnauthorized)
}

func (h *protectedHandler) serve(w http.ResponseWriter, r *http.Request) {
	// Strip query parameter token
	params := r.URL.Query()
	params.Del(tokenKey)
	r.URL.RawQuery = params.Encode()
	r.RequestURI = r.URL.RequestURI()

	// Strip cookie token
	cookies := r.Cookies()[:]
	r.Header.Del("Cookie")
	for _, c := range cookies {
		if c.Name == tokenKey {
			continue
		}
		r.AddCookie(c)
	}

	log.Printf("gallium=%q method=%s path=%q", "ok", r.Method, r.URL.Path)
	h.handler.ServeHTTP(w, r)
}
