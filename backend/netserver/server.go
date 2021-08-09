package netserver

import (
	"fmt"
	"log"
	"net"
)

func Start() {
	// Only localhost
	addr := "127.0.0.1:8080"

	// Open TCP Socket
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}

	l := struct{ listner net.Listener }{ln}
	defer l.listner.Close()

	// TODO: HTTP2
	// TODO: Track
	// TODO: Context
	// TODO: TLS

	for {
		// Accept reqest
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}

		// curl -I localhost:8080
		// HTTP/1.1 200 OK
		// Content-Type: application/json
		// Date: Mon, 09 Aug 2021 09:33:03 GMT
		// Content-Length: 2
		fmt.Fprintf(conn, "HTTP/1.1 200 OK\r\n")
		fmt.Fprintf(conn, "Content-Type: application/json\r\n")
		fmt.Fprintf(conn, "Date: Mon, 09 Aug 2021 09:33:03 GMT\r\n")
		fmt.Fprintf(conn, "Content-Length: 2\r\n")
		fmt.Fprintf(conn, "\r\n")
		fmt.Fprintf(conn, "ok\r\n")
		conn.Close()
	}
}
