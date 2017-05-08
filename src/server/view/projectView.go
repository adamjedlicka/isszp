package view

import (
	"net/http"

	"gitlab.fit.cvut.cz/isszp/isszp/src/model"
)

func ProjectsGET(w http.ResponseWriter, r *http.Request) {
	view := NewView(r, "Projects")
	view.AppendTemplates("projects/projects", "component/project-list")

	view.Vars["Projects"] = model.QueryProjects()

	view.Render(w)
}
