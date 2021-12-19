package main

import (
	"github.com/fun-to-projects/go_microservice/router"
	"github.com/fun-to-projects/go_microservice/server"
	"github.com/hashicorp/go-hclog"
)

func main() {
	serviceLogger := hclog.Default()
	router := router.NewRouter(serviceLogger)
	router.Configure()
	server.Start(serviceLogger, router)
}
