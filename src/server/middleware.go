package server

import (
	"net/http"

	"gitlab.fit.cvut.cz/isszp/isszp/src/server/session"
)

// MustLogin allows only logged in users. If user is not logged in, MustLogin returns HTTP Error
func MustLogin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !session.IsLoggedIn(r) {
			http.Error(w, "Not logged in!", http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// RedirectToLogin redirects not logged in users to the /login page
func RedirectToLogin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !session.IsLoggedIn(r) && r.RequestURI != "/login" {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}

		next.ServeHTTP(w, r)
	})
}
