package nexus

import (
	"errors"
	"github.com/google/uuid"
	"log"
	"net/rpc"
	"time"
)

type Client struct {
	ID      string
	handler *Handler[API]
	client  *rpc.Client
}

func NewClient(network string, address string) *Client {
	return &Client{
		ID:      uuid.New().String(),
		handler: NewHandler[API](network, address),
	}
}

func (a *Client) Start(channel chan string) {
	var err error
	a.client, err = a.handler.StartClient()
	if err != nil {

	}
	go a.keepClientAlive(channel)
	go a.wireChannel(channel)
}

func (a *Client) wireChannel(channel chan string) {
	for msg := range channel {
		log.Printf("[nexus] Outputting message [%s]\n", msg)
		// Send them out the remote
		a.ProcessPacket(msg)
	}
}

func (a *Client) keepClientAlive(channel chan string) {
	for {
		time.Sleep(time.Second)
		_, err := a.Ping("Is this connection active?")
		if err != nil {
			log.Println("[nexus] Client disconnected, reconnecting")
			break
		}
	}
	a.Start(channel)
}

func (a *Client) Ping(value string) (string, error) {
	var str string
	if a.client == nil {
		return str, errors.New("Client is nil")
	}
	err := a.client.Call("API.Ping", value, &str)
	if err != nil {
		log.Println("[nexus] Error: ", err)
		return str, err
	}
	return str, nil
}

func (a *Client) ProcessPacket(value string) {
	var str string
	if a.client != nil {
		err := a.client.Call("API.ProcessPacket", value, &str)
		if err != nil {
			log.Println("[nexus] Error: ", err)
		}
	}
}
