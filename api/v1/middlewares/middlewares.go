package middlewares

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type ApiFunc func(http.ResponseWriter, *http.Request) error // signature of our handler

type ApiError struct {
	// error signature
	Error string `json:"error"`
}

type ApiSuccess struct {
	// success signature
	Success string `json:"success"`
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	// It is a reponse to the request
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

func MethodNotAlowed(w http.ResponseWriter, method string) error {
	// middleware for unallowed methods
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	w.WriteHeader(http.StatusMethodNotAllowed)
	return json.NewEncoder(w).Encode(ApiError{Error: fmt.Sprintf("%s method not allowed", method)})
}

func MakeHttpHandler(f ApiFunc) http.HandlerFunc {
	// return http.HandlerFunc
	return func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w, r)
		log.Println(w)

		if err := f(w, r); err != nil {
			WriteJSON(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}

func GetId(r *http.Request) (int, error) {
	// returns id in int
	idstr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idstr)
	if err != nil {
		return id, fmt.Errorf("error in parsing the url paramater")
	}
	return id, nil
}

func GetCommentId(r *http.Request) (int, error) {
	// returns commet id in int
	idstr := mux.Vars(r)["comment_id"]
	id, err := strconv.Atoi(idstr)
	if err != nil {
		return id, fmt.Errorf("error in parsing the url paramater")
	}
	return id, nil
}

func GetSlug(r *http.Request) string {
	// middlewares that returns slug
	slug := mux.Vars(r)["slug"]
	return slug
}

func enableCors(w *http.ResponseWriter, r *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, HEAD, OPTIONS")
	(*w).Header().Set("Access-Control-Allow-Credentials", "true")
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With, XMLHttpRequest")
	if r.Method == "OPTIONS" {
		(*w).WriteHeader(http.StatusOK)
		return
	}
}
