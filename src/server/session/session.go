// Package session contains implementation of session handling
// and some useful functions that will help you to correctly a safely extract certain data from sessions
package session

import (
	"net/http"

	"github.com/gorilla/sessions"
)

// Login is cookie name used to store login token
const Login = "LOGIN"

// Store is a cookie store that is used to obtain user cookies based on its cookie name
var Store = sessions.NewCookieStore([]byte("something-very-secret"))

// IsLoggedIn returns true if user that sent request r is lgged int.
// If not or no session/cookie exists IsLoggedIn returns false
func IsLoggedIn(r *http.Request) bool {
	s, err := Store.Get(r, Login)
	if err != nil {
		return false
	}

	isLoggedIn, ok := s.Values["IsLoggedIn"].(bool)
	if !ok {
		return false
	}

	return isLoggedIn
}

// GetUserName returns user name of currently logged user.
// If no user is logged in GetUserName returns empty string ""
func GetUserName(r *http.Request) string {
	s, err := Store.Get(r, Login)
	if err != nil {
		return ""
	}

	userName, ok := s.Values["UserName"].(string)
	if !ok {
		return ""
	}

	return userName
}
