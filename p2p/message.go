package p2p

import "net"

// Message represents any arbitrary data that is being sent over
// each transport between two nodes
type Message struct {
	From    net.Addr
	Payload []byte
}
