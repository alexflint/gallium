package webserver

import (
	"log"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"

	"golang.org/x/net/websocket"
)

//WebServer represents web server side.
type WebServer struct {
	client *rpc.Client
	ch     chan struct{}
}

//New return WebServer obj.
func New(host string, strs ...interface{}) (*WebServer, error) {
	for _, str := range strs {
		if err := rpc.Register(str); err != nil {
			return nil, err
		}
	}
	w := &WebServer{
		ch: make(chan struct{}),
	}
	go w.start(host)
	<-w.ch
	return w, nil
}

//Close closes client RPC connection.
func (w *WebServer) Close() {
	w.ch <- struct{}{}
}

//Call calls calls RPC.
func (w *WebServer) Call(m string, args interface{}, reply interface{}) error {
	return w.client.Call(m, args, reply)
}

func (w *WebServer) start(host string) {
	http.HandleFunc("/ws-server",
		func(w http.ResponseWriter, req *http.Request) {
			s := websocket.Server{
				Handler: websocket.Handler(func(ws *websocket.Conn) {
					jsonrpc.ServeConn(ws)
				}),
			}
			s.ServeHTTP(w, req)
			log.Println("connected ws-server")
		})
	http.HandleFunc("/ws-client",
		func(rw http.ResponseWriter, req *http.Request) {
			s := websocket.Server{
				Handler: websocket.Handler(func(ws *websocket.Conn) {
					w.client = jsonrpc.NewClient(ws)
					w.ch <- struct{}{}
					<-w.ch
				}),
			}
			s.ServeHTTP(rw, req)
			log.Println("connected ws-client")
		})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, r.URL.Path[1:])
	})
	if err := http.ListenAndServe(host, nil); err != nil {
		log.Fatal(err)
	}
}
