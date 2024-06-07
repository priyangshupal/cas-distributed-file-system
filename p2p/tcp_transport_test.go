package p2p

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTCPTransport (t *testing.T) {
	opts := TCPTransportOpts {
		ListenAddr: ":3000",
		HandshakeFunc: NOPHandshakeFunc,
		Decoder: DefaultDecoder{},
	}
	
	listenAddr := ":3000"
	tr := NewTCPTransport(opts)

	assert.Equal(t, tr.ListenAddr, listenAddr)

	assert.Nil(t, tr.ListenAndAccept())
}