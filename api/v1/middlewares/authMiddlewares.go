package middlewares

import (
	"log"
	"net/http"

	"github.com/ishanshre/Go-blog/api/v1/db"
	"github.com/ishanshre/Go-blog/internals/pkg/utils"
)

func JwtAuthHandler(handlerFunc http.HandlerFunc, s db.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId, err := utils.ExractTokenMetaData(r)
		if err != nil {
			log.Println(err)
			permissionDenied(w)
			return
		}
		account, err := s.UserInfoById(userId.ID)
		if err != nil {
			log.Println(err)
			permissionDenied(w)
			return
		}
		if err := utils.VerifyUser(account.ID, r); err != nil {
			log.Println(err)
			permissionDenied(w)
			return
		}
		handlerFunc(w, r)
	}
}

func JwtAuthPermissionHandler(handlerFunc http.HandlerFunc, s db.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId, err := utils.ExractTokenMetaData(r)
		if err != nil {
			log.Println(err)
			permissionDenied(w)
			return
		}
		account, err := s.UserInfoById(userId.ID)
		if err != nil {
			log.Println(err)
			permissionDenied(w)
			return
		}
		if err := utils.VerifyUser(account.ID, r); err != nil {
			log.Println(err)
			permissionDenied(w)
			return
		}
		paramsId, err := GetId(r)
		if err != nil {
			log.Println(err)
			permissionDenied(w)
			return
		}
		if paramsId != account.ID {
			log.Println(err)
			permissionDenied(w)
			return
		}
		handlerFunc(w, r)
	}
}

func JwtAuthPostOwnerPermissionHandler(handlerFunc http.HandlerFunc, s db.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId, err := utils.ExractTokenMetaData(r)
		if err != nil {
			log.Println(err)
			permissionDenied(w)
			return
		}
		account, err := s.UserInfoById(userId.ID)
		if err != nil {
			log.Println(err)
			permissionDenied(w)
			return
		}
		if err := utils.VerifyUser(account.ID, r); err != nil {
			log.Println(err)
			permissionDenied(w)
			return
		}
		post_id, err := GetId(r)
		if err != nil {
			log.Println(err)
			permissionDenied(w)
			return
		}
		postOwner, err := s.PostGetOwner(post_id)
		if err != nil {
			log.Println(err)
			permissionDenied(w)
			return
		}
		if account.ID != postOwner.User_id {
			permissionDenied(w)
			return
		}
		handlerFunc(w, r)
	}
}

func JwtAuthCommentOwnerPermissionHandler(handlerFunc http.HandlerFunc, s db.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId, err := utils.ExractTokenMetaData(r)
		if err != nil {
			log.Println(err)
			permissionDenied(w)
			return
		}
		account, err := s.UserInfoById(userId.ID)
		if err != nil {
			log.Println(err)
			permissionDenied(w)
			return
		}
		if err := utils.VerifyUser(account.ID, r); err != nil {
			log.Println(err)
			permissionDenied(w)
			return
		}
		comment_id, err := GetCommentId(r)
		if err != nil {
			log.Println(err)
			permissionDenied(w)
			return
		}
		commentOwner, err := s.CommentOwner(comment_id)
		if err != nil {
			log.Println(err)
			permissionDenied(w)
			return
		}
		if account.ID != commentOwner.User_id {
			permissionDenied(w)
			return
		}
		handlerFunc(w, r)
	}
}

func JwtAuthAdminHandler(handlerFunc http.HandlerFunc, s db.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId, err := utils.ExractTokenMetaData(r)
		if err != nil {
			log.Println(err)
			permissionDenied(w)
			return
		}
		account, err := s.UserInfoById(userId.ID)
		if err != nil {
			permissionDenied(w)
			return
		}
		if err := utils.VerifyUser(account.ID, r); err != nil {
			permissionDenied(w)
			return

		}
		if !account.IsAdmin {
			permissionDenied(w)
			return
		}
		handlerFunc(w, r)
	}
}
