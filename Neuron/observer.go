package main

import (
	"bufio"
	"log"
	"os"
)

type Observer struct {
	OutputStream chan string
}

func NewStdInObserver() *Observer {
	m := &Observer{
		OutputStream: make(chan string),
	}
	scanner := bufio.NewScanner(os.Stdin)
	go m.handleScanner(scanner)
	return m
}

func (m Observer) handleScanner(scanner *bufio.Scanner) {
	for scanner.Scan() {
		m.OutputStream <- scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func (m Observer) MuxChannel(source chan string) {
	for msg := range source {
		m.OutputStream <- msg
	}
}
