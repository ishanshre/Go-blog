package utils

import "net/http"

func CheckHttpProtocol(r *http.Request) string {
	if r.TLS != nil && r.TLS.HandshakeComplete {
		return "https"
	} else {
		return "http"
	}
}
