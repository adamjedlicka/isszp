package view

import (
	"net/http"

	"gitlab.fit.cvut.cz/isszp/isszp/src/model"
	"gitlab.fit.cvut.cz/isszp/isszp/src/server/session"
)

func ProfileGET(w http.ResponseWriter, r *http.Request) {
	view := NewView(r, "Profile")
	view.AppendTemplates("profile/profile")

	view.Vars["Tasks"] = model.QueryTasks("WorkerID = ?", session.GetUserUUID) // Nefunguje

	view.Render(w)
}
