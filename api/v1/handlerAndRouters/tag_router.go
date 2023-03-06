package handlerAndRouters

import (
	"github.com/gorilla/mux"
	"github.com/ishanshre/Go-blog/api/v1/middlewares"
)

func tagRouter(r *mux.Router, s *ApiServer) {
	r.HandleFunc("/api/v1/tags", middlewares.JwtAuthHandler(middlewares.MakeHttpHandler(s.handleTag), s.store))
	r.HandleFunc("/api/v1/tag/{id}", middlewares.JwtAuthAdminHandler(middlewares.MakeHttpHandler(s.handleTagsById), s.store))
}
