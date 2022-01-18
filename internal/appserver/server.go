package appserver

import (
	"github.com/gorilla/mux"
	"mocking-goway/internal/config"
	"net/http"
)

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
func (a *AppServer) StartServer() {
	a.router.HandleFunc("/getData", GetData).Methods("GET")
	a.server = &http.Server{
		Addr:    a.config.Port,
		Handler: a.router}
	a.router.Handle("/", a.router)
	a.server.ListenAndServe()

}

func GetData(writer http.ResponseWriter, request *http.Request) {

}
