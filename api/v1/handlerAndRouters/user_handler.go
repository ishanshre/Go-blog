package handlerAndRouters

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ishanshre/Go-blog/api/v1/middlewares"
	"github.com/ishanshre/Go-blog/api/v1/models"
	"github.com/ishanshre/Go-blog/internals/pkg/utils"
)

func (s *ApiServer) handleUserSignUp(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "POST" {
		req := new(models.RegisterUserRequest)
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			return err
		}
		encPass, err := utils.HashPassword(req.Password)
		if err != nil {
			return err
		}
		user := models.RegisterNewUser(req.Username, req.Email, encPass)
		if err := s.store.UserSignUp(user); err != nil {
			return err
		}
		return middlewares.WriteJSON(w, http.StatusCreated, middlewares.ApiSuccess{Success: "user created"})
	}
	return fmt.Errorf("%s method not allowed", r.Method)
}
