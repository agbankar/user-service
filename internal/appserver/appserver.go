package appserver

import (
	"context"
	"github.com/gorilla/mux"
	"log"
	"mocking-goway/cmd/webservices/controllers"
	"mocking-goway/internal/config"
	"net/http"
	"time"
)

type AppServer interface {
	Start()
	Stop()
}
type HttpServer struct {
	config *config.Config
	router *mux.Router
	server *http.Server
}

func NewAppServer(config *config.Config) *HttpServer {
	return &HttpServer{
		config: config,
		router: mux.NewRouter(),
	}
}
func (a *HttpServer) Start() {
	a.router.HandleFunc("/getData", controllers.GetBonus).Methods("GET")
	a.server = &http.Server{
		Addr:    a.config.Port,
		Handler: a.router,
	}
	if err := a.server.ListenAndServe(); err != nil {
		log.Fatal("Not able to start the server")
	}
}
func (a HttpServer) Stop() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Nanosecond)
	defer cancel()
	a.server.Shutdown(ctx)
}
