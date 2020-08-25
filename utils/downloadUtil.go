package utils

import (
	"mime"
	"net/http"
	"path/filepath"
)

func ParseFileInfoFrom(response *http.Response) string {
	contentDisposition := response.Header.Get("Content-Disposition")
	if contentDisposition != "" {
		_, params, err := mime.ParseMediaType(contentDisposition)
		if err != nil {
			panic(err)
		}
		return params["filename"]
	}
	filename := filepath.Base(response.Request.URL.Path)
	return filename
}
