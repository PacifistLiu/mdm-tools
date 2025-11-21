package crate

import (
	"bufio"
	"io"
	"log"
)

// ReadLines from io.Reader, return a slice of lines
func ReadLines(r io.Reader) []string {
	var buf []string

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		buf = append(buf, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal("reading line from STDIN, err")
	}

	return buf
}
