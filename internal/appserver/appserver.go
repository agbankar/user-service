package appserver

import (
	"context"
	"github.com/gorilla/mux"
	"log"
	"mocking-goway/internal/config"
	"net/http"
	"time"
)

type Server interface {
	Start()
	Stop()
}

type AppServer struct {
	config *config.Config
	router *mux.Router
	server *http.Server
}

func NewAppServer(config *config.Config) *AppServer {
	return &AppServer{
		config: config,
		router: mux.NewRouter(),
	}
}
func (a *AppServer) Start() {
	a.server = &http.Server{
		Addr:    ":" + a.config.Port,
		Handler: a.router,
	}
	if err := a.server.ListenAndServe(); err != nil {
		log.Fatal("Not able to start the server")
	}
}

func (a *AppServer) AddGetApi(Path string, HandledFunc func(w http.ResponseWriter, r *http.Request)) {
	a.router.HandleFunc(Path, HandledFunc)
}

func (a *AppServer) Stop() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	a.server.Shutdown(ctx)
}
