package p2p

import (
	"net"
)

// Peer is an interface that represents remote node
type Peer interface {
	net.Conn
	Send ([]byte) error
	CloseStream()
}

// Transport is anything that handles communication
// between network nodes. This can be TCP, UDP, WebSockets...
type Transport interface {
	Addr() string
	Dial(string) error
	ListenAndAccept() error
	Consume() <- chan RPC
	Close() error
}