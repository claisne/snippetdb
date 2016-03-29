package handlers

import (
	"html/template"
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/claisne/snippetdb/models"
	"github.com/claisne/snippetdb/store"
	"github.com/gorilla/context"
	"github.com/gorilla/sessions"
)

var getSignTemplates *template.Template

func init() {
	var err error
	getSignTemplates, err = template.ParseFiles("templates/layout.html", "templates/auth/sign.html", "templates/sign.html")
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Warn("Unable to parse GetSign templates")
	}
}

func GetSign(w http.ResponseWriter, r *http.Request) {
	sessionStore := context.Get(r, "sessionStore").(sessions.Store)
	session, _ := sessionStore.Get(r, "snippetdb-session")

	errors := getFlashesErrors(session)
	session.Save(r, w)

	data := struct {
		Errors []*Error
	}{
		Errors: errors,
	}

	err := getSignTemplates.Execute(w, data)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Warn("Unable to execute GetSign templates")
	}
}

func PostLogin(w http.ResponseWriter, r *http.Request) {
	// Get credentials
	r.ParseForm()
	username := r.Form.Get("username")
	password := r.Form.Get("password")

	// Potential Error
	loginErr := &Error{
		StatusCode:   403,
		RedirectPath: "/login",
	}

	// Find User
	store := context.Get(r, "store").(store.Store)
	user, err := store.User().GetByUsername(username)
	if err != nil {
		loginErr.Message = "No user with this username."
		loginErr.Render(w, r)
		return
	}

	// Check password
	if !user.ComparePassword(password) {
		loginErr.Message = "The password is incorrect."
		loginErr.Render(w, r)
		return
	}

	// Update Session
	sessionStore := context.Get(r, "sessionStore").(sessions.Store)
	session, _ := sessionStore.Get(r, "snippetdb-session")
	session.Values["user"] = user

	err = session.Save(r, w)
	if err != nil {
		http.Redirect(w, r, "/error", 500)
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
