package p2p

import "testing"

func TestNewP2P(t *testing.T) {
	p := NewP2P("123")
	p.makeBasicHost(8000)
}
