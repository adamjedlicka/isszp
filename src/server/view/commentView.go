package view

import (
	"net/http"

	"gitlab.fit.cvut.cz/isszp/isszp/src/controller"
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

	http.Redirect(w, r, "/task/view/"+taskID, http.StatusSeeOther)
}
