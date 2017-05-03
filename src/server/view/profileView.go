package view

import (
	"log"
	"net/http"

	"gitlab.fit.cvut.cz/isszp/isszp/src/model"
	"gitlab.fit.cvut.cz/isszp/isszp/src/server/session"
)

func ProfileGET(w http.ResponseWriter, r *http.Request) {
	view := NewView(r, "Profile")
	view.AppendTemplates("profile/profile")

	view.Vars["Tasks"] = model.QueryTasks("WorkerID = ?", session.GetUserUUID(r)) // Nefunguje

	view.Render(w)
}

func StartHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	log.Println(r)

	//log.Printf("%#v\n", r.FormValue("startDate"))

}

func StopHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	log.Println(r)

	//log.Printf("%#v\n", r.FormValue("time"))

}
