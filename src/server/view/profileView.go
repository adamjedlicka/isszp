package view

import (
	"net/http"
	"time"

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
	timer := model.NewTimeRecord()

	r.ParseForm()

	taskID := r.FormValue("taskID")
	userID := session.GetUserUUID(r)
	date := time.Now().Local().Format("2006-01-02")
	time := time.Now().Local()

	timer.SetTaskByID(taskID)
	timer.SetUserByID(userID)
	timer.SetDate(date)
	timer.SetStart(time.Format("15:04:05"))

	err := timer.Save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func StopHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	//log.Println(r)

	//log.Printf("%#v\n", r.FormValue("time"))

}
