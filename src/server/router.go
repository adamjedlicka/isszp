package server

import (
	"net/http"

	"github.com/gorilla/mux"

	"gitlab.fit.cvut.cz/isszp/isszp/src/server/api"
	"gitlab.fit.cvut.cz/isszp/isszp/src/server/view"
)

// NewRouter creates new router for HTTP server
// router decides when users visits certain address which function will handle that request
func NewRouter() http.Handler {
	r := mux.NewRouter()

	r.Handle("/", use(view.HomeGET, RedirectToLogin)).Methods("GET")

	r.Handle("/firms", use(view.FirmsGET, MustLogin)).Methods("GET")
	r.Handle("/firm/new", use(view.FirmNewGET, MustLogin)).Methods("GET")
	r.Handle("/firm/view/{ID}", use(view.FirmViewGET, MustLogin)).Methods("GET")
	r.Handle("/firm/edit/{ID}", use(view.FirmEditGET, MustLogin)).Methods("GET")
	r.Handle("/firm/delete/{ID}", use(view.FirmDelGET, MustLogin)).Methods("GET")
	r.Handle("/firm/save", use(view.FirmSavePOST, MustLogin)).Methods("POST")

	r.Handle("/projects", use(view.ProjectsGET, MustLogin)).Methods("GET")
	r.Handle("/project/new", use(view.ProjectNewGET, CanManageProjects, MustLogin)).Methods("GET")
	r.Handle("/project/view/{ID}", use(view.ProjectViewGET, MustLogin)).Methods("GET")
	r.Handle("/project/edit/{ID}", use(view.ProjectEditGET, CanManageProjects, MustLogin)).Methods("GET")
	r.Handle("/project/delete/{ID}", use(view.ProjectDeleteGET, CanManageProjects, MustLogin)).Methods("GET")
	r.Handle("/project/save", use(view.ProjectSavePOST, CanManageProjects, MustLogin)).Methods("POST")

	r.Handle("/tasks", use(view.TasksGET, MustLogin)).Methods("GET")
	r.Handle("/task/new", use(view.TaskNewGET, CanManageTasks, MustLogin)).Methods("GET")
	r.Handle("/task/view/{ID}", use(view.TaskViewGET, MustLogin)).Methods("GET")
	r.Handle("/task/edit/{ID}", use(view.TaskEditGET, CanManageTasks, MustLogin)).Methods("GET")
	r.Handle("/task/delete/{ID}", use(view.TaskDeleteGET, CanManageTasks, MustLogin)).Methods("GET")
	r.Handle("/task/save", use(view.TaskSavePOST, CanManageTasks, MustLogin)).Methods("POST")

	r.Handle("/comment/save", use(view.CommentSavePOST, MustLogin)).Methods("POST")

	r.Handle("/users", use(view.UsersGET, MustLogin)).Methods("GET")
	r.Handle("/user/new", use(view.UserNewGET, CanManageUsers, MustLogin)).Methods("GET")
	r.Handle("/user/view/{ID}", use(view.UserViewGET, MustLogin)).Methods("GET")
	r.Handle("/user/edit/{ID}", use(view.UserEditGET, MustLogin)).Methods("GET")
	r.Handle("/user/delete/{ID}", use(view.UserDeleteGET, CanManageUsers, MustLogin)).Methods("GET")
	r.Handle("/user/save", use(view.UserSavePOST, CanManageUsers, MustLogin)).Methods("POST")
	r.Handle("/user/password", use(view.UserPasswordPOST, MustLogin)).Methods("POST")

	r.Handle("/timerecords", use(view.TimerecordsGET, MustLogin)).Methods("GET")
	r.Handle("/timerecord/view/{ID}", use(view.TimerecordViewGET, MustLogin)).Methods("GET")
	r.Handle("/timerecord/edit/{ID}", use(view.TimerecordEditGET, MustLogin)).Methods("GET")
	r.Handle("/timerecord/delete/{ID}", use(view.TimerecordDeleteGET, MustLogin)).Methods("GET")
	r.Handle("/timerecord/save", use(view.TimerecordSavePOST, MustLogin)).Methods("POST")

	r.Handle("/login", use(view.LoginGET)).Methods("GET")
	r.Handle("/login", use(view.LoginPOST)).Methods("POST")
	r.Handle("/logout", use(view.LogoutGET)).Methods("GET")

	r.Handle("/api/timer", use(api.TimerGET, MustLogin)).Methods("GET")
	r.Handle("/api/timer/start", use(api.TimerStartPOST, MustLogin)).Methods("POST")
	r.Handle("/api/timer/stop", use(api.TimerStopPOST, MustLogin)).Methods("POST")
	r.Handle("/api/notify", use(api.NotifyGET, MustLogin)).Methods("GET")

	// serve files from ./static/ directory without any special routing
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	return r
}

// use is function that helps with chaining middleware
// first argument is final handler function and the you can append various number of middleware functions
func use(handler func(http.ResponseWriter, *http.Request), middleware ...func(http.Handler) http.Handler) http.Handler {
	var h http.Handler
	h = http.HandlerFunc(handler)

	for _, fn := range middleware {
		h = fn(h)
	}

	return h
}
