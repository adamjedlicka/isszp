package api

import (
	"encoding/json"
	"net/http"

	"gitlab.fit.cvut.cz/isszp/isszp/src/controller"
	"gitlab.fit.cvut.cz/isszp/isszp/src/model"
	"gitlab.fit.cvut.cz/isszp/isszp/src/server/session"
)

// NotifyGET returns JSON-formatted array of notifications for currently logged int user
func NotifyGET(w http.ResponseWriter, r *http.Request) {
	u := model.NewUser()
	u.FillByUserName(session.GetUserName(r))

	notifications := controller.GetNotifications(u.GetID())
	bytes, err := json.MarshalIndent(notifications, "", "\t")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(bytes)
}
