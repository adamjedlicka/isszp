package view

import (
	"net/http"

	"gitlab.fit.cvut.cz/isszp/isszp/src/model"

	"github.com/gorilla/mux"
)

func TasksGET(w http.ResponseWriter, r *http.Request) {
	view := NewView(r, "Tasks")
	view.AppendTemplates("tasks/tasks", "component/task-list")

	view.Vars["Tasks"] = model.QueryTasks()

	view.Render(w)
}

func TaskNewGET(w http.ResponseWriter, r *http.Request) {
	view := NewView(r, "New task")
	view.AppendTemplates("tasks/task-view")

	view.Vars["Action"] = "new"
	view.Vars["Projects"] = model.QueryProjects()
	view.Vars["Users"] = model.QueryUsers()

	view.Render(w)
}

func TaskViewGET(w http.ResponseWriter, r *http.Request) {
	view := NewView(r, "New task")
	view.AppendTemplates("tasks/task-view", "component/comment")

	task := model.NewTask()
	err := task.FillByID(mux.Vars(r)["ID"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	view.Vars["Action"] = "view"
	view.Vars["Task"] = task
	view.Vars["Projects"] = model.QueryProjects()
	view.Vars["Users"] = model.QueryUsers()
	view.Vars["Comments"] = model.QueryComments("TaskID = ?", task.GetID())

	view.Vars["readonly"] = "readonly"
	view.Vars["disabled"] = "disabled"

	view.Render(w)
}

func TaskEditGET(w http.ResponseWriter, r *http.Request) {
	view := NewView(r, "New task")
	view.AppendTemplates("tasks/task-view")

	task := model.NewTask()
	err := task.FillByID(mux.Vars(r)["ID"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	view.Vars["Action"] = "edit"
	view.Vars["Task"] = task
	view.Vars["Projects"] = model.QueryProjects()
	view.Vars["Users"] = model.QueryUsers()

	view.Render(w)
}

func TaskDeleteGET(w http.ResponseWriter, r *http.Request) {
	task := model.NewTask()
	err := task.FillByID(mux.Vars(r)["ID"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = task.Delete()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/tasks", http.StatusSeeOther)
}

func TaskSavePOST(w http.ResponseWriter, r *http.Request) {
	t := model.NewTask()

	id := r.FormValue("ID")
	if id != "" {
		err := t.FillByID(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	t.SetName(r.FormValue("Name"))
	t.SetDescription(r.FormValue("Description"))
	t.SetStartDate(r.FormValue("StartDate"))

	planEndDate := r.FormValue("PlanEndDate")
	if planEndDate == "" {
		t.SetPlanEndDate(nil)
	} else {
		t.SetPlanEndDate(&planEndDate)
	}

	endDate := r.FormValue("EndDate")
	if endDate == "" {
		t.SetEndDate(nil)
	} else {
		t.SetEndDate(&endDate)
	}

	maintainer := model.NewUser()
	err := maintainer.FillByID(r.FormValue("MaintainerID"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	worker := model.NewUser()
	err = worker.FillByID(r.FormValue("WorkerID"))
	if err != nil {
		worker = nil
	}

	project := model.NewProject()
	err = project.FillByID(r.FormValue("ProjectID"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	t.SetMaintainer(maintainer)
	t.SetWorker(worker)
	t.SetProject(project)

	err = t.Save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/task/view/"+t.GetID(), http.StatusSeeOther)
}
