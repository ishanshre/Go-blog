package utils

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func UploadPostImage(r *http.Request) (string, error) {
	if err := r.ParseMultipartForm(8 * 1024 * 1024); err != nil {
		return "", fmt.Errorf("image size must be less than 8 MB") //limit upload size
	}
	file, handler, err := r.FormFile("image")
	if err != nil {
		return "", nil
	}
	defer file.Close()
	// create new directroy if does not already exists
	if err := os.MkdirAll("./media/uploads/posts", os.ModePerm); err != nil {
		return "", err
	}
	filename := fmt.Sprintf("%d%s", time.Now().UnixNano(), filepath.Ext(handler.Filename))
	path := fmt.Sprintf("./media/uploads/posts/%v", filename)
	new, err := os.Create(path)
	if err != nil {
		return "", err
	}
	defer new.Close()
	_, err = io.Copy(new, file)
	if err != nil {
		return "", err
	}

	return filename, nil
}
