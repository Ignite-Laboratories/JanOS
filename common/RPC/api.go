package RPC

import "net/rpc"

type RPC interface {
	StartServer()
	StartClient() *rpc.Client
}
