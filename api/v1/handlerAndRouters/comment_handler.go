package handlerAndRouters

import (
	"encoding/json"
	"net/http"

	"github.com/ishanshre/Go-blog/api/v1/middlewares"
	"github.com/ishanshre/Go-blog/api/v1/models"
	"github.com/ishanshre/Go-blog/internals/pkg/utils"
)

func (s *ApiServer) handleCommentCreate(w http.ResponseWriter, r *http.Request) error {
	// create new comment for a specific post
	if r.Method == "POST" {
		req := new(models.CommentReq)
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			return err
		}
		post_id, err := middlewares.GetId(r)
		if err != nil {
			return err
		}
		tokenData, err := utils.ExractTokenMetaData(r)
		if err != nil {
			return err
		}
		comment := models.NewCommentCreate(req.Content, req.Rating, post_id, tokenData.ID)
		if err := s.store.CommentCreate(comment); err != nil {
			return err
		}
		return middlewares.WriteJSON(w, http.StatusCreated, middlewares.ApiSuccess{Success: "comment created"})

	}
	return middlewares.MethodNotAlowed(w, r.Method)
}

func (s *ApiServer) handleCommentByPost(w http.ResponseWriter, r *http.Request) error {
	// handler for comment by id
	if r.Method == "GET" {
		post_id, err := middlewares.GetId(r)
		if err != nil {
			return err
		}
		comments, err := s.store.CommentAllByPost(post_id)
		if err != nil {
			return err
		}
		return middlewares.WriteJSON(w, http.StatusOK, comments)
	}
	return middlewares.MethodNotAlowed(w, r.Method)
}

func (s *ApiServer) handleCommentUpdateDelete(w http.ResponseWriter, r *http.Request) error {
	// handler for updating and deleting comment by owner
	if r.Method == "PUT" {
		return s.handleCommentUpdate(w, r)
	}
	if r.Method == "DELETE" {
		return s.handleCommentDelete(w, r)
	}
	return middlewares.MethodNotAlowed(w, r.Method)
}

func (s *ApiServer) handleCommentUpdate(w http.ResponseWriter, r *http.Request) error {
	// handler for updating comment by comment owner
	post_id, err := middlewares.GetId(r)
	if err != nil {
		return err
	}
	comment_id, err := middlewares.GetCommentId(r)
	if err != nil {
		return err
	}
	req := new(models.CommentReq)
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return err
	}
	comment := models.NewCommentUpdate(req.Content, req.Rating)
	if err := s.store.CommentUpdate(post_id, comment_id, comment); err != nil {
		return err
	}
	return middlewares.WriteJSON(w, http.StatusOK, middlewares.ApiSuccess{Success: "comment updated"})
}

func (s *ApiServer) handleCommentDelete(w http.ResponseWriter, r *http.Request) error {
	// handler comment delete by comment owner
	post_id, err := middlewares.GetId(r)
	if err != nil {
		return err
	}
	comment_id, err := middlewares.GetCommentId(r)
	if err != nil {
		return err
	}
	if err := s.store.CommentDelete(post_id, comment_id); err != nil {
		return err
	}
	return middlewares.WriteJSON(w, http.StatusOK, middlewares.ApiSuccess{Success: "comment deleted"})
}
