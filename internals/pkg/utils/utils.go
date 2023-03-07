package utils

import "net/http"

func CheckHttpProtocol(r *http.Request) string {
	// returns http or https based on connection type
	if r.TLS != nil && r.TLS.HandshakeComplete {
		return "https"
	} else {
		return "http"
	}
}
