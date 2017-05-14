package view

import (
	"net/http"

	"gitlab.fit.cvut.cz/isszp/isszp/src/model"
)

func TimerecordsGET(w http.ResponseWriter, r *http.Request) {
	view := NewView(r, "Timerecords")
	view.AppendTemplates("timerecords/timerecords", "component/timerecord-list")

	view.Vars["Timerecords"] = model.QueryTimeRecords()

	view.Render(w)
}
