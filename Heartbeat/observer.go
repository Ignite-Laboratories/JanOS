package main

import (
	"log"
)

type Observer struct {
	Stream chan string
}

func NewObserver() *Observer {
	m := &Observer{
		Stream: make(chan string),
	}
	go m.Write()

	return m
}

func (m Observer) Write() {
	for msg := range m.Stream {
		log.Println(msg)
	}
}

func (m Observer) Observe(source chan string) {
	for msg := range source {
		m.Stream <- msg
	}
}
