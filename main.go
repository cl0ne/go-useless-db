package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"time"

	"github.com/cl0ne/go-useless-db/database"
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

var db = database.NewDatabase()
var dbPath = "useless-db.json"

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

func save() {
	db_file, err := os.OpenFile(dbPath, os.O_CREATE|os.O_WRONLY, 0)
	if err != nil {
		log.Panic("Failed to open database file for saving")
	}
	defer db_file.Close()

	saved, err := db.Serialize(db_file)
	if err != nil {
		log.Println("Failed to save database:", err)
		return
	}

	if !saved {
		return
	}

	log.Println("Modified database saved")
	size, err := db_file.Seek(0, io.SeekCurrent)
	if err != nil {
		log.Println("Database file seek failed:", err)
		return
	}

	db_file.Truncate(size)
}

func autoSave(seconds time.Duration) {
	tickChan := time.NewTicker(time.Second * seconds).C
	run := true
	for run {
		<-tickChan
		log.Println("Running autosave task")

		save()
	}
	log.Println("Exiting autosave task")
}

func main() {
	db_file, err := os.Open(dbPath)
	if err == nil {
		log.Println("Loading data from file...")
		db.Deserialize(db_file)
	} else if !os.IsNotExist(err) {
		log.Fatal("Failed to load database from file")
	}

	listenAddress := ":1337"
	l, err := net.Listen("tcp", listenAddress)
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	go autoSave(2)
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
