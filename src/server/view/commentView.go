package view

import (
	"fmt"
	"net/http"

	"gitlab.fit.cvut.cz/isszp/isszp/src/controller"
	"gitlab.fit.cvut.cz/isszp/isszp/src/model"
	"gitlab.fit.cvut.cz/isszp/isszp/src/server/session"
)

func CommentSavePOST(w http.ResponseWriter, r *http.Request) {
	userName := session.GetUserName(r)
	taskID := r.URL.Query().Get("task_id")
	text := r.FormValue("Text")

	err := controller.NewComment(userName, taskID, text)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	t := model.NewTask()
	t.FillByID(taskID)

	if t.GetMaintainer().GetUserName() != session.GetUserName(r) {
		msg := fmt.Sprintf("Added new comment to task: '%s'", t.GetName())
		address := fmt.Sprintf("/task/view/%s", t.GetID())

		controller.SendNotification(t.GetMaintainer().GetID(), msg,
			controller.NotifyWithAddress(address))
	}

	http.Redirect(w, r, "/task/view/"+taskID, http.StatusSeeOther)
}
