package handlerAndRouters

import (
	"github.com/gorilla/mux"
	"github.com/ishanshre/Go-blog/api/v1/middlewares"
)

func mediaRouter(r *mux.Router, s *ApiServer) {
	r.HandleFunc("/media/image/{filename}", middlewares.MakeHttpHandler(s.handleMediaImage))
}
