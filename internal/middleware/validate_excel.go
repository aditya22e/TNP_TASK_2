package middleware

import (
	"net/http"
)

// ValidateExcel middleware checks that an Excel file is included in the request
func ValidateExcel(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
		if err := r.ParseMultipartForm(10 << 20); err != nil {
			http.Error(w, "failed to parse multipart form", http.StatusBadRequest)
			return
		}
		if _, _, err := r.FormFile("file"); err != nil {
			http.Error(w, "no file uploaded", http.StatusBadRequest)
			return
		}
		next.ServeHTTP(w, r)
	})
}
