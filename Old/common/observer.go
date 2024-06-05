package common

import (
	"bufio"
	"log"
	"os"
)

type Observer struct {
	OutputStream chan string
}

func NewObserver(sources ...chan string) *Observer {
	o := &Observer{OutputStream: make(chan string)}

	for _, source := range sources {
		go o.ReceiveFrom(source)
	}

	return o
}

func NewStdInObserver() *Observer {
	m := NewObserver()
	s := bufio.NewScanner(os.Stdin)

	go func(scanner *bufio.Scanner) {
		for scanner.Scan() {
			m.OutputStream <- scanner.Text()
		}
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}(s)

	return m
}

func (o *Observer) ForwardTo(output chan string) {
	for data := range o.OutputStream {
		output <- data
	}
}

func (o *Observer) ReceiveFrom(source chan string) {
	for msg := range source {
		o.OutputStream <- msg
	}
}
