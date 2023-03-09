package middlewares

import (
	"encoding/json"
	"fmt"
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
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

func MethodNotAlowed(w http.ResponseWriter, method string) error {
	// middleware for unallowed methods
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusMethodNotAllowed)
	return json.NewEncoder(w).Encode(ApiError{Error: fmt.Sprintf("%s method not allowed", method)})
}

func MakeHttpHandler(f ApiFunc) http.HandlerFunc {
	// return http.HandlerFunc
	return func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
		}
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

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Add("Access-Control-Allow-Origin", "*")
	(*w).Header().Add("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	(*w).Header().Add("Access-Control-Allow-Headers", "Accept, Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}
