package view

import (
	"net/http"

	"gitlab.fit.cvut.cz/isszp/isszp/src/model"

	"github.com/gorilla/mux"
)

const FirmsPerPage = 5

func FirmsGET(w http.ResponseWriter, r *http.Request) {
	view := NewView(r, "Firm :: list")
	view.AppendTemplates("firms/firms", "component/firm-list", "component/pagination")

	firms := model.QueryFirms()
	from, to := view.SetPagination(r, len(firms), FirmsPerPage)
	view.Vars["Firms"] = firms[from:to]

	view.Render(w)
}

func FirmNewGET(w http.ResponseWriter, r *http.Request) {
	view := NewView(r, "Firm :: new")
	view.AppendTemplates("firms/firm-view")

	view.Vars["Action"] = "new"

	view.Render(w)
}

func FirmViewGET(w http.ResponseWriter, r *http.Request) {
	view := NewView(r, "Firm :: detail")
	view.AppendTemplates("firms/firm-view")

	id := mux.Vars(r)["ID"]
	firm := model.NewFirm()
	err := firm.FillByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	view.Name += firm.GetName()

	view.Vars["Firm"] = firm
	view.Vars["Action"] = "view"

	view.Vars["readonly"] = "readonly"

	view.Render(w)
}

func FirmEditGET(w http.ResponseWriter, r *http.Request) {
	view := NewView(r, "Firm :: edit")
	view.AppendTemplates("firms/firm-view")

	id := mux.Vars(r)["ID"]
	firm := model.NewFirm()
	err := firm.FillByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	view.Vars["Firm"] = firm
	view.Vars["Action"] = "edit"

	view.Render(w)
}

func FirmDelGET(w http.ResponseWriter, r *http.Request) {
	f := model.NewFirm()
	err := f.FillByID(mux.Vars(r)["ID"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = f.Delete()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set URL to the URL user came from
	url := r.Header.Get("Referer")
	if url == "" {
		url = "/firms"
	}

	http.Redirect(w, r, url, http.StatusSeeOther)
}

func FirmSavePOST(w http.ResponseWriter, r *http.Request) {
	f := model.NewFirm()

	id := r.FormValue("ID")
	if id != "" {
		err := f.FillByID(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	f.SetName(r.FormValue("Name"))
	f.SetEmail(r.FormValue("Email"))
	f.SetTelNumber(r.FormValue("TelNumber"))
	f.SetDescription(r.FormValue("Description"))
	err := f.Save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/firm/view/"+f.GetID(), http.StatusSeeOther)
}
