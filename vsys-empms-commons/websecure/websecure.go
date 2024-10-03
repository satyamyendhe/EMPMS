package websecure

import (
	"log"
	"net/http"
	"strings"

	u "vsys.empms.commons/utils"
)

// CommonMiddleware is an HTTP middleware that handles authentication and static file serving
func CommonMiddleware(next http.Handler) http.Handler {
	// Define the directory for static files
	staticDir := "./static"

	// Create a file server handler for static files
	staticFileServer := http.StripPrefix("/static/", http.FileServer(http.Dir(staticDir)))

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path

		// Serve static files
		if strings.HasPrefix(path, "/static/") {
			staticFileServer.ServeHTTP(w, r)
			return
		}

		// Allow specific public URLs like login and signup
		allowedURLs := []string{"/", "/login", "/signup", "/get-emps", "/get-emp", "/get-logs"}
		for _, url := range allowedURLs {
			if path == url {
				next.ServeHTTP(w, r)
				return
			}
		}

		// Check for token in cookies
		token, err := r.Cookie("auth_token")
		if err == nil && token != nil {
			if u.ValidateJwtToken(token.Value) {
				next.ServeHTTP(w, r)
				return
			}
		}

		// Redirect to login if not valid
		log.Printf("Unauthorized request: %s", r.URL.Path)
		http.Redirect(w, r, "/login", http.StatusFound)
	})
}
