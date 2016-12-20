package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

var handlers = map[string]func(w io.Writer, args string){
	"example": exampleHandler,
	"help":    help,
	"motd":    motd,
	"insert":  handleInsert,
	"update":  handleUpdate,
	"delete":  handleDelete,
	"get":     handleGet,
	"len":     handleLen,
	"clear":   handleClear,
}

var db = NewDatabase()

func handleConnection(c net.Conn) {
	log.Println("Accepted connection from:", c.RemoteAddr())
	motd(c, "")

	writer := bufio.NewWriter(c)
	defer func() {
		log.Println("Client served:", c.RemoteAddr())
		writer.Flush()
		c.Close()
	}()

	scanner := bufio.NewScanner(c)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			continue
		}

		var key string
		fmt.Sscanf(line, "%s", &key)
		key = strings.ToLower(key)

		if key == "exit" {
			fmt.Fprintln(writer, "Bye!")
			return
		}

		args := strings.TrimSpace(line[len(key):])

		h, ok := handlers[key]
		if !ok {
			continue
		}

		h(writer, args)
		writer.Flush()
	}
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
