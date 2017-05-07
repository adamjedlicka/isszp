package view

import (
	"net/http"

	"github.com/gorilla/mux"

	"gitlab.fit.cvut.cz/isszp/isszp/src/model"
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

	view.Vars["readonly"] = "readonly"

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
	user.SetPassword(r.FormValue("Password"))

	err := user.Save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/user/view/"+user.GetID(), http.StatusSeeOther)
}
