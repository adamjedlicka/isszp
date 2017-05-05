package view

import (
	"net/http"

	"gitlab.fit.cvut.cz/isszp/isszp/src/model"
	"gitlab.fit.cvut.cz/isszp/isszp/src/server/session"
)

func HomeGET(w http.ResponseWriter, r *http.Request) {
	view := NewView(r, "Home")
	view.AppendTemplates("home", "component/task-list")

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

	view.Render(w)
}
