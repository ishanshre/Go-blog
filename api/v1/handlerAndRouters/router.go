package handlerAndRouters

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ishanshre/Go-blog/api/v1/db"
)

type ApiServer struct {
	listenAddr string
	store      db.Storage
}

func NewApiServer(listenAddr string, store db.Storage) *ApiServer {
	return &ApiServer{
		listenAddr: listenAddr,
		store:      store,
	}
}

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	return r
}

func (s *ApiServer) Run() {
	// initializing new gorilla mux router and running golang server
	router := NewRouter()
	userRouter(router, s)
	tagRouter(router, s)
	postRouter(router, s)
	mediaRouter(router, s)
	commentRouter(router, s)
	log.Println("Starting server at port ", s.listenAddr)
	http.ListenAndServe(s.listenAddr, router)
}
