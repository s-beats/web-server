package netserver

import (
	"log"
	"net"

	"github.com/go-pp/pp"
)

func Start() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	rw, err := ln.Accept()
	if err != nil {
		log.Fatal(err)
	}

	// TODO: net.TCPConnを操作してresponseを作る
	pp.Println(rw)
}
