package session

import (
	"net/http"

	"github.com/gorilla/sessions"
)

const Login = "LOGIN"

var Store = sessions.NewCookieStore([]byte("something-very-secret"))

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
