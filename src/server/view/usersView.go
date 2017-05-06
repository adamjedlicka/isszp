package view

import (
	"net/http"

	"gitlab.fit.cvut.cz/isszp/isszp/src/model"
)

func UsersGET(w http.ResponseWriter, r *http.Request) {
	view := NewView(r, "Users")
	view.AppendTemplates("users/users", "component/user-list")

	view.Vars["Users"] = model.QueryUsers()

	view.Render(w)
}
