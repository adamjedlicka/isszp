package view

import (
	"net/http"

	"github.com/gorilla/mux"

	"gitlab.fit.cvut.cz/isszp/isszp/src/model"
)

func ProjectsGET(w http.ResponseWriter, r *http.Request) {
	view := NewView(r, "Projects")
	view.AppendTemplates("projects/projects", "component/project-list")

	view.Vars["Projects"] = model.QueryProjects()

	view.Render(w)
}

func ProjectNewGET(w http.ResponseWriter, r *http.Request) {
	view := NewView(r, "New project")
	view.AppendTemplates("projects/project-view")

	view.Vars["Action"] = "new"
	view.Vars["Users"] = model.QueryUsers()
	view.Vars["Firms"] = model.QueryFirms()

	view.Render(w)
}

func ProjectViewGET(w http.ResponseWriter, r *http.Request) {
	view := NewView(r, "New project")
	view.AppendTemplates("projects/project-view")

	project := model.NewProject()
	err := project.FillByID(mux.Vars(r)["ID"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	view.Vars["Action"] = "view"
	view.Vars["Project"] = project
	view.Vars["Users"] = model.QueryUsers()
	view.Vars["Firms"] = model.QueryFirms()

	view.Vars["readonly"] = "readonly"
	view.Vars["disabled"] = "disabled"

	view.Render(w)
}

func ProjectSavePOST(w http.ResponseWriter, r *http.Request) {
	p := model.NewProject()

	id := r.FormValue("ID")
	if id != "" {
		err := p.FillByID(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	p.SetName(r.FormValue("Name"))
	p.SetCode(r.FormValue("Code"))
	p.SetDescription(r.FormValue("Description"))
	p.SetStartDate(r.FormValue("StartDate"))

	planEndDate := r.FormValue("PlanEndDate")
	if planEndDate == "" {
		p.SetPlanEndDate(nil)
	} else {
		p.SetPlanEndDate(&planEndDate)
	}

	endDate := r.FormValue("EndDate")
	if endDate == "" {
		p.SetEndDate(nil)
	} else {
		p.SetEndDate(&endDate)
	}

	maintainer := model.NewUser()
	err := maintainer.FillByID(r.FormValue("MaintainerID"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	p.SetMaintainer(maintainer)

	firm := model.NewFirm()
	err = firm.FillByID(r.FormValue("FirmID"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	p.SetFirm(firm)

	err = p.Save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/project/view/"+p.GetID(), http.StatusSeeOther)
}
