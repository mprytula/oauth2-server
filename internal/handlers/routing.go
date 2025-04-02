package handlers

import (
	"net/http"
	"os"
	"path/filepath"
)

func AttachRouting(mutex *http.ServeMux) {
	if mutex == nil {
		return
	}

	basePath, _ := os.Getwd()
	mainPage := filepath.Join(basePath, "web", "assets", "index.html")

	mutex.HandleFunc("/", LoadPage(mainPage, nil))
	mutex.HandleFunc("/callback", Auth(LoadPage(mainPage, nil)))
}
