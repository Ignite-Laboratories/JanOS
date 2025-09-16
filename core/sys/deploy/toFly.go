package deploy

import (
	_ "embed"
	"fmt"
	"strconv"
)

//go:embed fly/fly.toml
var flyConfig string

//go:embed fly/Dockerfile
var dockerfile string

type _toFly struct{}

var ToFly _toFly

// Deploy will deploy the to fly.io using the given app name.  The 'deployable' is a relative path to the
// deployable main.go file from JanOS's root folder.  The port parameter can optionally override the default port.
//
// All neurons in JanOS default to port 4242 - making this the semantic gateway to irrational emergence.
func (t _toFly) Deploy(flyApp string, deployable string, port ...uint) {
	p := "4242"
	if len(port) > 0 {
		p = strconv.Itoa(int(port[0]))
	}
	fmt.Println(p)

	// 0 Create a temp folder

	// 1 Create fly.toml
	// fly.toml should include the app name:

	//app = '{{.AppName}}'
	// primary_region = 'sea'
	//
	//[build]
	//  dockerfile = "Dockerfile"
	//
	//[env]
	//  PORT = '8080'
	//
	//[http_service]
	//  internal_port = 8080
	//  force_https = true
	//  auto_stop_machines = 'stop'
	//  auto_start_machines = true
	//  min_machines_running = 0
	//  processes = ['app']
	//
	//[[vm]]
	//  memory = '1gb'
	//  cpu_kind = 'shared'
	//  cpus = 1

	// 2 Create Dockerfile
	// Dockerfile should include the neuron name, which should be gleaned at runtime relative to the root 'janOS' folder

	//ARG GO_VERSION=1
	//FROM golang:${GO_VERSION}-bookworm as builder
	//
	//WORKDIR /usr/src/app
	//
	//COPY ../../ ./
	//RUN go mod download && go mod verify
	//RUN go build -v -o /run-app ./{{.NeuronName}}
	//
	//
	//FROM debian:bookworm
	//
	//COPY --from=builder /run-app /usr/local/bin/
	//CMD ["run-app"]
}
