package main

import (
	"log"
	"practice/p2p"
)

func main() {
	tcpOpts := p2p.TCPTransportOpts{
		ListenAddr:    ":3000",
		Decoder:       p2p.DefaultDecoder{},
		HandshakeFunc: p2p.NOPHandshakeFunc,
	}
	tr := p2p.NewTCPTransport(tcpOpts)

	err := tr.ListenAndAccept()
	if err != nil {
		log.Panic(err)
	}

	select {}
}
