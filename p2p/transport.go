package p2p

// Peer is an interface that represents the remote node
// Peer代表一个远程节点，通常在对等网络（P2P）中，节点之间会互相通信
type Peer interface {
}

// Transport is anything that handles communication
// between the nodes in the network. This can be the form
// (TCP, UDP, websockets, ...)
// Transport 通常用于处理节点之间的通信
type Transport interface {
	ListenAndAccept() error
}
