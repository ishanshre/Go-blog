package handlerAndRouters

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ishanshre/Go-blog/api/v1/middlewares"
	"github.com/ishanshre/Go-blog/api/v1/models"
	"github.com/ishanshre/Go-blog/internals/pkg/utils"
)

func (s *ApiServer) handlePostCreate(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "POST" {
		filename, err := utils.UploadPostImage(r)
		if err != nil {
			return err
		}
		req := r.MultipartForm.Value
		content := req["content"][0]
		title := req["title"][0]
		user, err := utils.ExractTokenMetaData(r)
		if err != nil {
			return err
		}
		log.Println(user.ID)
		post := models.NewPostCreate(title, content, filename, user.ID)
		if err := s.store.PostCreate(post); err != nil {
			return err
		}
		return middlewares.WriteJSON(w, http.StatusCreated, post)
	}
	return middlewares.MethodNotAlowed(w, r.Method)
}

func (s *ApiServer) handlePostAll(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		page := new(models.Page)
		if err := json.NewDecoder(r.Body).Decode(&page); err != nil {
			return err
		}
		protocol := utils.CheckHttpProtocol(r)
		domain := fmt.Sprintf("%s://%s", protocol, r.Host)
		posts, err := s.store.PostGetAll(page.Limit, page.Offset, domain)
		if err != nil {
			return err
		}
		return middlewares.WriteJSON(w, http.StatusOK, posts)
	}
	return middlewares.MethodNotAlowed(w, r.Method)
}

func (s *ApiServer) handlePostGetBySlug(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		slug := mux.Vars(r)["slug"]
		protocol := utils.CheckHttpProtocol(r)
		domain := fmt.Sprintf("%s://%s", protocol, r.Host)
		post, err := s.store.PostGetBySlug(slug, domain)
		if err != nil {
			return err
		}
		return middlewares.WriteJSON(w, http.StatusOK, post)
	}
	return middlewares.MethodNotAlowed(w, r.Method)
}
