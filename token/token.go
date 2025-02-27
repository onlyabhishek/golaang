package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Define authenticationMiddleware struct
type authenticationMiddleware struct {
	tokenUsers map[string]string
}

// Populate the tokenUsers map with sample tokens
func (amw *authenticationMiddleware) Populate() {
	amw.tokenUsers["00000000"] = "user0"
	amw.tokenUsers["aaaaaaaa"] = "userA"
	amw.tokenUsers["05f717e5"] = "randomUser"
	amw.tokenUsers["deadbeef"] = "user0"
}

// Middleware function to authenticate users based on token
func (amw *authenticationMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("X-Session-Token")

		if user, found := amw.tokenUsers[token]; found {
			// If token is found, user is authenticated
			log.Printf("Authenticated user: %s\n", user)
			next.ServeHTTP(w, r) // Call the next handler
		} else {
			// If token is invalid, return 403 Forbidden
			http.Error(w, "Forbidden", http.StatusForbidden)
		}
	})
}

// Example handler function for authenticated requests
func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome! You are authenticated."))
}

func main() {
	// Create a new Gorilla Mux router
	r := mux.NewRouter()

	// Register a handler for the root path "/"
	r.HandleFunc("/", handler)

	// Create an instance of authenticationMiddleware
	amw := authenticationMiddleware{tokenUsers: make(map[string]string)}
	amw.Populate() // Populate tokens

	// Apply middleware to the router
	r.Use(amw.Middleware)

	// Start the server
	log.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
