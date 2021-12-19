package router

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hashicorp/go-hclog"
)

type Router struct {
	Router    *mux.Router
	ApiLogger hclog.Logger
}

func (router Router) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	router.ApiLogger.Info("Request received")
	http.DefaultServeMux.ServeHTTP(rw, r)
}

func NewRouter(shopLog hclog.Logger) *Router {
	return &Router{
		Router:    mux.NewRouter(),
		ApiLogger: shopLog,
	}
}

func (router *Router) Configure() {
	router.initializeRoutes()
}

type Product struct {
	Name string
}

func (router *Router) initializeRoutes() {
	getR := router.Router.Methods(http.MethodGet).Subrouter()
	getR.HandleFunc("/products", func(rw http.ResponseWriter, r *http.Request) {
		router.ApiLogger.Info("[GET] Request received")
		product := &Product{Name: "hello"}
		_ = json.NewEncoder(rw).Encode(product)
	})
}
