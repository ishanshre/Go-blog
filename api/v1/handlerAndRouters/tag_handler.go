package handlerAndRouters

import (
	"encoding/json"
	"net/http"

	"github.com/ishanshre/Go-blog/api/v1/middlewares"
	"github.com/ishanshre/Go-blog/api/v1/models"
)

func (s *ApiServer) handleTag(w http.ResponseWriter, r *http.Request) error {
	//handler for reteriving all tags and creating tags
	if r.Method == "GET" {
		return s.handleGetAllTags(w, r)
	}
	if r.Method == "POST" {
		return s.handleCreateNewTag(w, r)
	}
	return middlewares.MethodNotAlowed(w, r.Method)
}

func (s *ApiServer) handleCreateNewTag(w http.ResponseWriter, r *http.Request) error {
	// handler for new tag by admin
	req := new(models.CreateTagRequest)
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return err
	}
	tag := models.CreateNewTag(req.Name)
	if err := s.store.TagCreate(tag); err != nil {
		return err
	}
	return middlewares.WriteJSON(w, http.StatusCreated, middlewares.ApiSuccess{Success: "new tag created"})
}

func (s *ApiServer) handleGetAllTags(w http.ResponseWriter, r *http.Request) error {
	// handler for reteriving all tags
	if r.Method == "GET" {
		req := new(models.Page)
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			return err
		}
		tags, err := s.store.TagAll(req.Limit, req.Offset)
		if err != nil {
			return err
		}
		return middlewares.WriteJSON(w, http.StatusOK, tags)
	}
	return middlewares.MethodNotAlowed(w, r.Method)
}

func (s *ApiServer) handleTagsById(w http.ResponseWriter, r *http.Request) error {
	// handler for fetching tag by id
	if r.Method == "GET" {
		return s.handleGetTagsById(w, r)
	}
	if r.Method == "PUT" {
		return s.handleUpdateTags(w, r)
	}
	if r.Method == "DELETE" {
		return s.handleDeleteTags(w, r)
	}
	return middlewares.MethodNotAlowed(w, r.Method)
}

func (s *ApiServer) handleDeleteTags(w http.ResponseWriter, r *http.Request) error {
	// handler for deleting tag
	id, err := middlewares.GetId(r)
	if err != nil {
		return err
	}
	if err := s.store.TagDelete(id); err != nil {
		return err
	}
	return middlewares.WriteJSON(w, http.StatusOK, middlewares.ApiSuccess{Success: "tags deleted"})
}

func (s *ApiServer) handleUpdateTags(w http.ResponseWriter, r *http.Request) error {
	// handler for updating tag
	id, err := middlewares.GetId(r)
	if err != nil {
		return err
	}
	req := new(models.CreateTagRequest)
	if json.NewDecoder(r.Body).Decode(&req); err != nil {
		return err
	}
	if err := s.store.TagUpdate(id, req); err != nil {
		return err
	}
	return middlewares.WriteJSON(w, http.StatusOK, middlewares.ApiSuccess{Success: "tag updated"})
}

func (s *ApiServer) handleGetTagsById(w http.ResponseWriter, r *http.Request) error {
	// handler for reteriving tag by id
	id, err := middlewares.GetId(r)
	if err != nil {
		return err
	}
	tag, err := s.store.TagByID(id)
	if err != nil {
		return err
	}
	return middlewares.WriteJSON(w, http.StatusOK, tag)
}
