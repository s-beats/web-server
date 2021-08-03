package netserver

import (
	"log"
	"net"
	"net/http"
)

func Start() {
	addr := ":8080"
	// handler := func(w http.ResponseWriter, r *http.Request) {}

	// Open TCP Socket
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	for {
		// Accept reqest
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}

		// curl -I localhost:8080
		// HTTP/0.0 200 OK
		// Content-Length: 0
		res := http.Response{
			StatusCode: 200,
		}
		res.Write(conn)
		conn.Close()
	}
}
