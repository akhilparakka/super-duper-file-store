package p2p

type Transport interface {
	ListenAndAccept() error
	Consume() <-chan RPC
}

type Peer interface {
	Close() error
}
