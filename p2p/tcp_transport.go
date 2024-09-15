package p2p

import (
	"fmt"
	"net"
	"sync"
)

// TCPPeer represents the remote node over a TCP established connection.
type TCPPeer struct {
	// conn is the underlying connection of the peer
	conn net.Conn
	// if we dial and retrieve a conn -> outbound == true
	// if we accept and retrieve a conn -> outbound == false
	outbound bool
}

func NewTCPPeer(conn net.Conn, outbound bool) *TCPPeer {
	return &TCPPeer{
		conn:     conn,
		outbound: outbound,
	}
}

type TCPTransportOps struct {
	ListenAddr    string
	HandshakeFunc HandshakeFunc
	Decoder       Decoder
}

// TCPTransport represents the TCP-based transport layer that handles network communication.
// TCPTransport 表示基于TCP的传输层，负责网络通信。
type TCPTransport struct {
	TCPTransportOps

	listener net.Listener

	mu sync.RWMutex

	peers map[net.Addr]Peer
}

// NewTCPTransport creates a new instance of TCPTransport with the specified listening address.
// NewTCPTransport 创建一个新的 TCPTransport 实例，并指定监听地址。
func NewTCPTransport(opts TCPTransportOps) *TCPTransport {
	return &TCPTransport{
		TCPTransportOps: opts,
	}
}

// ListenAndAccept starts the TCP listener and accepts incoming connections.
// This function creates a TCP listener and launches a goroutine to handle connections asynchronously.
// ListenAndAccept 启动TCP监听器，并接受传入的连接。
// 此函数创建一个TCP监听器，并启动一个goroutine来异步处理连接。
func (t *TCPTransport) ListenAndAccept() error {
	var err error

	t.listener, err = net.Listen("tcp", t.ListenAddr)
	if err != nil {
		return err
	}

	go t.startAcceptLoop()
	return nil
}

// startAcceptLoop continuously accepts incoming TCP connections in a loop.
// startAcceptLoop 持续循环接受传入的TCP连接。
func (t *TCPTransport) startAcceptLoop() {
	for {
		conn, err := t.listener.Accept()
		if err != nil {
			fmt.Printf("TCP accept error: %s\n", err)
			continue
		}

		fmt.Printf("new incoming connection %+v\n", conn)
		go t.handleConn(conn)
	}
}

// handleConn handles an individual TCP connection once accepted.
// handleConn 处理每个已接受的TCP连接。
func (t *TCPTransport) handleConn(conn net.Conn) {
	peer := NewTCPPeer(conn, true)

	if err := t.HandshakeFunc(peer); err != nil {
		conn.Close()
		fmt.Printf("TCP handshake error: %s\n", err)
		return
	}

	// read loop
	msg := &Message{}
	for {
		if err := t.Decoder.Decode(conn, msg); err != nil {
			fmt.Printf("TCP error: %s\n", err)
			continue
		}

		msg.From = conn.RemoteAddr()

		fmt.Printf("message: %+v\n", msg)
	}

}
