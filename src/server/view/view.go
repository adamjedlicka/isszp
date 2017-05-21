// package view is HTTP implementation of view layer of the application
package view

import (
	"html/template"
	"net/http"

	"strconv"

	"gitlab.fit.cvut.cz/isszp/isszp/src/common"
	"gitlab.fit.cvut.cz/isszp/isszp/src/model"
	"gitlab.fit.cvut.cz/isszp/isszp/src/server/session"
)

var (
	templateBase = "layout/base"
)

// View is default template of the web page. It is responsible for rendering the page, marking proper files
// to be sent to the client, ... Every view should be created using this struct
type View struct {
	Name      string
	Vars      map[string]interface{}
	L         map[string]string
	templates []string
}

// NewView creates new view based on passed in http.Request and name
// It always creates 2 global variables in template space
//   - Name : name of the current view
//   - IsLoggedInt : bool if user is logged in or not. if true then two new variables are added:
//     - LoggedUser : model.User of currently logged in user
//     - UserName : userName of currently logged in user
func NewView(r *http.Request, name string) *View {
	v := new(View)
	v.Name = name
	v.Vars = make(map[string]interface{})
	v.L = make(map[string]string)
	v.templates = make([]string, 0)
	v.templates = append(v.templates, templateBase)

	v.Vars["Name"] = name
	v.Vars["IsLoggedIn"] = session.IsLoggedIn(r)
	if session.IsLoggedIn(r) {
		user := model.NewUser()
		user.FillByUserName(session.GetUserName(r))
		v.Vars["LoggedUser"] = user
		v.Vars["UserName"] = user.GetUserName()
	}

	return v
}

// AppendTemplates adds html templates to the render pipeline which are later sent to the client
// every HTML template (except base.html) hase to be added with AppendTemplates
// template paths are without .html suffix and path begins in /template direcotry
// example:
//   v.AppendTemplates("/firms/firm-view", "/component/firm-list")
func (v *View) AppendTemplates(templates ...string) {
	v.templates = append(v.templates, templates...)
}

// Render renders the HTML page and sends it to he supplied ResponseWriter
func (v *View) Render(w http.ResponseWriter) {
	templateList := make([]string, len(v.templates))

	for i, name := range v.templates {
		path := "template/" + name + ".html"

		templateList[i] = path
	}

	template, err := template.ParseFiles(templateList...)
	if err != nil {
		http.Error(w, "Template Parse Error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	v.Vars["L"] = v.L
	err = template.Execute(w, v.Vars)
	if err != nil {
		http.Error(w, "Template Execute Error: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

// SetPagination computes pagination for the current view and set Pages variable which is later used in template
// to render the pagination buttons
func (v *View) SetPagination(r *http.Request, len, itemsPerPage int) (int, int) {
	nr := len / itemsPerPage
	if len%itemsPerPage != 0 {
		nr++
	}

	pages := make([]int, nr)
	for i := 0; i < nr; i++ {
		pages[i] = i + 1
	}
	v.Vars["Pages"] = pages

	page, err := strconv.Atoi(r.URL.Query().Get("p"))
	if err != nil {
		page = 0
	}

	if page >= nr {
		page = common.Max(nr-1, 0)
	}

	return page * itemsPerPage, common.Min((page+1)*itemsPerPage, len)
}
