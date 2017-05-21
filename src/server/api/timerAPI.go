// Package api is collection of REST methods that communicate with json instead of HTML
package api

import (
	"encoding/json"
	"net/http"
	"time"

	"gitlab.fit.cvut.cz/isszp/isszp/src/common"
	"gitlab.fit.cvut.cz/isszp/isszp/src/model"
	"gitlab.fit.cvut.cz/isszp/isszp/src/server/session"
)

// TimerGET returns a JSON-formatted message with info about currently running timer for logged in user
func TimerGET(w http.ResponseWriter, r *http.Request) {
	tr := model.QueryTimeRecords("UserID = ? AND End IS NULL", session.GetUserUUID(r))
	bytes, err := json.MarshalIndent(tr, "", "\t")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = w.Write(bytes)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// TimerStartPOST starts a timer for currently logged in user and creates new TimeRecord with nil End time
func TimerStartPOST(w http.ResponseWriter, r *http.Request) {
	task := model.NewTask()
	task.FillByID(r.FormValue("TaskID"))

	user := model.NewUser()
	user.FillByID(session.GetUserUUID(r))

	t := time.Now()

	record := model.NewTimeRecord()
	record.SetUser(user)
	record.SetTask(task)
	record.SetDescription(r.FormValue("Description"))
	record.SetDate(t.Format(common.DateFormat))
	record.SetStart(t.Format(common.TimeFormat))

	err := record.Save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// TimerStopPOST stops timer for currently logged in user and updates TimeRecord with proper End time
func TimerStopPOST(w http.ResponseWriter, r *http.Request) {
	records := model.QueryTimeRecords("UserID = ? AND End IS NULL", session.GetUserUUID(r))
	if len(records) == 0 {
		http.Error(w, "No active TimeRecord!", http.StatusInternalServerError)
		return
	}

	t := time.Now()
	now := t.Format(common.TimeFormat)

	for _, v := range records {
		v.SetStop(&now)
		err := v.Save()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
