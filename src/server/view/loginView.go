package view

import (
	"log"
	"net/http"
	"time"

	"gitlab.fit.cvut.cz/isszp/isszp/src/controller"
	"gitlab.fit.cvut.cz/isszp/isszp/src/model"
	"gitlab.fit.cvut.cz/isszp/isszp/src/server/session"
)

func LoginGET(w http.ResponseWriter, r *http.Request) {
	if session.IsLoggedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	view := NewView(r, "Login")
	view.AppendTemplates("login")

	view.Render(w)
}

func LoginPOST(w http.ResponseWriter, r *http.Request) {
	userName := r.PostFormValue("UserName")
	password := r.PostFormValue("Password")

	log.Println("Attemt to login from: ", userName)

	user := model.NewUser()
	err := user.FillByUserName(userName)
	if err != nil {
		log.Println(err)
		http.Error(w, "Bad UserName or Password!", http.StatusForbidden)
		return
	}

	ok := controller.CheckUserHashedPassword(user, password)
	if !ok {
		http.Error(w, "Bad UserName or Password!", http.StatusForbidden)
		return
	}

	s, err := session.Store.Get(r, session.Login)
	if err != nil {
		s.Options.MaxAge = -1
		err = s.Save(r, w)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	s.Values["LoginTime"] = time.Now().String()
	s.Values["UserName"] = user.GetUserName()
	s.Values["UUID"] = user.GetID()
	s.Values["IsLoggedIn"] = true

	err = s.Save(r, w)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func LogoutGET(w http.ResponseWriter, r *http.Request) {
	s, err := session.Store.Get(r, session.Login)
	if err != nil {
		log.Println(err)
		http.Error(w, "Not logged in!", http.StatusInternalServerError)
		return
	}

	s.Options.MaxAge = -1
	err = s.Save(r, w)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
