package handlerAndRouters

import (
	"github.com/gorilla/mux"
	"github.com/ishanshre/Go-blog/api/v1/middlewares"
)

func postRouter(r *mux.Router, s *ApiServer) {
	r.HandleFunc("/api/v1/post", middlewares.JwtAuthHandler(middlewares.MakeHttpHandler(s.handlePostCreate), s.store))
	r.HandleFunc("/api/v1/posts", middlewares.MakeHttpHandler(s.handlePostAll))
	r.HandleFunc("/api/v1/post/{id}", middlewares.MakeHttpHandler(s.handlePostGetById))
	r.HandleFunc("/api/v1/post/{id}/delete", middlewares.JwtAuthPostOwnerPermissionHandler(middlewares.MakeHttpHandler(s.handlePostDeleteById), s.store))
	r.HandleFunc("/api/v1/post/{id}/update", middlewares.JwtAuthPostOwnerPermissionHandler(middlewares.MakeHttpHandler(s.handlePostUpdateById), s.store))
	r.HandleFunc("/api/v1/post/{id}/tag", middlewares.JwtAuthPostOwnerPermissionHandler(middlewares.MakeHttpHandler(s.handlePostTagAddDelete), s.store))
	r.HandleFunc("/api/v1/post/{id}/tags", middlewares.MakeHttpHandler(s.handlePostTagsAll))
}
