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

	currentUserID := session.GetUserUUID(r)

	view.Vars["Tasks"] = model.QueryTasks("WorkerID = ?", currentUserID)

	taskRecord := model.QueryTimeRecords("UserID = ? AND End = '00:00:00'", currentUserID)

	if len(taskRecord) > 0 {
		view.Vars["StartTime"] = taskRecord[0]
	}

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
	timer.SetDescription(r.FormValue("startDate")) // Start time in miliseconds

	err := timer.Save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func StopHandler(w http.ResponseWriter, r *http.Request) {
	/*timer := model.NewTimeRecord()

	r.ParseForm()

	timeRecord := model.QueryTimeRecords("UserID = ?", session.GetUserUUID(r))

	err := timer.FillByID()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	userID := session.GetUserUUID(r)
	time := time.Now().Local()

	timer.SetStop(time.Format("15:04:05"))

	err = timer.Save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}*/
}
