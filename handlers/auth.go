package handlers

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/claisne/snippetdb/models"
	"github.com/claisne/snippetdb/store"
	"github.com/gorilla/context"
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
	err := getSignTemplates.Execute(w, nil)
	if err != nil {
		logrus.Warn(err.Error())
	}
}

func PostRegister(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	user, err := models.UserFromForm(r.Form)
	if err != nil {
		logrus.Warn(err.Error())
		return
	}

	store := context.Get(r, "store").(store.Store)
	fmt.Println("Saving")
	err = store.User().Save(user)
	if err != nil {
		logrus.Warn(err.Error())
		return
	}
}
