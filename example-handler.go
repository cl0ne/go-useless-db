package main

import (
	"fmt"
	"io"
	"log"
)

func exampleHandler(w io.Writer, args string) {
	log.Printf("Example handler called, args: %q", args)
	fmt.Fprintf(w, "You called example handler with args: %q\n", args)
}
