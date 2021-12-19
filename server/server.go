package server

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/hashicorp/go-hclog"
)

func configure(shopLog hclog.Logger, router *mux.Router, port string) *http.Server {
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})
	ch := handlers.CORS(headersOk, originsOk, methodsOk)

	// create a new server
	s := http.Server{
		Addr:         fmt.Sprintf(":%s", port),                               // configure the bind address
		Handler:      ch(router),                                             // set the default handler
		ErrorLog:     shopLog.StandardLogger(&hclog.StandardLoggerOptions{}), // set the logger for the server
		ReadTimeout:  5 * time.Second,                                        // max time to read request from the client
		WriteTimeout: 10 * time.Second,                                       // max time to write response to the client
		IdleTimeout:  120 * time.Second,                                      // max time for connections using TCP Keep-Alive
	}
	return &s
}

func Start(shopLog hclog.Logger, port string) {
	router := createRouter(shopLog)
	server := configure(shopLog, router, port)
	// start the server
	go func() {
		shopLog.Info("Starting server", "port", port)

		err := server.ListenAndServe()
		if err != nil {
			shopLog.Error("Error starting server", "error", err)
			os.Exit(1)
		}
	}()

	// trap sigterm or interupt and gracefully shutdown the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)

	// Block until a signal is received.
	sig := <-c
	shopLog.Info("Signal received", "signal", sig)

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	server.Shutdown(ctx)
}

func createRouter(shopLog hclog.Logger) *mux.Router {
	router := mux.NewRouter()
	getR := router.Methods(http.MethodGet).Subrouter()
	getR.HandleFunc("/products", func(rw http.ResponseWriter, r *http.Request) {
		shopLog.Info("[GET] Request received")

		_ = json.NewEncoder(rw).Encode(nil)
	})
	return router
}
