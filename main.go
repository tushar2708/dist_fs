package main

import (
	"fmt"
	"log"

	"github.com/tushar2708/dist_fs/p2p"
)

func main() {
	fmt.Println("distributed file-system")

	tr := p2p.NewTCPTransport(":3000")

	if err := tr.ListenAndAccept(); err != nil {
		log.Fatalf("failed to listen and accept, err: %v", err)
	}

	select {}

}
