package netsrv

import (
	"log"
	"net"
)

const (
	proto = "tcp4"
	addr  = "0.0.0.0:8000"
)

func Start() {
	listener, err := net.Listen(proto, addr)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

}
