package view

import (
	"net/http"

	"github.com/gorilla/mux"

	"gitlab.fit.cvut.cz/isszp/isszp/src/model"
)

func TimerecordsGET(w http.ResponseWriter, r *http.Request) {
	view := NewView(r, "Timerecords")
	view.AppendTemplates("timerecords/timerecords", "component/timerecord-list")

	view.Vars["Timerecords"] = model.QueryTimeRecords()

	view.Render(w)
}

func TimerecordViewGET(w http.ResponseWriter, r *http.Request) {
	view := NewView(r, "New Timerecord")
	view.AppendTemplates("timerecords/timerecord-view")

	tr := model.NewTimeRecord()
	err := tr.FillByID(mux.Vars(r)["ID"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	view.Vars["Action"] = "view"
	view.Vars["Timerecord"] = tr
	view.Vars["Users"] = model.QueryUsers()
	view.Vars["Tasks"] = model.QueryTasks()

	view.Vars["readonly"] = "readonly"
	view.Vars["disabled"] = "disabled"

	view.Render(w)
}

func TimerecordEditGET(w http.ResponseWriter, r *http.Request) {
	view := NewView(r, "New Timerecord")
	view.AppendTemplates("timerecords/timerecord-view")

	tr := model.NewTimeRecord()
	err := tr.FillByID(mux.Vars(r)["ID"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if tr.InProgress() == true {
		http.Redirect(w, r, "/timerecord/view/"+tr.GetID(), http.StatusSeeOther)
	}

	view.Vars["Action"] = "edit"
	view.Vars["Timerecord"] = tr
	view.Vars["Users"] = model.QueryUsers()
	view.Vars["Tasks"] = model.QueryTasks("WorkerID = ?", tr.GetUser().GetID())

	view.Render(w)
}

func TimerecordDeleteGET(w http.ResponseWriter, r *http.Request) {
	tr := model.NewTimeRecord()
	err := tr.FillByID(mux.Vars(r)["ID"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if tr.InProgress() {
		http.Redirect(w, r, "/timerecords", http.StatusSeeOther)
		return
	}

	err = tr.Delete()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/timerecords", http.StatusSeeOther)
}

func TimerecordSavePOST(w http.ResponseWriter, r *http.Request) {
	tr := model.NewTimeRecord()

	id := r.FormValue("ID")
	if id != "" {
		err := tr.FillByID(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	task := model.NewTask()
	task.FillByID(r.FormValue("TaskID"))

	user := model.NewUser()
	user.FillByID(r.FormValue("UserID"))

	tr.SetTask(task)
	tr.SetUser(user)
	tr.SetDate(r.FormValue("Date"))

	startTime := r.FormValue("StartTime")
	if startTime != "" {
		tr.SetStart(startTime)
	}

	stopTime := r.FormValue("StopTime")
	if stopTime != "" {
		tr.SetStop(&stopTime)
	}

	err := tr.Save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/timerecord/view/"+tr.GetID(), http.StatusSeeOther)
}
