package handlerAndRouters

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/ishanshre/Go-blog/api/v1/middlewares"
	"github.com/ishanshre/Go-blog/api/v1/models"
	"github.com/ishanshre/Go-blog/internals/pkg/utils"
)

func (s *ApiServer) handleUserSignUp(w http.ResponseWriter, r *http.Request) error {
	// hadner for user registration
	if r.Method == "POST" {
		req := new(models.RegisterUserRequest)
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			log.Println(err)
			return err
		}
		encPass, err := utils.HashPassword(req.Password)
		if err != nil {
			return err
		}
		user := models.RegisterNewUser(req.FirstName, req.LastName, req.Username, req.Email, encPass)
		if err := s.store.UserSignUp(user); err != nil {
			return err
		}
		return middlewares.WriteJSON(w, http.StatusOK, middlewares.ApiSuccess{Success: "user created"})
	}
	return middlewares.MethodNotAlowed(w, r.Method)
}

func (s *ApiServer) handleUserLogin(w http.ResponseWriter, r *http.Request) error {
	// handler for login process
	if r.Method == "POST" {
		log.Println("Post login ")
		req := new(models.LoginRequest)
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			log.Println(err)
			return err
		}
		user, err := s.store.UserLogin(req.Username)
		if err != nil {
			log.Println(err)
			return err
		}
		if err := utils.VerifyPassword(user.Password, req.Password); err != nil {
			log.Println(err)
			return err
		}
		res, err := utils.GenerateTokens(user.ID, req.Username)
		if err != nil {
			log.Println(err)
			return err
		}
		if err := s.store.UpdateLastLogin(user.ID); err != nil {
			log.Println(err)
			return err
		}
		return middlewares.WriteJSON(w, http.StatusOK, res)
	}
	return middlewares.MethodNotAlowed(w, r.Method)
}

func (s *ApiServer) handleUserById(w http.ResponseWriter, r *http.Request) error {
	// handler for  user account by id
	if r.Method == "GET" {
		return s.handleUserInfoById(w, r)
	}
	if r.Method == "DELETE" {
		return s.handleDeleteUserById(w, r)
	}
	return middlewares.MethodNotAlowed(w, r.Method)
}
func (s *ApiServer) handleUserInfoById(w http.ResponseWriter, r *http.Request) error {
	// handler for reteriving user info by id
	id, err := middlewares.GetId(r)
	if err != nil {
		return err
	}
	user, err := s.store.UserInfoById(id)
	if err != nil {
		return err
	}
	return middlewares.WriteJSON(w, http.StatusOK, user)
}

func (s *ApiServer) handleDeleteUserById(w http.ResponseWriter, r *http.Request) error {
	// handler for deleting user by id
	id, err := middlewares.GetId(r)
	if err != nil {
		return err
	}
	if err := s.store.UserDelete(id); err != nil {
		return err
	}
	return middlewares.WriteJSON(w, http.StatusOK, middlewares.ApiSuccess{Success: "user deleted"})
}

func (s *ApiServer) handleUsersAll(w http.ResponseWriter, r *http.Request) error {
	// handler for getting all users
	if r.Method == "GET" {
		users, err := s.store.UsersAll()
		if err != nil {
			return err
		}
		return middlewares.WriteJSON(w, http.StatusOK, users)
	}
	return middlewares.MethodNotAlowed(w, r.Method)
}

func (s *ApiServer) handleMe(w http.ResponseWriter, r *http.Request) error {
	// handler for reteriving profile
	if r.Method == "GET" {
		log.Println(r.Header.Get("Authorization"))
		userData, err := utils.ExractTokenMetaData(r)
		if err != nil {
			return err
		}
		user, err := s.store.UserInfoById(userData.ID)
		if err != nil {
			return err
		}
		return middlewares.WriteJSON(w, http.StatusOK, user)
	}
	return middlewares.MethodNotAlowed(w, r.Method)
}

func (s *ApiServer) handleValidToken(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "POST" {
		tokenData, err := utils.ExractTokenMetaData(r)
		if err != nil {
			return err
		}
		username, err := s.store.UserGetUsername(tokenData.ID)
		if err != nil {
			return err
		}
		return middlewares.WriteJSON(w, http.StatusOK, map[string]string{"isUserAuthenticated": "true", "username": username.Username})
	}
	return middlewares.MethodNotAlowed(w, r.Method)
}
