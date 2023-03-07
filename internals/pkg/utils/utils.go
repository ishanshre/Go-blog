package utils

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gosimple/slug"
)

func CheckHttpProtocol(r *http.Request) string {
	// returns http or https based on connection type
	if r.TLS != nil && r.TLS.HandshakeComplete {
		return "https"
	} else {
		return "http"
	}
}

func CreateSlug(text string) string {
	// create a slug and retuns it as string
	return fmt.Sprintf("%s-%s", slug.Make(text), time.Now().Format(time.DateTime))
}
