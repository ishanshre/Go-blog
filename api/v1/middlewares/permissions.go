package middlewares

import "net/http"

func permissionDenied(w http.ResponseWriter) {
	// permission message response
	WriteJSON(w, http.StatusForbidden, ApiError{Error: "permission denied"})
}
