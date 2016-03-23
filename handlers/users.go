package handlers

import (
	"html/template"
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/claisne/snippetdb/libhttp"
	"github.com/claisne/snippetdb/models"
	"github.com/claisne/snippetdb/store"

	"github.com/gorilla/context"
)

var getAccountTemplate *template.Template

func init() {
	var err error
	getAccountTemplate, err = template.ParseFiles("templates/layout.html", "templates/users/account.html", "templates/sign-modal.html")
	if err != nil {
		logrus.Fatal("Failed to parse account templates")
	}
}

func GetAccount(w http.ResponseWriter, r *http.Request) {
	user := context.Get(r, "user").(*models.User)

	err := getAccountTemplate.Execute(w, user)
	if err != nil {
		logrus.Warn(err.Error())
	}
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	id, err := getIdFromPath(r)
	if err != nil {
		libhttp.HandleErrorJson(w, err)
		return
	}

	store := context.Get(r, "store").(store.Store)

	user, err := store.User().Get(id)
	if err != nil {
		libhttp.HandleErrorJson(w, err)
		return
	}

	jsonBytes, err := user.ToJson()
	if err != nil {
		libhttp.HandleErrorJson(w, err)
		return
	}

	w.Header().Set("Content-Type", "text/json")
	w.Write(jsonBytes)
	return
}
