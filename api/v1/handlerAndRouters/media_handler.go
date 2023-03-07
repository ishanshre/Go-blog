package handlerAndRouters

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func (s *ApiServer) handleMediaImage(w http.ResponseWriter, r *http.Request) error {
	filename := mux.Vars(r)["filename"]
	media_url := "./media/uploads/posts/"
	path := fmt.Sprintf("%s%s", media_url, filename)
	log.Println(path)
	http.ServeFile(w, r, path)
	return nil
}
