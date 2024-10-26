package p2p

type Peer interface {
}

// Transport can be anything,
// that handles the communication b/w nodes in the network
type Transport interface {
	ListenAndAccept() error
}
