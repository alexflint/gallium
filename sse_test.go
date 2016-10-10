package gallium

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSSECallsUnderlyingHandler(t *testing.T) {
	stub := &handlerStub{}
	s := newSSESource(stub)

	tests := []struct {
		method string
		path   string
	}{
		{method: http.MethodGet, path: "/"},
		{method: http.MethodGet, path: "/gallium"},
		{method: http.MethodHead, path: "/"},
		{method: http.MethodGet, path: "/sub/path"},
		{method: http.MethodPost, path: "/sub/path"},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%s %s", test.method, test.path), func(t *testing.T) {
			w, r := httptest.NewRecorder(), httptest.NewRequest(test.method, test.path, nil)
			s.ServeHTTP(w, r)
			if w.Code != http.StatusOK {
				t.Fatalf("Request failed: %d", w.Code)
			}

			req := stub.lastRequest()
			if req.Method != test.method {
				t.Fatalf("Expected method: %s Actual method: %s", test.method, req.Method)
			}
			if req.RequestURI != test.path {
				t.Fatalf("Expected method: %s Actual method: %s", test.path, req.RequestURI)
			}
		})
	}
}

func TestSSEFlushesOnStart(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	s := newSSESource(nil)
	w, r := httptest.NewRecorder(), httptest.NewRequest(http.MethodGet, "/gallium.events", nil)
	r = r.WithContext(ctx)
	cancel()
	s.ServeHTTP(w, r)
	if w.Code != http.StatusOK {
		t.Errorf("Failed request: %d", w.Code)
	}
	if ct := w.Header().Get("Content-Type"); ct != "text/event-stream" {
		t.Errorf("Unexpected Content Type: %s", ct)
	}
	if w.Body.String() != "\n" {
		t.Errorf(`
		Expected body: %q
		Received body: %q
		`, "\n", w.Body.String())
	}
	if !w.Flushed {
		t.Error("Connection must be flushed when started")
	}
}

func TestReceiveEvents(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	s := newSSESource(nil)
	w, r := httptest.NewRecorder(), httptest.NewRequest(http.MethodGet, "/gallium.events", nil)
	r = r.WithContext(ctx)
	go func() {
		s.Emit("first", struct {
			Foo string `json:"foo"`
		}{Foo: "bar"})
		s.Emit("second", `second event`)
		cancel()
	}()

	var expected bytes.Buffer
	fmt.Fprintln(&expected, ``)
	fmt.Fprintln(&expected, "event: first")
	fmt.Fprintln(&expected, `data: {"foo":"bar"}`)
	fmt.Fprintln(&expected, ``)
	fmt.Fprintln(&expected, "event: second")
	fmt.Fprintln(&expected, `data: "second event"`)
	fmt.Fprintln(&expected, ``)

	s.ServeHTTP(w, r)
	s.Emit("missed", `miss this event`)

	if w.Code != http.StatusOK {
		t.Errorf("Failed request: %d", w.Code)
	}
	if w.Body.String() != expected.String() {
		t.Errorf(`
		Expected body: %q
		Received body: %q
		`, expected.String(), w.Body.String())
	}
	if ct := w.Header().Get("Content-Type"); ct != "text/event-stream" {
		t.Errorf("Unexpected Content Type: %s", ct)
	}
}

type handlerStub struct {
	reqs []*http.Request
}

func (h *handlerStub) lastRequest() *http.Request {
	return h.reqs[len(h.reqs)-1]
}

func (h *handlerStub) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.reqs = append(h.reqs, r)
}
