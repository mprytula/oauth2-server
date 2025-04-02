package handlers

import (
	"log/slog"
	"net/http"
	"os"
)
func LoadPage(htmlPath string, next http.HandlerFunc) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {	
		var file, err = os.ReadFile(htmlPath)
		if err != nil {
			slog.Warn("Error reading file", "error", err)
		} else {
			defer w.Write(file)
		}

		if (next != nil) {
			next(w, r)
		}
	}
}