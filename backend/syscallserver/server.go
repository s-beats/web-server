package syscallserver

import (
	"log"
	"net"
	"syscall"
)

func Start() {
	syscall.ForkLock.Lock()
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
		nfd, addr, err := syscall.Accept(fd)
		if err == nil {
			syscall.CloseOnExec(nfd)
		}
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("Sockaddr is %#v", addr)
		log.Printf("nfd is %#v", nfd)

		if _, err := syscall.Write(fd, []byte("response")); err != nil {
			log.Println("syscall.Write error")
			log.Fatal(err)
		}
		// FIXME: socket is not conncted
	}
}
