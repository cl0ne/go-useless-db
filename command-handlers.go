package main

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"github.com/cl0ne/go-useless-db/database"
)

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func parseRecord(w io.Writer, r *database.Record, j string) (ok bool) {
	err := json.Unmarshal([]byte(j), r)
	ok = (err == nil)
	if !ok {
		fmt.Fprintln(w, "Failed to parse record:", err)
	}
	return
}

func handleInsert(w io.Writer, args string) {
	values := strings.SplitN(args, " ", 2)

	if len(values) < 2 {
		fmt.Fprintln(w, "Error: insert requires 2 arguments,", min(1, len(args)), "given.")
		return
	}

	var index int
	_, err := fmt.Sscan(values[0], &index)
	if err != nil {
		fmt.Fprintln(w, "Failed to get index:", err)
		return
	}

	var r database.Record
	if !parseRecord(w, &r, values[1]) {
		return
	}

	ok := db.Insert(index, r)
	if ok {
		fmt.Fprintln(w, "OK")
	} else {
		fmt.Fprintf(w, "Error: %d is not valid index.\n", index)
	}
}

func handleGet(w io.Writer, args string) {
	if len(args) == 0 {
		fmt.Fprintln(w, "Error: get requires 1 argument, none given")
		return
	}
	var index int
	_, err := fmt.Sscan(args, &index)
	if err != nil {
		fmt.Fprintln(w, args, "is not a valid index:", err)
		return
	}

	r, ok := db.Get(index)
	if !ok {
		fmt.Fprintf(w, "Error: %d is not valid index.\n", index)
		return
	}

	b, err := json.Marshal(r)
	if err != nil {
		fmt.Fprintln(w, "Failed to serialize record", r, ":", err)
		return
	}

	w.Write(b)
	fmt.Fprintln(w)
}

func handleUpdate(w io.Writer, args string) {
	values := strings.SplitN(args, " ", 2)

	if len(values) < 2 {
		fmt.Fprintln(w, "Error: update requires 2 arguments,", min(1, len(args)), "given.")
		return
	}

	var index int
	_, err := fmt.Sscan(values[0], &index)
	if err != nil {
		fmt.Fprintln(w, "Failed to get index:", err)
		return
	}
	var r database.Record
	if !parseRecord(w, &r, values[1]) {
		return
	}
	ok := db.Update(index, r)
	if !ok {
		fmt.Fprintf(w, "Error: %d is not valid index.\n", index)
		return
	}
	fmt.Fprintln(w, "OK")
}

func handleLen(w io.Writer, args string) {
	fmt.Fprintln(w, "The length of DB:", db.Length())
}

func handleDelete(w io.Writer, args string) {
	if len(args) == 0 {
		fmt.Fprintln(w, "Error: get requires 1 argument, none given")
		return
	}

	var index int
	_, err := fmt.Sscan(args, &index)
	if err != nil {
		fmt.Fprintln(w, "Failed to get index:", err)
		return
	}

	ok := db.Remove(index)
	if !ok {
		fmt.Fprintln(w, "Invalid index")
		return
	}
	fmt.Fprintln(w, "OK")
}

func handleClear(w io.Writer, args string) {
	db.Clear()
	fmt.Fprintln(w, "DB is empty")
}
