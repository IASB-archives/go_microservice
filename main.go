package main

import (
	"github.com/fun-to-projects/go_microservice/server"
	"github.com/hashicorp/go-hclog"
)

func main() {
	serviceLogger := hclog.Default()
	server.Start(serviceLogger)
}
