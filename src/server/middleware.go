package server

import (
	"net/http"

	"gitlab.fit.cvut.cz/isszp/isszp/src/server/session"
)

func MustLogin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if session.IsLoggedIn(r) {
			next.ServeHTTP(w, r)
		} else {
			http.Error(w, "Not logged in!", http.StatusForbidden)
		}
	})
}

func RedirectNotLoggedInt(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !session.IsLoggedIn(r) && r.RequestURI != "/login" {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		next.ServeHTTP(w, r)
	})
}
