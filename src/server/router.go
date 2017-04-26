package server

import (
	"net/http"

	"gitlab.fit.cvut.cz/isszp/isszp/src/server/view"

	"github.com/gorilla/mux"
)

func NewRouter() http.Handler {
	r := mux.NewRouter()

	r.Handle("/", http.HandlerFunc(view.HomeGET)).Methods("GET")

	r.Handle("/firms", MustLogin(http.HandlerFunc(view.FirmsGET))).Methods("GET")
	r.Handle("/firm/new", MustLogin(http.HandlerFunc(view.FirmNewGET))).Methods("GET")
	r.Handle("/firm/view/{ID}", MustLogin(http.HandlerFunc(view.FirmViewGET))).Methods("GET")
	r.Handle("/firm/edit/{ID}", MustLogin(http.HandlerFunc(view.FirmEditGET))).Methods("GET")
	r.Handle("/firm/delete/{ID}", MustLogin(http.HandlerFunc(view.FirmDelGET))).Methods("GET")
	r.Handle("/firm/save", MustLogin(http.HandlerFunc(view.FirmSavePOST))).Methods("POST")

	r.Handle("/tasks", MustLogin(http.HandlerFunc(view.TasksGET))).Methods("GET")
	r.Handle("/task/new", MustLogin(http.HandlerFunc(view.TaskNewGET))).Methods("GET")
	r.Handle("/task/view/{ID}", MustLogin(http.HandlerFunc(view.TaskViewGET))).Methods("GET")
	r.Handle("/task/edit/{ID}", MustLogin(http.HandlerFunc(view.TaskEditGET))).Methods("GET")
	r.Handle("/task/delete/{ID}", MustLogin(http.HandlerFunc(view.TaskDeleteGET))).Methods("GET")
	r.Handle("/task/save", MustLogin(http.HandlerFunc(view.TaskSavePOST))).Methods("POST")

	r.Handle("/comment/save", MustLogin(http.HandlerFunc(view.CommentSavePOST))).Methods("POST")

	r.Handle("/login", http.HandlerFunc(view.LoginGET)).Methods("GET")
	r.Handle("/login", http.HandlerFunc(view.LoginPOST)).Methods("POST")
	r.Handle("/logout", http.HandlerFunc(view.LogoutGET)).Methods("GET")

	// serve files from ./static/ directory without any special routing
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	return r
}
