package syscallserver

import (
	"log"
	"net"
	"syscall"
)

// reference: https://gist.github.com/jschaf/93f37aedb5327c54cb356b2f1f0427e3
func Start() {
	// not affcted by other parallel process
	syscall.ForkLock.Lock()
	// domain is protocol family (IPv4)
	// type is communication
	// proto is sepecify socket type
	fd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, 0)
	if err != nil {
		log.Fatal(err)
	}
	syscall.ForkLock.Unlock()

	sa := syscall.SockaddrInet4{Port: 8080}
	ip := net.ParseIP("127.0.0.1")
	copy(sa.Addr[:], ip)
	if err := syscall.Bind(fd, &sa); err != nil {
		log.Fatalln(err)
	}

	if err := syscall.Listen(fd, syscall.SOMAXCONN); err != nil {
		log.Fatal(err)
	}
	defer syscall.Close(fd)

	for {
		// Accept reqest
		nfd, _, err := syscall.Accept(fd)
		if err == nil {
			syscall.CloseOnExec(nfd)
		}
		if err != nil {
			log.Fatal(err)
		}

		b := make([]byte, 1024)
		n, err := syscall.Read(nfd, b)
		if err != nil {
			log.Println("syscall.Read error")
			log.Fatal(err)
		}
		log.Printf("read %d bytes\n", n)
		log.Println(string(b))

		res := "HTTP/1.1 200 OK\r\nContent-Type: application/json\r\nDate: Mon, 09 Aug 2021 09:33:03 GMT\r\nContent-Length: 2\r\n\r\nok\r\n"
		if _, err := syscall.Write(nfd, []byte(res)); err != nil {
			log.Println("syscall.Write error")
			log.Fatal(err)
		}
	}
}
