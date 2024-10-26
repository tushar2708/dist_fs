package p2p

import (
	"fmt"
	"net"
	"sync"
)

// TCPPeer represtnts a remote node over a TCP established connection
type TCPPeer struct {
	// conn is the underlying TCP connection of the peer
	conn net.Conn

	// if we dial a connection => outbound = true
	// if we accept and retrieve a connection => outbound = false
	outboundPeer bool
}

func NewTCPPeer(conn net.Conn, outbound bool) *TCPPeer {
	return &TCPPeer{
		conn:         conn,
		outboundPeer: outbound,
	}
}

type TCPTransport struct {
	listenAddress string
	listener      net.Listener
	shakeHands    HandshakeFunc
	decoder       Decoder

	mu    sync.RWMutex
	peers map[net.Addr]Peer
}

func NewTCPTransport(listenAddr string) *TCPTransport {
	return &TCPTransport{
		shakeHands:    NoOPHandShakeFunc,
		listenAddress: listenAddr,
	}
}

func (t *TCPTransport) ListenAndAccept() error {

	var err error

	t.listener, err = net.Listen("tcp", t.listenAddress)
	if err != nil {
		return err
	}

	go t.startAcceptLoop()

	return nil
}

func (t *TCPTransport) startAcceptLoop() {

	for {
		conn, err := t.listener.Accept()
		if err != nil {
			fmt.Printf("TCP accept error: %v\n", err)
		}

		go t.handleConn(conn)
	}
}

type Temp struct {
}

func (t *TCPTransport) handleConn(conn net.Conn) {

	peer := NewTCPPeer(conn, false)

	if err := t.shakeHands(peer); err != nil {
		conn.Close()
		return
	}

	fmt.Printf("New peer connected: %v\n", peer)

	// lenDecodeError := 0
	msg := &Temp{}
	for {
		if err := t.decoder.Decode(conn, msg); err != nil {
			// lenDecodeError++
			// if lenDecodeError > 10 {
			// 	conn.Close()
			// }
			fmt.Printf("failed to decode, err: %v\n", err)
			continue
		}

	}

	// peer := NewPeer(conn, t)
	// t.addPeer(peer)
}
