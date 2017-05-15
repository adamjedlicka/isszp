package view

import (
	"net/http"

	"github.com/gorilla/mux"

	"gitlab.fit.cvut.cz/isszp/isszp/src/controller"
	"gitlab.fit.cvut.cz/isszp/isszp/src/model"
	"gitlab.fit.cvut.cz/isszp/isszp/src/server/session"
)

func UsersGET(w http.ResponseWriter, r *http.Request) {
	view := NewView(r, "Users")
	view.AppendTemplates("users/users", "component/user-list")

	view.Vars["Users"] = model.QueryUsers()

	view.Render(w)
}

func UserNewGET(w http.ResponseWriter, r *http.Request) {
	view := NewView(r, "User :: new")
	view.AppendTemplates("users/user-view")

	view.Vars["Action"] = "new"

	view.Render(w)
}

func UserViewGET(w http.ResponseWriter, r *http.Request) {
	view := NewView(r, "User :: detail")
	view.AppendTemplates("users/user-view")

	id := mux.Vars(r)["ID"]
	user := model.NewUser()
	err := user.FillByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	view.Name += user.GetUserName()

	view.Vars["User"] = user
	view.Vars["Action"] = "view"

	view.Vars["IsAdmin"] = user.GetPermission()&model.IsAdmin == model.IsAdmin
	view.Vars["CanManageProjects"] = user.GetPermission()&model.CanManageProjects == model.CanManageProjects
	view.Vars["CanManageTasks"] = user.GetPermission()&model.CanManageTasks == model.CanManageTasks
	view.Vars["CanManageUsers"] = user.GetPermission()&model.CanManageUsers == model.CanManageUsers

	view.Vars["readonly"] = "readonly"
	view.Vars["disabled"] = "disabled"

	view.Render(w)
}

func UserEditGET(w http.ResponseWriter, r *http.Request) {
	view := NewView(r, "User :: edit")
	view.AppendTemplates("users/user-view")

	id := mux.Vars(r)["ID"]
	user := model.NewUser()
	err := user.FillByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	view.Vars["User"] = user
	view.Vars["Action"] = "edit"

	view.Vars["IsAdmin"] = user.GetPermission()&model.IsAdmin == model.IsAdmin
	view.Vars["CanManageProjects"] = user.GetPermission()&model.CanManageProjects == model.CanManageProjects
	view.Vars["CanManageTasks"] = user.GetPermission()&model.CanManageTasks == model.CanManageTasks
	view.Vars["CanManageUsers"] = user.GetPermission()&model.CanManageUsers == model.CanManageUsers

	viewer := model.NewUser()
	viewer.FillByID(session.GetUserUUID(r))
	if viewer.GetPermission() != model.IsAdmin {
		view.Vars["disabled"] = "disabled"
	}

	view.Render(w)
}

func UserDeleteGET(w http.ResponseWriter, r *http.Request) {
	user := model.NewUser()
	err := user.FillByID(mux.Vars(r)["ID"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = user.Delete()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set URL to the URL user came from
	url := r.Header.Get("Referer")
	if url == "" {
		url = "/users"
	}

	http.Redirect(w, r, url, http.StatusSeeOther)
}

func UserSavePOST(w http.ResponseWriter, r *http.Request) {
	user := model.NewUser()

	id := r.FormValue("ID")
	if id != "" {
		err := user.FillByID(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	user.SetUserName(r.FormValue("Username"))
	user.SetFirstName(r.FormValue("FirstName"))
	user.SetLastName(r.FormValue("LastName"))
	controller.SetUserHashedPassword(user, r.FormValue("Password"))

	viewer := model.NewUser()
	viewer.FillByID(session.GetUserUUID(r))
	if viewer.GetPermission() == model.IsAdmin {
		user.SetPermission(0)
		if r.FormValue("IsAdmin") == "on" {
			user.AddPermission(model.IsAdmin)
		}
		if r.FormValue("CanManageProjects") == "on" {
			user.AddPermission(model.CanManageProjects)
		}
		if r.FormValue("CanManageTasks") == "on" {
			user.AddPermission(model.CanManageTasks)
		}
		if r.FormValue("CanManageUsers") == "on" {
			user.AddPermission(model.CanManageUsers)
		}
	}

	err := user.Save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/user/view/"+user.GetID(), http.StatusSeeOther)
}
