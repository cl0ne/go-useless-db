package main

import (
	"fmt"
	"io"
)

func help(w io.Writer, args string) {
	fmt.Fprintln(w, "motd <index>")
	fmt.Fprintln(w, "   show MOTD again")
	fmt.Fprintln(w, "insert <index> <json>")
	fmt.Fprintln(w, "   insert record at index")
	fmt.Fprintln(w, "get <index>")
	fmt.Fprintln(w, "   get record at index")
	fmt.Fprintln(w, "update <index> <json>")
	fmt.Fprintln(w, "   update record at index")
	fmt.Fprintln(w, "delete <index>")
	fmt.Fprintln(w, "   delete record at index")
	fmt.Fprintln(w, "clear")
	fmt.Fprintln(w, "   clear all collection")
	fmt.Fprintln(w, "len")
	fmt.Fprintln(w, "   get collection's length")
	fmt.Fprintln(w, "exit")
	fmt.Fprintln(w, "   just exit from DB")
}
