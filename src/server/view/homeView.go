package view

import (
	"net/http"

	"gitlab.fit.cvut.cz/isszp/isszp/src/model"
	"gitlab.fit.cvut.cz/isszp/isszp/src/server/session"
)

func HomeGET(w http.ResponseWriter, r *http.Request) {
	view := NewView(r, "Home")
	view.AppendTemplates("home", "component/task-list", "component/timer")

	userID := session.GetUserUUID(r)
	user := model.NewUser()
	user.FillByID(userID)

	view.Vars["Worker"] = map[string]interface{}{
		"Tasks":          model.QueryTasks("WorkerID = ? AND State != 'success' AND State != 'fail'", user.GetID()),
		"HideWorker":     true,
		"HideMaintainer": true,
	}

	view.Vars["Maintainer"] = map[string]interface{}{
		"Tasks":          model.QueryTasks("MaintainerID = ? AND State != 'success' AND State != 'fail'", user.GetID()),
		"HideMaintainer": true,
	}

	view.Vars["TimerTasks"] = model.QueryTasks("WorkerID = ? AND EndDate IS NULL", user.GetID())

	view.Render(w)
}
