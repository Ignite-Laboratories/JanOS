package support

import (
	"bufio"
	"log"
	"os"
)

type Observer struct {
	OutputStream chan string
}

func NewObserver() *Observer {
	return &Observer{OutputStream: make(chan string)}
}

func NewStdInObserver() *Observer {
	m := NewObserver()
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
