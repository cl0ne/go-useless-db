package main

import (
	"io"
	"log"
	"net"
)

func handleConnection(c net.Conn) {
	log.Println("Accepted connection from:", c.RemoteAddr())
	io.Copy(c, c)
	log.Println("Client served:", c.RemoteAddr())
	c.Close()
}

func main() {
	listenAddress := ":1337"
	l, err := net.Listen("tcp", listenAddress)
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	log.Println("Serving on", l.Addr())
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handleConnection(conn)
	}
}
