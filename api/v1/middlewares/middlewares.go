package middlewares

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/ishanshre/Go-blog/api/v1/models"
)

type ApiFunc func(http.ResponseWriter, *http.Request) error // signature of our handler

type ApiError struct {
	Error string `json:"error"`
}

type ApiSuccess struct {
	Success string `json:"success"`
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

func MakeHttpHandler(f ApiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJSON(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}

func ScanUserPass(rows *sql.Rows) (*models.UserPass, error) {
	user := new(models.UserPass)
	err := rows.Scan(
		&user.ID,
		&user.Password,
	)
	return user, err
}
