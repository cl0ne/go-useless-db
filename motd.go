package main

import (
	"fmt"
	"io"
)

func motd(w io.Writer, args string) {
	fmt.Fprintln(w, "Welcome! Type 'help' to get available commands")
}
