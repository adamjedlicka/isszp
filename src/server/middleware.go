package server

import (
	"log"
	"net/http"

	"gitlab.fit.cvut.cz/isszp/isszp/src/model"
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

// IsAdmin if currently logged in user has admin rights. If not returns an HTTP error
func IsAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := model.NewUser()
		userID := session.GetUserUUID(r)

		log.Println(userID)

		err := user.FillByID(userID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if user.GetPermission()&model.IsAdmin == 0 {
			http.Error(w, "You are not admin!", http.StatusBadRequest)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// IsAdmin if currently logged in user can manage projects. If not returns an HTTP error
func CanManageProjects(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := model.NewUser()
		userID := session.GetUserUUID(r)

		log.Println(userID)

		err := user.FillByID(userID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if user.GetPermission()&model.CanManageProjects == 0 {
			http.Error(w, "You cannot manage projects!", http.StatusBadRequest)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// IsAdmin if currently logged in user can manage tasks. If not returns an HTTP error
func CanManageTasks(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := model.NewUser()
		userID := session.GetUserUUID(r)

		log.Println(userID)

		err := user.FillByID(userID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if user.GetPermission()&model.CanManageTasks == 0 {
			http.Error(w, "You cannot manage tasks!", http.StatusBadRequest)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// IsAdmin if currently logged in user can manage users. If not returns an HTTP error
func CanManageUsers(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := model.NewUser()
		userID := session.GetUserUUID(r)

		log.Println(userID)

		err := user.FillByID(userID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if user.GetPermission()&model.CanManageUsers == 0 {
			http.Error(w, "You cannot manage users!", http.StatusBadRequest)
			return
		}

		next.ServeHTTP(w, r)
	})
}
