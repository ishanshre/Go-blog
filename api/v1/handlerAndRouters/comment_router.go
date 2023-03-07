package handlerAndRouters

import (
	"github.com/gorilla/mux"
	"github.com/ishanshre/Go-blog/api/v1/middlewares"
)

func commentRouter(r *mux.Router, s *ApiServer) {
	r.HandleFunc("/api/v1/post/{id}/comment", middlewares.JwtAuthHandler(middlewares.MakeHttpHandler(s.handleCommentCreate), s.store))
	r.HandleFunc("/api/v1/post/{id}/comments", middlewares.MakeHttpHandler(s.handleCommentByPost))
	r.HandleFunc("/api/v1/post/{id}/comment/{comment_id}", middlewares.JwtAuthCommentOwnerPermissionHandler(middlewares.MakeHttpHandler(s.handleCommentUpdateDelete), s.store))
}
