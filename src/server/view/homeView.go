package view

import (
	"net/http"

	"gitlab.fit.cvut.cz/isszp/isszp/src/model"
	"gitlab.fit.cvut.cz/isszp/isszp/src/server/session"
)

func HomeGET(w http.ResponseWriter, r *http.Request) {
	view := NewView(r, "Home")
	view.AppendTemplates("home", "component/task-list", "component/stopwatch")

	userName := session.GetUserName(r)
	u := model.NewUser()
	u.FillByUserName(userName)

	view.Vars["Worker"] = map[string]interface{}{
		"Tasks":          model.QueryTasks("WorkerID = ? AND State != 'success' AND State != 'fail'", u.GetID()),
		"HideWorker":     true,
		"HideMaintainer": true,
	}

	view.Vars["Maintainer"] = map[string]interface{}{
		"Tasks":          model.QueryTasks("MaintainerID = ? AND State != 'success' AND State != 'fail'", u.GetID()),
		"HideMaintainer": true,
	}

	currentUserID := session.GetUserUUID(r)

	view.Vars["Tasks"] = model.QueryTasks("WorkerID = ?", currentUserID)
	currentUserName := model.QueryUsers("id = ?", currentUserID)

	taskRecord := model.QueryTimeRecords("UserID = ? AND End IS NULL", currentUserID)

	if len(taskRecord) > 0 { // Max one timeRecord with userID = x and End = 00:00:00, but there can be non
		view.Vars["StartTime"] = taskRecord[0]
	}

	if len(currentUserName) > 0 { // Max one user with this UUID, but there can be non
		view.Vars["CUser"] = currentUserName[0]
	}

	view.Render(w)
}
