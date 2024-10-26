package p2p

// HandshakeFunc is the function that performs the handshake
type HandshakeFunc func(Peer) error

func NoOPHandShakeFunc(Peer) error {
	return nil
}
