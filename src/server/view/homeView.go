package view

import (
	"net/http"
)

func HomeGET(w http.ResponseWriter, r *http.Request) {
	view := NewView(r, "Home")

	view.Render(w)
}
