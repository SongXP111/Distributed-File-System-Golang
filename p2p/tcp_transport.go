package p2p

import (
	"fmt"
	"net"
	"sync"
)

// TCPTransport represents the TCP-based transport layer that handles network communication.
// TCPTransport 表示基于TCP的传输层，负责网络通信。
type TCPTransport struct {
	// The address on which the transport listens for incoming connections.
	// 传输层监听传入连接的地址。
	listenAddress string
	// TCP listener that accepts incoming connections.
	// 用于接受传入连接的TCP监听器。
	listener net.Listener
	// Mutex to ensure safe access to shared resources (peers map).
	// 读写锁，用于保证对共享资源（peers映射）的安全访问。
	mu sync.RWMutex
	// Map to store connected peers, indexed by their network address.
	// 用于存储已连接的节点（Peer），键为网络地址。
	peers map[net.Addr]Peer
}

// NewTCPTransport creates a new instance of TCPTransport with the specified listening address.
// NewTCPTransport 创建一个新的 TCPTransport 实例，并指定监听地址。
func NewTCPTransport(listenAddr string) *TCPTransport {
	return &TCPTransport{
		listenAddress: listenAddr, // Initialize the listen address. 初始化监听地址。
	}
}

// ListenAndAccept starts the TCP listener and accepts incoming connections.
// This function creates a TCP listener and launches a goroutine to handle connections asynchronously.
// ListenAndAccept 启动TCP监听器，并接受传入的连接。
// 此函数创建一个TCP监听器，并启动一个goroutine来异步处理连接。
func (t *TCPTransport) ListenAndAccept() error {
	var err error

	// Start listening on the provided address.
	// 在提供的地址上开始监听。
	t.listener, err = net.Listen("tcp", t.listenAddress)
	if err != nil {
		return err // Return error if listener fails to start. 如果监听器启动失败，则返回错误。
	}

	// Start a goroutine to asynchronously accept connections.
	// 启动一个goroutine来异步接受连接。
	go t.startAcceptLoop()
	return nil
}

// startAcceptLoop continuously accepts incoming TCP connections in a loop.
// startAcceptLoop 持续循环接受传入的TCP连接。
func (t *TCPTransport) startAcceptLoop() {
	for {
		// Accept an incoming connection. 接受传入的连接。
		conn, err := t.listener.Accept()
		if err != nil {
			// Print error message if the accept fails. 如果接受连接失败，则打印错误信息。
			fmt.Printf("TCP accept error: %s\n", err)
			continue // Continue to the next iteration. 继续下一个循环。
		}
		// Handle the accepted connection in a separate goroutine.
		// 在一个单独的goroutine中处理已接受的连接。
		go t.handleConn(conn)
	}
}

// handleConn handles an individual TCP connection once accepted.
// handleConn 处理每个已接受的TCP连接。
func (t *TCPTransport) handleConn(conn net.Conn) {
	// Print a message indicating a new connection has been established.
	// 打印一条消息，表明建立了一个新的连接。
	fmt.Printf("new incoming connection %+v\n", conn)

	// Here you would typically add the peer to the peers map and manage the connection.
	// 在这里，你通常会将节点添加到peers映射中，并管理这个连接。
}
