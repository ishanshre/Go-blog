package handlerAndRouters

import (
	"github.com/gorilla/mux"
	"github.com/ishanshre/Go-blog/api/v1/middlewares"
)

func userRouter(r *mux.Router, s *ApiServer) {
	r.HandleFunc("/api/v1/auth/signup", middlewares.MakeHttpHandler(s.handleUserSignUp))
	r.HandleFunc("/api/v1/auth/login", middlewares.MakeHttpHandler(s.handleUserLogin))
	r.HandleFunc("/api/v1/auth/user/{id}", middlewares.JwtAuthHandler(middlewares.MakeHttpHandler(s.handleUserInfoById), s.store))
	r.HandleFunc("/api/v1/auth/users", middlewares.JwtAuthAdminHandler(middlewares.MakeHttpHandler(s.handleUsersAll), s.store))
}
