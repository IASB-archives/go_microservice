package main

import (
	"log"
	"os"
	"time"

	"github.com/fun-to-projects/go_microservice/data"
	"github.com/fun-to-projects/go_microservice/server"
	"github.com/hashicorp/go-hclog"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(data.ENV_FILE_PATH)
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	serviceLogger := configureLogger(os.Getenv(data.SERVICE_NAME))
	server.Start(serviceLogger, os.Getenv(data.SERVICE_PROD_PORT))
}

func configureLogger(appName string) hclog.Logger {
	return hclog.New(&hclog.LoggerOptions{
		Name:       appName,
		Output:     os.Stdout,
		JSONFormat: false,
		TimeFormat: time.RFC822,
	})
}
