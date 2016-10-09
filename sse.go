package gallium

import (
	"bufio"
	"bytes"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"sync"
)

type sseSource struct {
	handler http.Handler
	mut     sync.RWMutex
	clients map[string]chan<- sseEvent
}

type sseEvent struct {
	Name string
	Data []byte
}

func (e sseEvent) WriteTo(w io.Writer) (int64, error) {
	var (
		total int64
	)
	if e.Name != "" {
		n, err := fmt.Fprintf(w, "event: %s\n", e.Name)
		if err != nil {
			return total, err
		}
		total += int64(n)
	}
	scanner := bufio.NewScanner(bytes.NewReader(e.Data))
	for scanner.Scan() {
		n, err := fmt.Fprintf(w, "data: %s\n", scanner.Text())
		if err != nil {
			return total, err
		}
		total += int64(n)
	}
	if err := scanner.Err(); err != nil {
		return total, err
	}
	n, err := fmt.Fprintf(w, "\n")
	if err != nil {
		return total, err
	}
	total += int64(n)
	return total, nil
}

func newSSESource(h http.Handler) *sseSource {
	return &sseSource{
		handler: h,
		clients: make(map[string]chan<- sseEvent),
	}
}

func (s *sseSource) Emit(name string, data interface{}) error {
	d, err := json.Marshal(data)
	if err != nil {
		return err
	}
	event := sseEvent{
		Name: name,
		Data: d,
	}

	s.mut.RLock()
	defer s.mut.RUnlock()

	for _, cl := range s.clients {
		cl <- event
	}
	return nil
}

func (s *sseSource) register(id string, ch chan<- sseEvent) (func(), error) {
	s.mut.Lock()
	defer s.mut.Unlock()
	if _, ok := s.clients[id]; ok {
		return nil, errors.New("client already registered")
	}
	cancel := func() {
		s.mut.Lock()
		defer s.mut.Unlock()
		delete(s.clients, id)
	}
	s.clients[id] = ch
	return cancel, nil
}

func (s *sseSource) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.RequestURI != "/gallium.events" {
		s.handler.ServeHTTP(w, r)
		return
	}

	f, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "unable to stream", http.StatusInternalServerError)
		return
	}

	id := s.clientID(r)
	ch := make(chan sseEvent)
	cancel, err := s.register(id, ch)
	if err != nil {
		http.Error(w, "failed to register client", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")

	ctx := r.Context()
	for {
		select {
		case e := <-ch:
			e.WriteTo(w)
			f.Flush()
		case <-ctx.Done():
			cancel()
			close(ch)
			for e := range ch {
				e.WriteTo(w)
			}
			return
		}
	}
}

func (s *sseSource) clientID(r *http.Request) string {
	var id [32]byte
	_, err := rand.Read(id[:])
	if err != nil {
		panic(err)
	}
	return base64.RawURLEncoding.EncodeToString(id[:])
}
