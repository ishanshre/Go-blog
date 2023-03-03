package handlerAndRouters

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ishanshre/Go-blog/api/v1/db"
)

type ApiServer struct {
	ListenAddr string
	Store      db.Storage
}

func NewApiServer(listenAddr string, store db.Storage) *ApiServer {
	return &ApiServer{
		ListenAddr: listenAddr,
		Store:      store,
	}
}

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	return r
}

func (s *ApiServer) Run() {
	router := NewRouter()
	log.Println("Starting server at port ", s.ListenAddr)
	http.ListenAndServe(s.ListenAddr, router)
}
