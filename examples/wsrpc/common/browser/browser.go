package browser

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"

	"github.com/gopherjs/websocket"
)

//Browser represents RPC of browser side
type Browser struct {
	Client *rpc.Client
	s      net.Conn
	c      net.Conn
}

//New connects websocket and returns Browser obj.
func New(dest string, strs ...interface{}) (*Browser, error) {
	var err error
	b := &Browser{}
	for _, str := range strs {
		if err := rpc.Register(str); err != nil {
			return nil, err
		}
	}
	b.s, err = websocket.Dial("ws://" + dest + "/ws-client") // Blocks until connection is established
	if err != nil {
		return nil, err
	}
	log.Println("ws-client connected")
	go jsonrpc.ServeConn(b.s)

	b.c, err = websocket.Dial("ws://" + dest + "/ws-server") // Blocks until connection is established
	if err != nil {
		return nil, err
	}
	log.Println("client connected")
	b.Client = jsonrpc.NewClient(b.c)
	return b, nil
}

//Call Calls calls RPC.
func (b *Browser) Call(m string, args interface{}, reply interface{}) error {
	return b.Client.Call(m, args, reply)
}

//Close closes RPC client and server connections.
func (b *Browser) Close() {
	if err := b.c.Close(); err != nil {
		log.Println(err)
	}
	if err := b.s.Close(); err != nil {
		log.Println(err)
	}
}
