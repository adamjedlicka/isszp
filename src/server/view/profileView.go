package view

import (
	"net/http"
)

func ProfileGET(w http.ResponseWriter, r *http.Request) {
	view := NewView(r, "Profile")
	view.AppendTemplates("profile/profile")

	// TODO

	view.Render(w)
}
