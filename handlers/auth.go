package handlers

import (
	"errors"
	"html/template"
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/claisne/snippetdb/libhttp"
	"github.com/claisne/snippetdb/models"
	"github.com/claisne/snippetdb/store"
	"github.com/gorilla/context"
	"github.com/gorilla/sessions"
)

var getSignTemplates *template.Template

func init() {
	var err error
	getSignTemplates, err = template.ParseFiles("templates/layout.html", "templates/auth/sign.html")
	if err != nil {
		logrus.Fatal("Failed to parse sign templates")
	}
}

func GetSign(w http.ResponseWriter, r *http.Request) {
	err := getSignTemplates.Execute(w, nil)
	if err != nil {
		logrus.Warn(err.Error())
	}
}

func PostLogin(w http.ResponseWriter, r *http.Request) {
	// Get credentials
	r.ParseForm()
	username := r.Form.Get("username")
	password := r.Form.Get("password")

	// Find User
	store := context.Get(r, "store").(store.Store)
	user, err := store.User().GetByUsername(username)
	if err != nil {
		libhttp.HandleErrorJson(w, errors.New("No user found"))
		return
	}

	// Check password
	if !user.ComparePassword(password) {
		libhttp.HandleErrorJson(w, errors.New("Password mismatch"))
		return
	}

	// Update Session
	sessionStore := context.Get(r, "sessionStore").(sessions.Store)
	session, _ := sessionStore.Get(r, "snippetdb-session")
	session.Values["user"] = user

	err = session.Save(r, w)
	if err != nil {
		libhttp.HandleErrorJson(w, err)
		return
	}

	http.Redirect(w, r, "/account", 302)
}

func PostRegister(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	user, err := models.NewUserFromForm(r.Form)
	if err != nil {
		logrus.Warn(err.Error())
		return
	}

	store := context.Get(r, "store").(store.Store)
	err = store.User().Save(user)
	if err != nil {
		logrus.Warn(err.Error())
		return
	}

	PostLogin(w, r)
}
