package handlerAndRouters

import (
	"github.com/gorilla/mux"
	"github.com/ishanshre/Go-blog/api/v1/middlewares"
)

func postRouter(r *mux.Router, s *ApiServer) {
	r.HandleFunc("/api/v1/post", middlewares.JwtAuthHandler(middlewares.MakeHttpHandler(s.handlePostCreate), s.store))
	r.HandleFunc("/api/v1/posts", middlewares.MakeHttpHandler(s.handlePostAll))
	r.HandleFunc("/api/v1/post/{slug}", middlewares.MakeHttpHandler(s.handlePostGetBySlug))
}
